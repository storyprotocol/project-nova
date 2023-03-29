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
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/abi/erc721"
	"github.com/project-nova/backend/pkg/auth"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
	"gorm.io/gorm"
)

// NewGetNftCollectionsHandler create the handler to get nft collections data.
// Doc: (To Be Added)
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
			c.JSON(http.StatusBadRequest, ErrorMessage(err.Error()))
			return
		}

		results, err := nftCollectionRepository.GetCollections(collectionAddresses)
		if err != nil {
			logger.Errorf("Failed to get collections: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

// NewUpdateNftBackstoryHandler: https://documenter.getpostman.com/view/25015244/2s935ppNga#d4af7069-ec5a-440a-b158-55524412da58
func NewUpdateNftBackstoryHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {

		var requestBody entity.UpdateNftBackstoryRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(requestBody.CollectionAddress)
		if err != nil {
			logger.Errorf("Invalid collection address: %s", requestBody.CollectionAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		walletAddress, err := utils.SanitizeAddress(requestBody.WalletAddress)
		if err != nil {
			logger.Errorf("Invalid wallet address: %s", requestBody.WalletAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid wallet address"))
			return
		}

		recoveredAddress, err := auth.RecoverAddress(requestBody.Message, requestBody.Signature)
		if err != nil {
			logger.Errorf("Failed to recover address: %v", err)
			c.JSON(http.StatusForbidden, ErrorMessage("The wallet doesn't have permission for this operation"))
			return
		}

		recoveredAddress, err = utils.SanitizeAddress(recoveredAddress)
		if err != nil {
			logger.Errorf("Wallet verification failed, invalid recovered address: %s", recoveredAddress)
			c.JSON(http.StatusForbidden, ErrorMessage("The wallet doesn't have permission for this operation"))
			return
		}

		if recoveredAddress != walletAddress {
			logger.Errorf("Wallet verification failed, recovered address: %s, wallet address: %s", recoveredAddress, walletAddress)
			c.JSON(http.StatusForbidden, ErrorMessage("The wallet doesn't have permission for this operation"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("token id is invalid"))
			return
		}

		nftToken, err := nftTokenRepository.GetNftByTokenId(tokenId, collectionAddress)
		if err != nil {
			logger.Errorf("Failed to get nft by token id: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		if nftToken.OwnerAddress == nil || *nftToken.OwnerAddress != walletAddress {
			logger.Errorf("The request wallet is not the owner of the nft, owner address: %s, wallet address: %s", *nftToken.OwnerAddress, walletAddress)
			c.JSON(http.StatusForbidden, ErrorMessage("The wallet doesn't have permission for this operation"))
			return
		}

		nftToken, err = nftTokenRepository.UpdateNftBackstory(tokenId, collectionAddress, &requestBody.Backstory)
		if err != nil {
			logger.Errorf("Failed to update nft backstory: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

// NewGetNftsHandler creates the handler to handle POST /nft/list request.
// Doc: (To be added)
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
			c.JSON(http.StatusBadRequest, ErrorMessage(err.Error()))
			return
		}

		if walletAddress != "" {
			walletAddress, err = utils.SanitizeAddress(walletAddress)
			if err != nil {
				logger.Errorf("Invalid wallet address: %s", walletAddress)
				c.JSON(http.StatusBadRequest, ErrorMessage("Invalid wallet address"))
				return
			}
		}

		var offset *int
		var limit *int
		if limitStr != "" {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				logger.Errorf("Failed to convert limit string to integer : %v", err)
				c.JSON(http.StatusBadRequest, ErrorMessage("limit is invalid"))
				return
			}
			if limitInt < 0 || limitInt > constant.NftMaxLimit {
				logger.Errorf("Invalid input limit: %d", limitInt)
				c.JSON(http.StatusBadRequest, ErrorMessage(fmt.Sprintf("limit is invalid. limit should be <= %d", constant.NftMaxLimit)))
				return
			}
			limit = &limitInt
		}

		if offsetStr != "" {
			offsetInt, err := strconv.Atoi(offsetStr)
			if err != nil {
				logger.Errorf("Failed to convert offset string to integer : %v", err)
				c.JSON(http.StatusBadRequest, ErrorMessage("offset is invalid"))
				return
			}
			if offsetInt < 0 {
				logger.Errorf("Invalid input offset: %d", offsetInt)
				c.JSON(http.StatusBadRequest, ErrorMessage("offset is invalid. offset should be >= 0"))
				return
			}
			offset = &offsetInt
		}

		// listOption includes options for filtering, ordering and pagination.
		// Currently only it implements filtering.
		// Example:
		// {
		//   "filter":{
		//	   "operator":"and",
		//	   "operands":[
		//	   	 {
		//		   "operator":"eq",
		//		   "field":"character",
		//		   "value":"Human"
		//	   	 }
		//	   ]
		//   }
		// }
		var listOption entity.ListOption
		err = c.ShouldBindJSON(&listOption)
		filter := listOption.Filter

		var nfts []*entity.NftTokenResponse
		var totalUnfiltered int
		if err == nil && filter != nil && filter.Validate() {
			nfts, totalUnfiltered, err = nftTokenRepository.GetFilteredNfts(collectionAddresses, walletAddress, filter, offset, limit)
			if err != nil {
				logger.Errorf("Failed to get filtered nfts: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
		} else {
			nfts, err = nftTokenRepository.GetNfts(collectionAddresses, walletAddress, offset, limit)
			if err != nil {
				logger.Errorf("Failed to get nfts: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}

			totalNfts, err := nftTokenRepository.GetNfts(collectionAddresses, walletAddress, nil, nil)
			if err != nil {
				logger.Errorf("Failed to get total nfts: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
			totalUnfiltered = len(totalNfts)
		}

		c.JSON(http.StatusOK, &entity.NftTokensResponse{
			Total: totalUnfiltered,
			Data:  nfts,
		})
	}
}

// NewAdminGetCollectionsHandler creates a handler to handle requests to get all collections
func NewAdminGetCollectionsHandler(nftCollectionRepository repository.NftCollectionRepository, client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get all collections here should be fine as we won't have even more than 100 collections for a while
		results, err := nftCollectionRepository.GetAllCollections()
		if err != nil {
			logger.Errorf("Failed to get collections: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

// NewAdminCreateOrUpdateNftHandler creates a handler to handle requests to create or update a nft token data
func NewAdminCreateOrUpdateNftHandler(nftTokenRepository repository.NftTokenRepository, client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("token id is invalid"))
			return
		}

		collectionAddress := c.DefaultQuery("collectionAddress", "")
		collectionAddress, err = utils.SanitizeAddress(collectionAddress)
		if err != nil {
			logger.Errorf("Invalid collection address: %s", collectionAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		address := common.HexToAddress(collectionAddress)
		contract, err := erc721.NewErc721(address, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		uri, err := contract.TokenURI(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		ownerAddress, err := contract.OwnerOf(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		walletAddress, err := utils.SanitizeAddress(ownerAddress.String())
		if err != nil {
			logger.Errorf("Invalid wallet address: %s", walletAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid wallet address"))
			return
		}

		nft, err := createNftRecord(uri, int(tokenId), walletAddress, collectionAddress)
		if err != nil {
			logger.Errorf("Failed to construct nft record: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		currentNft, err := nftTokenRepository.GetNftByTokenId(int(tokenId), collectionAddress)
		if err != nil {
			if err == gorm.ErrRecordNotFound { // nft token doesn't exist in DB. Create the nft record
				nftToken, err := nftTokenRepository.CreateNft(nft)
				if err != nil {
					logger.Errorf("Failed to create nft token db record: %v", err)
					c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
					return
				}
				c.JSON(http.StatusOK, nftToken)
				return
			}
			logger.Errorf("Failed to get nft record: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// nft token exists. Update the record
		nft.ID = currentNft.ID
		nftToken, err := nftTokenRepository.UpdateNft(nft)
		if err != nil {
			logger.Errorf("Failed to update nft token db record: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

func createNftRecord(uri string, tokenId int, ownerAddress string, collectionAddress string) (*entity.NftTokenModel, error) {
	nft := &entity.NftTokenModel{
		ID:                uuid.New().String(),
		TokenId:           tokenId,
		OwnerAddress:      &ownerAddress,
		CollectionAddress: collectionAddress,
	}

	if uri == "" {
		return nft, nil
	}

	splittedStr := strings.Split(uri, ",")
	if len(splittedStr) != 2 {
		return nil, fmt.Errorf("failed to split uri string to 2 parts. uri: %v", uri)
	}

	decodedMetadata, err := base64.StdEncoding.DecodeString(splittedStr[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode uri: %v", err)
	}

	var result entity.NftOnchainMeta
	err = json.Unmarshal(decodedMetadata, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %v", err)
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

// (Deprecated) NewAdminUpdateNftOwnerHandler creates a handler to handle requests to update nft owner of an nft
func NewAdminUpdateNftOwnerHandler(nftTokenRepository repository.NftTokenRepository, client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("token id is invalid"))
			return
		}

		collectionAddress := c.DefaultQuery("collectionAddress", "")
		collectionAddress, err = utils.SanitizeAddress(collectionAddress)
		if err != nil {
			logger.Errorf("Invalid collection address: %s", collectionAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		address := common.HexToAddress(collectionAddress)
		contract, err := erc721.NewErc721(address, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		ownerAddress, err := contract.OwnerOf(nil, big.NewInt(tokenId))
		if err != nil {
			logger.Errorf("Failed to query uri: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		walletAddress, err := utils.SanitizeAddress(ownerAddress.String())
		if err != nil {
			logger.Errorf("Invalid wallet address: %s", walletAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid wallet address"))
			return
		}

		nftToken, err := nftTokenRepository.UpdateNftOwner(int(tokenId), collectionAddress, walletAddress)
		if err != nil {
			logger.Errorf("Failed to update nft owner: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}

// NewAdminDeleteNftHandler creates a handler to handle requests for deleting a nft record from db
func NewAdminDeleteNftHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("token id is invalid"))
			return
		}

		collectionAddress := c.DefaultQuery("collectionAddress", "")
		collectionAddress, err = utils.SanitizeAddress(collectionAddress)
		if err != nil {
			logger.Errorf("Invalid collection address: %s", collectionAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		err = nftTokenRepository.DeleteNft(int(tokenId), collectionAddress)
		if err != nil {
			logger.Errorf("Failed to delete nft: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

// NewAdminUpdateCollectionAddressHandler creates a handler to handle requests for updating nft collection address
func NewAdminUpdateCollectionAddressHandler(
	nftCollectionRepository repository.NftCollectionRepository,
	franchiseCollectionRepository repository.FranchiseCollectionRepository,
	nftAllowlistRepository repository.NftAllowlistRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		oldAddress, err := utils.SanitizeAddress(c.Param("address"))
		if err != nil {
			logger.Errorf("Invalid old collection address: %s", oldAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid old collection address"))
			return
		}

		newAddress, err := utils.SanitizeAddress(c.DefaultQuery("newAddress", ""))
		if err != nil {
			logger.Errorf("Invalid new collection address: %s", newAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid new collection address"))
			return
		}

		err = nftCollectionRepository.UpdateCollectionAddress(oldAddress, newAddress)
		if err != nil {
			logger.Errorf("Failed to update collection address in nft collection: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		err = franchiseCollectionRepository.UpdateCollectionAddress(oldAddress, newAddress)
		if err != nil {
			logger.Errorf("Failed to update collection address in franchise collection: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		err = nftAllowlistRepository.UpdateCollectionAddress(oldAddress, newAddress)
		if err != nil {
			logger.Errorf("Failed to update collection address in nft allowlist: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
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
		collectionAddress, err := utils.SanitizeAddress(collectionAddress)
		if err != nil {
			return nil, fmt.Errorf("invalid collection address: %v", err)
		}
		collectionAddresses = append(collectionAddresses, collectionAddress)
	} else if franchiseId != "" {
		franchiseIdInt, err := strconv.ParseInt(franchiseId, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert franchise id: %v", err)
		}
		collectionAddresses, err = franchiseCollectionRepository.GetCollectionAddressesByFranchise(franchiseIdInt)
		if err != nil {
			return nil, fmt.Errorf("failed to get collection addresses by franchise id: %v", err)
		}
	} else {
		return nil, fmt.Errorf("neither collection address or franchise id is passed")
	}

	return collectionAddresses, nil
}
