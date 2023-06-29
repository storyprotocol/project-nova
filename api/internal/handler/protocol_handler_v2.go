package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/proto/v1/web3_gateway"
)

// GET /character/:franchiseId
func NewGetCharactersHandlerV2(
	characterInfoRepository repository.CharacterInfoRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		characters, err := characterInfoRepository.GetCharacters(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get characters: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		resp := []*entity.GetCharactersResp{}
		for _, character := range characters {
			charResp, err := character.ToGetCharactersResp()
			if err != nil {
				logger.Errorf("Failed to convert to character resp: %v", err)
				continue
			}
			resp = append(resp, charResp)
		}

		c.JSON(http.StatusOK, resp)
	}
}

// GET /character/:franchiseId/:tokenId
func NewGetCharacterHandlerV2(
	characterInfoRepository repository.CharacterInfoRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("tokenId"))
		if err != nil {
			logger.Errorf("Invalid token id: %s", c.Param("tokenId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid token id"))
			return
		}

		character, err := characterInfoRepository.GetCharacter(franchiseId, int64(tokenId))
		if err != nil {
			logger.Errorf("Failed to get character, id %s: %v", c.Param("tokenId"), err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		characterResp := character.ToGetCharacterResp()
		c.JSON(http.StatusOK, characterResp)
	}
}

// POST /character/:franchiseId
func NewCreateCharacterHandlerV2(
	characterInfoRepository repository.CharacterInfoRepository,
	storyInfoV2Repository repository.StoryInfoV2Repository,
	web3Gateway gateway.Web3GatewayClient,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		var requestBody entity.CreateCharacterRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}
		if err := requestBody.Validate(); err != nil {
			logger.Errorf("Failed to validate request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		characterInfo := requestBody.ToCharacterInfoModel()
		characterInfo.FranchiseId = franchiseId

		storyMeta := &entity.CharacterMetadata{
			Name:      characterInfo.CharacterName,
			Backstory: *characterInfo.Backstory,
		}
		storyMetaBytes, err := json.Marshal(storyMeta)
		if err != nil {
			logger.Errorf("Failed to marshal the story meta: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		contentBase64 := base64.StdEncoding.EncodeToString(storyMetaBytes)
		resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
			Storage:     web3_gateway.StorageType_ARWEAVE,
			Content:     []byte(contentBase64),
			ContentType: web3_gateway.ContentType_MARKDOWN,
			Tags: []*web3_gateway.Tag{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
			},
		})
		if err != nil {
			logger.Errorf("Failed to upload content to web3-gateway: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		characterInfo.MediaUri = &resp.ContentUrl
		err = characterInfoRepository.CreateCharacter(characterInfo)
		if err != nil {
			logger.Errorf("Failed to create character info in db: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		createStoryReq := &entity.CreateStoryRequestBody{
			OwnerAddress: &characterInfo.OwnerAddress,
			Content:      characterInfo.Backstory,
		}
		storyMediaUrl, err := createStoryV2(franchiseId, createStoryReq, storyInfoV2Repository, web3Gateway)
		if err != nil {
			logger.Errorf("Failed to create story: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, &entity.CreateCharacterResp{
			CharacterMediaUri: resp.ContentUrl,
			BackstoryMediaUri: *storyMediaUrl,
		})
	}
}

// GET /story/:franchiseId
func NewGetStoriesHandlerV2(
	storyInfoV2Repository repository.StoryInfoV2Repository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		stories, err := storyInfoV2Repository.GetStories(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get stories: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		resp := []*entity.GetStoryResp{}
		for _, story := range stories {
			storyResp, err := story.ToGetStoryResp()
			if err != nil {
				logger.Errorf("Failed to convert to story resp: %v", err)
				continue
			}
			resp = append(resp, storyResp)
		}

		c.JSON(http.StatusOK, resp)
	}
}

// GET /story/:franchiseId/:tokenId
func NewGetStoryHandlerV2(
	storyInfoV2Repository repository.StoryInfoV2Repository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("tokenId"))
		if err != nil {
			logger.Errorf("Invalid token id: %s", c.Param("tokenId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid token id"))
			return
		}

		story, err := storyInfoV2Repository.GetStory(franchiseId, int64(tokenId))
		if err != nil {
			logger.Errorf("Failed to get story, id %s: %v", c.Param("tokenId"), err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		storyResp, err := story.ToGetStoryDetailsResp()
		if err != nil {
			logger.Errorf("Failed to convert to story details, id %s: %v", c.Param("tokenId"), err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, storyResp)
	}
}

// POST /story/:franchiseId
func NewCreateStoryHandlerV2(
	storyInfoV2Repository repository.StoryInfoV2Repository,
	web3Gateway gateway.Web3GatewayClient,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		var requestBody entity.CreateStoryRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}
		if err := requestBody.Validate(); err != nil {
			logger.Errorf("Failed to validate request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		contentUrl, err := createStoryV2(franchiseId, &requestBody, storyInfoV2Repository, web3Gateway)
		if err != nil {
			logger.Errorf("Failed to create story: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, &entity.CreateStoryResp{
			MediaUri: *contentUrl,
		})
	}
}

func createStoryV2(
	franchiseId int64,
	request *entity.CreateStoryRequestBody,
	storyInfoV2Repository repository.StoryInfoV2Repository,
	web3Gateway gateway.Web3GatewayClient,
) (*string, error) {
	storyInfoV2 := request.ToStoryInfoV2Model()
	storyInfoV2.FranchiseId = franchiseId

	storyMeta := &entity.StoryMetadata{
		Content: *storyInfoV2.Content,
	}
	storyMetaBytes, err := json.Marshal(storyMeta)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the story meta: %v", err)
	}

	contentBase64 := base64.StdEncoding.EncodeToString(storyMetaBytes)
	resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
		Storage:     web3_gateway.StorageType_ARWEAVE,
		Content:     []byte(contentBase64),
		ContentType: web3_gateway.ContentType_MARKDOWN,
		Tags: []*web3_gateway.Tag{
			{
				Name:  "Content-Type",
				Value: "application/json",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload content to web3-gateway: %v", err)
	}

	storyInfoV2.MediaUri = &resp.ContentUrl
	err = storyInfoV2Repository.CreateStory(storyInfoV2)
	if err != nil {
		return nil, fmt.Errorf("failed to create story info v2 in db: %v", err)
	}

	return &resp.ContentUrl, nil
}
