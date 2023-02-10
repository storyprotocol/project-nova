package main

import (
	"errors"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/config"
	"github.com/project-nova/backend/api/internal/handler"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/database"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"gorm.io/gorm"
)

type WhitelistWalletModel struct {
	ID          string `gorm:"primaryKey;column:id"`
	Address     string
	MerkleProof string
	CreatedAt   time.Time
}

func (WhitelistWalletModel) TableName() string {
	return "whitelist_wallet"
}

type MembershipModel struct {
	ID        string
	Address   string
	Logins    uint64
	CreatedAt time.Time
	Username  string
}

func (MembershipModel) TableName() string {
	return "membership"
}

type MembershipResp struct {
	Name     string
	Symbol   string
	UserName string
	Grade    string
	Count    uint64
	URI      string
	LogIns   uint64
	JoinedAt time.Time
}

func main() {
	r := gin.Default()

	flag.Parse()

	Logger, err := logger.InitLogger(logger.Levels.Info)
	if err != nil {
		logger.Fatalf("Failed to init logger, error: %v", err)
	}
	defer func() {
		_ = Logger.Sync()
	}()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to init config, error: %v", err)
	}

	db, err := database.NewGormDB(cfg.DatabaseURI)
	if err != nil {
		logger.Fatalf("Failed to connect to DB, error: %v", err)
	}

	client, err := ethclient.Dial(cfg.ProviderURL)
	if err != nil {
		logger.Fatalf("Failed to connect to the blockchain provider, error: %v", err)
	}

	walletMerkleProofRepository := repository.NewWalletMerkleProofDbImpl(db)
	nftTokenRepository := repository.NewNftTokenDbImpl(db)
	storyChapterRepository := repository.NewStoryChapterDbImpl(db)
	storyInfoRepository := repository.NewStoryInfoDbImpl(db)
	storyContentRepository, err := repository.NewStoryContentFsImpl(cfg.ContentPath)
	if err != nil {
		logger.Errorf("Failed to init story content fs implementation: %v", err)
		return
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Healthy")
	})

	// Endpoint to get the metadata of all story nfts owned by the wallet
	r.GET("/v1/wallet/:walletAddress/nfts", handler.NewGetWalletNftsHandler(nftTokenRepository))

	// Endpoint to get the merkle proof for the wallet address per allowlist
	r.GET("/v1/wallet/:walletAddress/proof", handler.NewGetWalletProofHandler(walletMerkleProofRepository))

	// Endpoint to get all story chapters' information
	r.GET("/v1/story/:storyNum/chapters", handler.NewGetStoryChaptersHandler(storyChapterRepository, storyInfoRepository))

	// Endpoint to get story chapter contents
	r.GET("/v1/story/:storyNum/chapter/:chapterNum/contents", handler.NewGetStoryChapterContentsHandler(storyContentRepository))

	// Endpoint to update nft backstory for the nft owner
	r.POST("/v1/nft/:id/backstory", handler.NewUpdateNftBackstoryHandler(nftTokenRepository))

	// Endpoint to get the metadata of story nfts
	r.GET("/v1/nfts", handler.NewGetNftsHandler(nftTokenRepository))

	// Admin Endpoint to batch update nft metadata
	r.POST("/v1/nft/:id", handler.NewUpdateNftHandler(nftTokenRepository, client))

	// Deprecated
	r.GET("/mint/proof", func(c *gin.Context) {
		address := c.DefaultQuery("address", "")

		// Validate address
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		result := &WhitelistWalletModel{}
		r := db.Where("address = ?", address).First(&result)
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		if r.Error != nil {
			logger.Errorf("Failed to query db: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"proof": result.MerkleProof,
		})
	})

	// Deprecated
	r.GET("/membership", func(c *gin.Context) {
		privateKey := "a6cd3f393b1cddf8be66e2ff784640adbafbce852267ec1ec000000000000000" // Fake Key
		contractAddress := "0x64432E5A76a93e79be2f7F3F12982059a32Fd794"
		address := c.DefaultQuery("address", "")
		// Validate address
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		dbResult := &MembershipModel{}
		r := db.Where("address = ?", address).First(&dbResult)
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		if r.Error != nil {
			logger.Errorf("Failed to query db: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
			PrivateKey: privateKey,
		})
		if err != nil {
			logger.Errorf("Failed to set up sdk: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		content, err := os.ReadFile(cfg.AbiPath + "story_pass.json")
		if err != nil {
			logger.Errorf("Failed to read JSON files: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		abi := string(content)
		contract, err := sdk.GetContractFromAbi(contractAddress, abi)
		if err != nil {
			logger.Errorf("Failed to get contract: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balance, err := contract.Call(c, "balanceOf", address)
		if err != nil {
			logger.Errorf("Failed to get balance: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balanceBigInt, ok := balance.(*big.Int)
		if !ok {
			logger.Errorf("Failed to convert balance\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		if balanceBigInt.Uint64() == 0 {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		name, err := contract.Call(c, "name")
		if err != nil {
			logger.Errorf("Failed to get name: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		nameStr, ok := name.(string)
		if !ok {
			logger.Errorf("Failed to convert name to string\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		symbol, err := contract.Call(c, "symbol")
		if err != nil {
			logger.Errorf("Failed to get symbol: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		symbolStr, ok := symbol.(string)
		if !ok {
			logger.Errorf("Failed to convert symbol to string\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenID, err := contract.Call(c, "tokenOfOwnerByIndex", address, 0)
		if err != nil {
			logger.Errorf("Failed to get tokenID: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenIDBigInt, ok := tokenID.(*big.Int)
		if !ok {
			logger.Errorf("Failed to convert tokenID to bigint\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenURI, err := contract.Call(c, "tokenURI", int(tokenIDBigInt.Int64()))
		if err != nil {
			logger.Errorf("Failed to get tokenURI: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenURIStr, ok := tokenURI.(string)
		if !ok {
			logger.Errorf("Failed to convert tokenURI to string\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenGradeID, err := contract.Call(c, "tokenGrades", int(tokenIDBigInt.Int64()))
		if err != nil {
			logger.Errorf("Failed to get tokenGradeID: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenGradeIDBigInt, ok := tokenGradeID.(*big.Int)
		if !ok {
			logger.Errorf("Failed to convert tokenGradeID to big int\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenGrade, err := contract.Call(c, "grades", int(tokenGradeIDBigInt.Int64()))
		if err != nil {
			logger.Errorf("Failed to get tokenGrade: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenGradeStr, ok := tokenGrade.(string)
		if !ok {
			logger.Errorf("Failed to convert tokenGradeID to big int\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, &MembershipResp{
			Name:     nameStr,
			Symbol:   symbolStr,
			UserName: dbResult.Username,
			Count:    balanceBigInt.Uint64(),
			URI:      tokenURIStr,
			Grade:    tokenGradeStr,
			LogIns:   dbResult.Logins,
			JoinedAt: dbResult.CreatedAt,
		})
	})

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	_ = r.Run(port)
}
