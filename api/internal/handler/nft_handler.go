package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/abi"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/auth"
	"github.com/project-nova/backend/pkg/logger"
)

func NewUpdateNftBackstoryHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		message := "hello"

		type UpdateNftBackstoryRequestBody struct {
			CollectionAddress string `json:"collectionAddress"`
			WalletAddress     string `json:"walletAddress"`
			Backstory         string `json:"backstory"`
			Signature         string `json:"signature"`
		}

		var requestBody UpdateNftBackstoryRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.String(http.StatusBadRequest, "invalid request body")
			return
		}

		recoveredAddress, err := auth.RecoverAddress(message, requestBody.Signature)
		if err != nil {
			logger.Errorf("failed to recover address: %v", err)
			return
		}

		if recoveredAddress != requestBody.WalletAddress {
			logger.Errorf("wallet verification failed, recovered address: %s, wallet address: %s", recoveredAddress, requestBody.WalletAddress)
			c.String(http.StatusForbidden, "The wallet doesn't have permission for this operation")
			return
		}

		tokenId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.String(http.StatusBadRequest, "token id is invalid")
			return
		}

		nftToken, err := nftTokenRepository.GetNftByTokenId(tokenId, requestBody.CollectionAddress)
		if err != nil {
			logger.Errorf("Failed to get nft by token id: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		if *nftToken.OwnerAddress != requestBody.WalletAddress {
			logger.Errorf("The request wallet is not the owner of the nft, owner address: %s, wallet address: %s", nftToken.OwnerAddress, requestBody.WalletAddress)
			c.String(http.StatusForbidden, "The wallet doesn't have permission for this operation")
			return
		}

		nftToken, err = nftTokenRepository.UpdateNftBackstory(tokenId, requestBody.CollectionAddress, &requestBody.Backstory)
		if err != nil {
			logger.Errorf("Failed to update nft backstory: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

func NewGetNftsHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func NewUpdateNftHandler(nftTokenRepository repository.NftTokenRepository, client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert franchise id: %v", err)
			c.String(http.StatusBadRequest, "franchise id is invalid")
			return
		}

		tokenId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.String(http.StatusBadRequest, "token id is invalid")
			return
		}

		collectionAddress := c.DefaultQuery("collectionAddress", "")
		if collectionAddress == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", collectionAddress))
			return
		}

		address := common.HexToAddress(collectionAddress)
		contract, err := abi.NewAbi(address, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		uri, err := contract.TokenURI(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		var nft *repository.NftTokenModel
		if strings.Contains(uri, "ipfs") {
			nft, err = createIpfsNftRecord(uri, franchiseId, int(tokenId), collectionAddress)
			logger.Errorf("Failed to construct ipfs nft record: %v", err)
		} else {
			nft, err = createSvgNftRecord(uri, franchiseId, int(tokenId), collectionAddress)
			logger.Errorf("Failed to construct svg nft record: %v", err)
		}
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		nftToken, err := nftTokenRepository.CreateNft(nft)
		if err != nil {
			logger.Errorf("Failed to create nft token db record: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

func createSvgNftRecord(uri string, franchiseId int64, tokenId int, collectionAddress string) (*repository.NftTokenModel, error) {
	splittedStr := strings.Split(uri, ",")
	if len(splittedStr) != 2 {
		return nil, fmt.Errorf("failed to split uri string to 2 parts. uri: %v", uri)
	}

	decodedUri, err := base64.URLEncoding.DecodeString(splittedStr[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode uri: %v", err)
	}

	var result map[string]any
	err = json.Unmarshal(decodedUri, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	resultStr := string(decodedUri)

	return &repository.NftTokenModel{
		ID:                uuid.New().String(),
		FranchiseId:       franchiseId,
		TokenId:           tokenId,
		CollectionAddress: collectionAddress,
		Traits:            &resultStr,
	}, nil
}

func createIpfsNftRecord(uri string, franchiseId int64, tokenId int, collectionAddress string) (*repository.NftTokenModel, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to query ipfs link: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read ipfs response body: %v", err)
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	imageUrl := result["animation_url"].(string)
	traits, err := json.Marshal(result["attributes"])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %v", err)
	}
	traitsStr := string(traits)

	return &repository.NftTokenModel{
		ID:                uuid.New().String(),
		FranchiseId:       franchiseId,
		TokenId:           tokenId,
		CollectionAddress: collectionAddress,
		ImageUrl:          &imageUrl,
		Traits:            &traitsStr,
	}, nil
}
