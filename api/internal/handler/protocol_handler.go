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
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/pkg/abi/character_registry"
	"github.com/project-nova/backend/pkg/abi/erc721"
	"github.com/project-nova/backend/pkg/abi/franchise"
	"github.com/project-nova/backend/pkg/abi/story_registry"
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
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		franchise, ok := entity.FranchiseMap[franchiseAddress]
		if !ok {
			logger.Errorf("Unkown franchise address: %s", franchiseAddress)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		c.JSON(http.StatusOK, franchise)
	}
}

// GET /character/:franchiseAddress/:collectionAddress
func NewGetCharactersHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("collectionAddress"))
		if err != nil {
			logger.Errorf("Invalid collection address: %s", c.Param("collectionAddress"))
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
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("collectionAddress"))
		if err != nil {
			logger.Errorf("Invalid collection address: %s", c.Param("collectionAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("tokenId"))
		if err != nil {
			logger.Errorf("Invalid token id: %s", c.Param("tokenId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid token id"))
			return
		}

		registryAddress := common.HexToAddress(entity.FranchiseMap[franchiseAddress].CharacterRegistry)
		registryContract, err := character_registry.NewCharacterRegistry(registryAddress, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the character registry contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		collectionAddr := common.HexToAddress(collectionAddress)
		erc721Contract, err := erc721.NewErc721(collectionAddr, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the character collection contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		charInfo, err := registryContract.Character(nil, collectionAddr, big.NewInt(int64(tokenId)))
		if err != nil {
			logger.Errorf("Failed to get character info: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		uri, err := erc721Contract.TokenURI(nil, big.NewInt(int64(tokenId)))
		if err != nil {
			logger.Errorf("Failed to get total uri for token %d: %v", tokenId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		ownerAddress, err := erc721Contract.OwnerOf(nil, big.NewInt(int64(tokenId)))
		if err != nil {
			logger.Errorf("Failed to get owner address for token %d: %v", tokenId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		character, err := createCharacterReponse(uri, tokenId, &charInfo, ownerAddress.String(), collectionAddress)
		if err != nil {
			logger.Errorf("Failed to create character response for %d: %v", tokenId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, character)
	}
}

// GET /character/:franchiseAddress/:collectionAddress/:tokenId/collectors
func NewGetCollectorsHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("collectionAddress"))
		if err != nil {
			logger.Errorf("Invalid collection address: %s", c.Param("collectionAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}
		collectionAddr := common.HexToAddress(collectionAddress)

		tokenId, err := strconv.Atoi(c.Param("tokenId"))
		if err != nil {
			logger.Errorf("Invalid token id: %s", c.Param("tokenId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid token id"))
			return
		}

		franchiseAddr := common.HexToAddress(franchiseAddress)
		franchiseContract, err := franchise.NewFranchise(franchiseAddr, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the franchise contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		var collectors []common.Address
		_, isCharacter := entity.CharacterContractMap[collectionAddress]
		if isCharacter {
			collectors, err = franchiseContract.GetCharacterCollectors(nil, collectionAddr, big.NewInt(int64(tokenId)))
			if err != nil {
				logger.Errorf("Failed to get character collectors for id %d: %v", tokenId, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
		} else if _, isStory := entity.StoryContractMap[collectionAddress]; isStory {
			collectors, err = franchiseContract.GetStoryCollectors(nil, collectionAddr, big.NewInt(int64(tokenId)))
			if err != nil {
				logger.Errorf("Failed to get story collectors for id %d: %v", tokenId, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
		} else {
			logger.Errorf("Invalid input. franchise %s, collection %s, id %d", franchiseAddress, collectionAddress, tokenId)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Invalid input"))
			return
		}

		var collectorsResp []*entity.Collector
		for _, c := range collectors {
			collectorsResp = append(collectorsResp, &entity.Collector{
				Address: c.String(),
			})
		}

		c.JSON(http.StatusOK, collectorsResp)
	}
}

// Get /story/:franchiseAddress/:collectionAddress
func NewGetStoriesHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("collectionAddress"))
		if err != nil {
			logger.Errorf("Invalid collection address: %s", c.Param("collectionAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		registryAddress := common.HexToAddress(entity.FranchiseMap[franchiseAddress].StoryRegistry)
		registryContract, err := story_registry.NewStoryRegistry(registryAddress, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the story registry contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		collectionAddr := common.HexToAddress(collectionAddress)
		erc721Contract, err := erc721.NewErc721(collectionAddr, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the story collection contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		supply, err := erc721Contract.TotalSupply(nil)
		if err != nil {
			logger.Errorf("Failed to get total supply for collection %s: %v", collectionAddress, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		collectionInfo, ok := entity.StoryContractMap[collectionAddress]
		if !ok {
			logger.Errorf("Failed to get collection info for collection %s: %v", collectionAddress, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		var resp []*entity.Story
		logger.Infof("total supply: %d", supply)

		supplyInt := supply.Int64()
		for i := 1; i < int(supplyInt); i++ {
			storyInfo, err := registryContract.Story(nil, collectionAddr, big.NewInt(int64(i)))
			if err != nil {
				logger.Errorf("Failed to get story info for token %d: %v", i, err)
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

			story, err := createStoryReponse(uri, i, &storyInfo, collectionInfo, ownerAddress.String(), collectionAddress)
			if err != nil {
				logger.Errorf("Failed to create story response for %d: %v", i, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}

			resp = append(resp, story)
		}
		c.JSON(http.StatusOK, resp)
	}
}

// Get /story/:franchiseAddress/:collectionAddress/:tokenId
func NewGetStoryHandler(client *ethclient.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseAddress"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("collectionAddress"))
		if err != nil {
			logger.Errorf("Invalid collection address: %s", c.Param("collectionAddress"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid collection address"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("tokenId"))
		if err != nil {
			logger.Errorf("Invalid token id: %s", c.Param("tokenId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid token id"))
			return
		}

		registryAddress := common.HexToAddress(entity.FranchiseMap[franchiseAddress].StoryRegistry)
		registryContract, err := story_registry.NewStoryRegistry(registryAddress, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the story registry contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		collectionAddr := common.HexToAddress(collectionAddress)
		erc721Contract, err := erc721.NewErc721(collectionAddr, client)
		if err != nil {
			logger.Errorf("Failed to instantiate the story collection contract: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		collectionInfo, ok := entity.StoryContractMap[collectionAddress]
		if !ok {
			logger.Errorf("Failed to get collection info for collection %s: %v", collectionAddress, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		storyInfo, err := registryContract.Story(nil, collectionAddr, big.NewInt(int64(tokenId)))
		if err != nil {
			logger.Errorf("Failed to get story info: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		uri, err := erc721Contract.TokenURI(nil, big.NewInt(int64(tokenId)))
		if err != nil {
			logger.Errorf("Failed to get total uri for token %d: %v", tokenId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		ownerAddress, err := erc721Contract.OwnerOf(nil, big.NewInt(int64(tokenId)))
		if err != nil {
			logger.Errorf("Failed to get owner address for token %d: %v", tokenId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		story, err := createStoryReponse(uri, tokenId, &storyInfo, collectionInfo, ownerAddress.String(), collectionAddress)
		if err != nil {
			logger.Errorf("Failed to create story response for %d: %v", tokenId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, story)
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

func createStoryReponse(uri string, tokenId int, storyInfo *story_registry.StoryInfo, collectionInfo *entity.StoryCollection, ownerAddress string, collectionAddress string) (*entity.Story, error) {
	story := &entity.Story{
		TokenId:           tokenId,
		OwnerAddress:      ownerAddress,
		CollectionAddress: collectionAddress,
		Title:             storyInfo.Title,
		IsCanon:           collectionInfo.IsCanon,
	}

	for _, author := range storyInfo.Author {
		story.AuthorAddress = append(story.AuthorAddress, author.String())
	}

	for _, character := range storyInfo.Characters {
		story.Characters = append(story.Characters, &entity.Character{
			CollectionAddress: character.Collection.String(),
			TokenId:           int(character.TokenId.Int64()),
		})
	}

	if uri == "" {
		return story, nil
	}

	splittedStr := strings.Split(uri, ",")
	if len(splittedStr) != 2 {
		return nil, fmt.Errorf("failed to split uri string to 2 parts. uri: %v", uri)
	}

	decodedMetadata, err := base64.StdEncoding.DecodeString(splittedStr[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode uri: %v", err)
	}

	var result entity.StoryNftOnchainMeta
	err = json.Unmarshal(decodedMetadata, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %v", err)
	}

	if result.Image != nil {
		story.ImageUrl = *result.Image
	}
	if result.Description != nil {
		story.ContentUrl = *result.Description
	}

	return story, nil
}
