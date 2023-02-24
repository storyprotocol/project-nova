package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/config"
	"github.com/project-nova/backend/api/internal/handler"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/database"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/middleware"
	"github.com/project-nova/backend/pkg/s3"
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

	ethClient, err := ethclient.Dial(cfg.ProviderURL)
	if err != nil {
		logger.Fatalf("Failed to connect to the blockchain provider, error: %v", err)
	}

	kmsClient := keymanagement.NewKmsClient(cfg.Region)

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
	})
	if err != nil {
		logger.Fatalf("Failed to create aws session: %v", err)
	}
	s3Client := s3.NewS3Client(awsSession)

	walletMerkleProofRepository := repository.NewWalletMerkleProofDbImpl(db)
	nftTokenRepository := repository.NewNftTokenDbImpl(db)
	nftCollectionRepository := repository.NewNftCollectionDbImpl(db)
	storyChapterRepository := repository.NewStoryChapterDbImpl(db)
	storyInfoRepository := repository.NewStoryInfoDbImpl(db)

	storyContentRepository, err := repository.NewStoryContentS3Impl(s3Client)
	if err != nil {
		logger.Fatalf("Failed to init story content s3 implementation: %v", err)
	}

	franchiseCollectionRepository := repository.NewFranchiseCollectionDbImpl(db)
	err = franchiseCollectionRepository.GetAndLoadFranchiseCollections()
	if err != nil {
		logger.Errorf("Failed to get and load franchise collections: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Healthy")
	})

	publicV1 := r.Group("/v1")
	publicV1.Use(cors.Default())
	{
		// Endpoint to get the merkle proof for the wallet address per allowlist
		publicV1.GET("/wallet/:walletAddress/proof", handler.NewGetWalletProofHandler(walletMerkleProofRepository))

		// Endpoint to get all story chapters' information
		publicV1.GET("/story/:franchiseId/:storyNum", handler.NewGetStoryChaptersHandler(storyChapterRepository, storyInfoRepository))

		// Endpoint to get story chapter contents
		publicV1.GET("/story/:franchiseId/:storyNum/:chapterNum", handler.NewGetStoryChapterContentsHandler(storyContentRepository))

		// Endpoint to get the metadata of story nfts
		publicV1.POST("/nft/list", handler.NewGetNftsHandler(nftTokenRepository, franchiseCollectionRepository))

		// Endpoint to update nft backstory for the nft owner
		publicV1.POST("/nft/:id/backstory", handler.NewUpdateNftBackstoryHandler(nftTokenRepository, cfg.AdminAuthMessage))

		// Endpoint to get the metadata of nft collection
		publicV1.GET("/nft/collections", handler.NewGetNftCollectionsHandler(nftCollectionRepository, franchiseCollectionRepository))

	}

	adminV1 := r.Group("/admin/v1")
	adminV1.Use(middleware.AuthAdmin(kmsClient, []byte(cfg.AdminAuthMessage), cfg.AuthKeyId))
	{
		// Admin Endpoint to fetch and create nft metadata
		adminV1.POST("/nft/:id", handler.NewCreateOrUpdateNftHandler(nftTokenRepository, ethClient))

		// Admin Endpoint to update nft owner address
		adminV1.POST("/nft/:id/owner", handler.NewUpdateNftOwnerHandler(nftTokenRepository, ethClient))

		// Admin Endpoint to delete nft
		adminV1.DELETE("/nft/:id", handler.NewDeleteNftHandler(nftTokenRepository))
	}

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	_ = r.Run(port)
}
