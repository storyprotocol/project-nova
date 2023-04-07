package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/pkg/abi/character_registry"
	"github.com/project-nova/backend/pkg/abi/erc721"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
)

// GET /franchise
func NewGetFranchisesHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Franchises)
	}
}

// GET /franchise/:franchiseAddress
func NewGetFranchiseCollectionsHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.SingleFranchise)
	}
}

// GET /character/:franchiseAddress/:collectionAddress
func NewGetCharactersHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", franchiseAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("collectionAddress"))
		if err != nil {
			logger.Errorf("Invalid collection address: %s", collectionAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		registryAddress := common.HexToAddress(entity.FranchiseMap[franchiseAddress].CharacterRegistry)
		registryContract, err := character_registry.NewCharacterRegistry(registryAddress, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		collectionAddr := common.HexToAddress(collectionAddress)
		erc721Contract, err := erc721.NewErc721(collectionAddr, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		supply, err := erc721Contract.TotalSupply(nil)
		if err != nil {
			logger.Errorf("Failed to get total supply for collection %s: %v", collectionAddress, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		var resp []*entity.Character

		supplyInt := supply.Int64()
		for i := 0; i < int(supplyInt); i++ {
			charInfo, err := registryContract.Character(nil, collectionAddr, big.NewInt(int64(i)))
			if err != nil {
				logger.Errorf("Failed to get character info: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}

			uri, err := erc721Contract.TokenURI(nil, big.NewInt(int64(i)))
			if err != nil {
				logger.Errorf("Failed to get total uri for token %d: %v", i, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}

			ownerAddress, err := erc721Contract.OwnerOf(nil, big.NewInt(int64(i)))
			if err != nil {
				logger.Errorf("Failed to get owner address for token %d: %v", i, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}

			character, err := createCharacterReponse(uri, i, &charInfo, ownerAddress.String(), collectionAddress)
			if err != nil {
				logger.Errorf("Failed to create character response for %d: %v", i, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}

			resp = append(resp, character)
		}

		c.JSON(http.StatusOK, resp)
	}
}

// GET /character/:franchiseAddress/:collectionAddress/:tokenId
func NewGetCharacterHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.SingleCharacter)
	}
}

// GET /character/:franchiseAddress/:collectionAddress/:tokenId/collectors
func NewGetCollectorsHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Collectors)
	}
}

// Get /story/:franchiseAddress/:collectionAddress
func NewGetStoriesHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Stories)
	}
}

// Get /story/:franchiseAddress/:collectionAddress/:tokenId
func NewGetStoryHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.SingleStory)
	}
}

// GET
func NewGetDerivativesHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func createCharacterReponse(uri string, tokenId int, charInfo *character_registry.ICharacterRegistryCharacterInfo, ownerAddress string, collectionAddress string) (*entity.Character, error) {
	character := &entity.Character{
		TokenId:           tokenId,
		OwnerAddress:      ownerAddress,
		CollectionAddress: collectionAddress,
		Name:              charInfo.Name,
		Description:       charInfo.Description,
		AuthorAddress:     charInfo.Author.FeeCollector.String(),
	}

	if uri == "" {
		return character, nil
	}

	splittedStr := strings.Split(uri, ",")
	if len(splittedStr) != 2 {
		return nil, fmt.Errorf("failed to split uri string to 2 parts. uri: %v", uri)
	}

	decodedMetadata, err := base64.StdEncoding.DecodeString(splittedStr[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode uri: %v", err)
	}

	var result entity.CharacterNftOnchainMeta
	err = json.Unmarshal(decodedMetadata, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %v", err)
	}

	if result.Image != nil {
		character.ImageUrl = *result.Image
	}

	if result.Attributes != nil {
		for _, attr := range result.Attributes {
			character.Traits = append(character.Traits, &entity.NftTraitResponse{
				TraitType: attr.TraitType,
				Value:     attr.Value,
			})
		}
	}

	return character, nil
}
