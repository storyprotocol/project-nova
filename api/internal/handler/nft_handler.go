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
	"github.com/project-nova/backend/api/internal/constant"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/abi/erc721"
	"github.com/project-nova/backend/pkg/auth"
	"github.com/project-nova/backend/pkg/logger"
	"gorm.io/gorm"
)

// NewGetNftCollectionsHandler create the handler to get nft collections data
func NewGetNftCollectionsHandler(
	nftCollectionRepository repository.NftCollectionRepository,
	franchiseCollectionRepository repository.FranchiseCollectionRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		collectionAddress := c.DefaultQuery("collectionAddress", "")
		franchiseId := c.DefaultQuery("franchiseId", "")

		collectionAddresses, err := getCollectionAddresses(franchiseId, collectionAddress, franchiseCollectionRepository)
		if err != nil {
			logger.Errorf("Failed to get collection addresses: %v", err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		results, err := nftCollectionRepository.GetCollections(collectionAddresses)
		if err != nil {
			logger.Errorf("Failed to get collections: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

// NewUpdateNftBackstoryHandler: https://documenter.getpostman.com/view/25015244/2s935ppNga#d4af7069-ec5a-440a-b158-55524412da58
func NewUpdateNftBackstoryHandler(nftTokenRepository repository.NftTokenRepository, authMessage string) func(c *gin.Context) {
	return func(c *gin.Context) {

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

		recoveredAddress, err := auth.RecoverAddress(authMessage, requestBody.Signature)
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

func NewGetNftsHandler(nftTokenRepository repository.NftTokenRepository, franchiseCollectionRepository repository.FranchiseCollectionRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		walletAddress := c.DefaultQuery("walletAddress", "")
		collectionAddress := c.DefaultQuery("collectionAddress", "")
		limitStr := c.DefaultQuery("limit", "")
		offsetStr := c.DefaultQuery("offset", "")
		franchiseId := c.DefaultQuery("franchiseId", "")

		collectionAddresses, err := getCollectionAddresses(franchiseId, collectionAddress, franchiseCollectionRepository)
		if err != nil {
			logger.Errorf("Failed to get collection addresses: %v", err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		var offset *int
		var limit *int
		if limitStr != "" {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				logger.Errorf("Failed to convert limit string to integer : %v", err)
				c.String(http.StatusBadRequest, "limit is invalid")
				return
			}
			if limitInt < 0 || limitInt > constant.NftMaxLimit {
				logger.Errorf("Invalid input limit: %d", limitInt)
				c.JSON(http.StatusBadRequest, gin.H{"message": "limit is invalid. limit should be <= 500"})
				return
			}
			limit = &limitInt
		}

		if offsetStr != "" {
			offsetInt, err := strconv.Atoi(offsetStr)
			if err != nil {
				logger.Errorf("Failed to convert offset string to integer : %v", err)
				c.String(http.StatusBadRequest, "offset is invalid")
				return
			}
			if offsetInt < 0 {
				logger.Errorf("Invalid input offset: %d", offsetInt)
				c.JSON(http.StatusBadRequest, gin.H{"message": "offset is invalid. offset should be >= 0"})
				return
			}
			offset = &offsetInt
		}

		result, err := nftTokenRepository.GetNfts(collectionAddresses, walletAddress, offset, limit)
		if err != nil {
			logger.Errorf("Failed to get nfts: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func NewCreateOrUpdateNftHandler(nftTokenRepository repository.NftTokenRepository, client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "token id is invalid"})
			return
		}

		collectionAddress := c.DefaultQuery("collectionAddress", "")
		if collectionAddress == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("input address is invalid, address: %s", collectionAddress)})
			return
		}

		address := common.HexToAddress(collectionAddress)
		contract, err := erc721.NewErc721(address, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		uri, err := contract.TokenURI(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		ownerAddress, err := contract.OwnerOf(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		nft, err := createNftRecord(uri, int(tokenId), ownerAddress.String(), collectionAddress)
		if err != nil {
			logger.Errorf("Failed to construct nft record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		_, err = nftTokenRepository.GetNftByTokenId(int(tokenId), collectionAddress)
		if err != nil {
			if err == gorm.ErrRecordNotFound { // nft token doesn't exist in DB. Create the nft record
				nftToken, err := nftTokenRepository.CreateNft(nft)
				if err != nil {
					logger.Errorf("Failed to create nft token db record: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
					return
				}
				c.JSON(http.StatusOK, nftToken)
				return
			}
			logger.Errorf("Failed to get nft record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
		// nft token exists. Update the record
		nftToken, err := nftTokenRepository.UpdateNft(nft)
		if err != nil {
			logger.Errorf("Failed to update nft token db record: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

func createNftRecord(uri string, tokenId int, ownerAddress string, collectionAddress string) (*repository.NftTokenModel, error) {
	splittedStr := strings.Split(uri, ",")
	if len(splittedStr) != 2 {
		return nil, fmt.Errorf("failed to split uri string to 2 parts. uri: %v", uri)
	}

	decodedMetadata, err := base64.StdEncoding.DecodeString(splittedStr[1])
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
		TokenId:           tokenId,
		OwnerAddress:      &ownerAddress,
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

func NewUpdateNftOwnerHandler(nftTokenRepository repository.NftTokenRepository, client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
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
		contract, err := erc721.NewErc721(address, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		ownerAddress, err := contract.OwnerOf(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		nftToken, err := nftTokenRepository.UpdateNftOwner(int(tokenId), collectionAddress, ownerAddress.String())
		if err != nil {
			logger.Errorf("Failed to update nft owner: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

func NewDeleteNftHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.String(http.StatusBadRequest, "token id is invalid")
			return
		}

		collectionAddress := c.DefaultQuery("collectionAddress", "")
		if collectionAddress == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input collection address is invalid, address: %s", collectionAddress))
			return
		}

		err = nftTokenRepository.DeleteNft(int(tokenId), collectionAddress)
		if err != nil {
			logger.Errorf("Failed to delete nft: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

func getCollectionAddresses(
	franchiseId string,
	collectionAddress string,
	franchiseCollectionRepository repository.FranchiseCollectionRepository,
) ([]string, error) {
	collectionAddresses := []string{}
	if collectionAddress != "" {
		collectionAddresses = append(collectionAddresses, collectionAddress)
	} else if franchiseId != "" {
		franchiseIdInt, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert franchise id: %v", err)
		}
		collectionAddresses, err = franchiseCollectionRepository.GetCollectionAddressesByFranchise(franchiseIdInt)
		if err != nil {
			return nil, fmt.Errorf("Failed to get collection addresses by franchise id: %v", err)
		}
	} else {
		return nil, fmt.Errorf("Neither collection address or franchise id is passed")
	}

	return collectionAddresses, nil
}
