package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

// NewUpdateNftBackstoryHandler: https://documenter.getpostman.com/view/25015244/2s935ppNga#d4af7069-ec5a-440a-b158-55524412da58
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

		if nftToken.OwnerAddress == nil || *nftToken.OwnerAddress != requestBody.WalletAddress {
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

		nft, err := createNftRecord(uri, franchiseId, int(tokenId), collectionAddress)
		if err != nil {
			logger.Errorf("Failed to construct svg nft record: %v", err)
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

func createNftRecord(uri string, franchiseId int64, tokenId int, collectionAddress string) (*repository.NftTokenModel, error) {
	splittedStr := strings.Split(uri, ",")
	if len(splittedStr) != 2 {
		return nil, fmt.Errorf("failed to split uri string to 2 parts. uri: %v", uri)
	}

	decodedMetadata, err := base64.URLEncoding.DecodeString(splittedStr[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode uri: %v", err)
	}

	result := struct {
		Name         *string
		Description  *string
		Image        *string
		AnimationUrl *string
		Attributes   []struct {
			TraitType string
			Value     string
		}
	}{}

	err = json.Unmarshal(decodedMetadata, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %v", err)
	}

	nft := &repository.NftTokenModel{
		ID:                uuid.New().String(),
		FranchiseId:       franchiseId,
		TokenId:           tokenId,
		CollectionAddress: collectionAddress,
	}

	if result.Name != nil {
		nft.Name = result.Name
	}
	if result.Description != nil {
		nft.Description = result.Description
	}
	if result.Image != nil {
		nft.Image = result.Image
	}
	if result.AnimationUrl != nil {
		nft.AnimationUrl = result.AnimationUrl
	}
	if result.Attributes != nil {
		traits, err := json.Marshal(result.Attributes)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal attributes: %v", err)
		}
		traitsStr := string(traits)
		nft.Traits = &traitsStr
	}

	return nft, nil
}
