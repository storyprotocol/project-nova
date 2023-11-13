package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/abi/story_blocks_registry"
	"github.com/project-nova/backend/pkg/gateway"
	xhttp "github.com/project-nova/backend/pkg/http"
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

		characterMetadata := requestBody.ToCharacterMetadata()

		characterMetaBytes, err := json.Marshal(characterMetadata)
		if err != nil {
			logger.Errorf("Failed to marshal the character metadata: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		contentBase64 := base64.StdEncoding.EncodeToString(characterMetaBytes)
		resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
			Storage:     web3_gateway.StorageType_ARWEAVE,
			Content:     []byte(contentBase64),
			ContentType: "text/markdown",
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

		createStoryReq := &entity.CreateStoryRequestBody{
			OwnerAddress: requestBody.OwnerAddress,
			Content:      requestBody.Backstory,
		}
		storyMediaUrl, err := createStoryV2(franchiseId, createStoryReq, web3Gateway)
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

// POST /files/upload
func NewUploadFileHandlerV2(
	web3Gateway gateway.Web3GatewayClient,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody entity.UploadFileRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		// decode base64 string into byte array
		decoded, err := base64.StdEncoding.DecodeString(requestBody.Base64)
		if err != nil {
			logger.Errorf("Failed to decode base64 string: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid base64 string"))
			return
		}

		resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
			Storage:     web3_gateway.StorageType_ARWEAVE,
			Content:     decoded,
			ContentType: requestBody.ContentType,
		})
		if err != nil {
			logger.Errorf("Failed to upload content to web3-gateway: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		contentUrl := []string{
			resp.ContentUrl,
		}

		c.JSON(http.StatusOK, &entity.FileUploadResp{
			URIs: contentUrl,
		})

		// TODO(Rex): Re-enable until S3 permission is fixed
		// headers := c.Request.Header
		// logger.Infof("headers: %v", headers)
		// form, err := c.MultipartForm()
		// if err != nil {
		// 	logger.Errorf("Failed to read multipart form: %v", err)
		// 	c.JSON(http.StatusBadRequest, ErrorMessage("Invalid request"))
		// 	return
		// }

		// files := form.File["file[]"]
		// logger.Infof("files: %v", files)

		// contentUrl := []string{}
		// for _, file := range files {
		// 	contentType := c.PostForm("content_type")
		// 	// Get the file bytes in memory
		// 	f, _ := file.Open()
		// 	fileBytes, _ := ioutil.ReadAll(f)

		// 	resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
		// 		Storage:     web3_gateway.StorageType_ARWEAVE,
		// 		Content:     fileBytes,
		// 		ContentType: contentType,
		// 	})
		// 	if err != nil {
		// 		logger.Errorf("Failed to upload content to web3-gateway: %v", err)
		// 		contentUrl = append(contentUrl, "")
		// 		continue
		// 	}
		// 	contentUrl = append(contentUrl, resp.ContentUrl)
		// 	logger.Infof("file size: %d, content-type: %s", len(fileBytes), contentType)
		// }

		// c.JSON(http.StatusOK, &entity.FileUploadResp{
		// 	URIs: contentUrl,
		// })
	}
}

// POST /story/:franchiseId
func NewCreateStoryHandlerV2(
	web3Gateway gateway.Web3GatewayClient,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		logger.Infof(">>> franchiseId: %d", franchiseId)
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

		contentUrl, err := createStoryV2(franchiseId, &requestBody, web3Gateway)
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

// POST /character/:franchiseId/:characterId/:storyId
func NewAdminCreateCharacterWithBackstoryHandler(
	characterInfoRepository repository.CharacterInfoRepository,
	storyInfoV2Repository repository.StoryInfoV2Repository,
	web3Gateway gateway.Web3GatewayClient,
	httpClient xhttp.Client,
	storyBlocksRegistry *story_blocks_registry.StoryBlocksRegistry,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		characterId, err := strconv.ParseInt(c.Param("characterId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid character id: %s", c.Param("characterId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid character id"))
			return
		}

		storyId, err := strconv.ParseInt(c.Param("storyId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid story id: %s", c.Param("storyId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story id"))
			return
		}

		var requestBody gateway.CreateCharacterWithBackstoryRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		// 1. Get story data from protocol
		storyOwner, err := storyBlocksRegistry.OwnerOf(nil, big.NewInt(storyId))
		if err != nil {
			logger.Errorf("Failed to create get the owner of the story %d: %v", storyId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		storyOwnerStr := storyOwner.String()

		storyBlock, err := storyBlocksRegistry.ReadStoryBlock(nil, big.NewInt(storyId))
		if err != nil {
			logger.Errorf("Failed to create get the story block for the story %d: %v", storyId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		if storyBlock.BlockType != 1 {
			logger.Errorf("Invalid block type for story. type: %d, id: %s", storyBlock.BlockType, storyId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story id"))
			return
		}

		storyInfo := &entity.StoryInfoV2Model{
			ID:               uuid.New().String(),
			FranchiseId:      franchiseId,
			StoryId:          &storyId,
			StoryName:        storyBlock.Name,
			StoryDescription: &storyBlock.Description,
			OwnerAddress:     &storyOwnerStr,
			MediaUri:         &storyBlock.MediaUrl,
			Txhash:           &requestBody.TxHash,
		}

		// 2. Use the mediaURL to fetch story content from Arweave
		var storyMetaData entity.StoryMetadata
		_, err = httpClient.Request("GET", storyBlock.MediaUrl, nil, &storyMetaData)
		if err != nil {
			logger.Errorf("Failed to get story metadata from remote storage. url: %s, error: %v", storyBlock.MediaUrl, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 3. Populate all story data to DB
		storyInfo.Content = &storyMetaData.Content
		err = storyInfoV2Repository.CreateStory(storyInfo)
		if err != nil {
			logger.Errorf("Failed to create the story: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 4. Do the same for character
		characterOwner, err := storyBlocksRegistry.OwnerOf(nil, big.NewInt(characterId))
		if err != nil {
			logger.Errorf("Failed to create get the owner of the character %d: %v", characterId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		characterOwnerStr := characterOwner.String()

		characterBlock, err := storyBlocksRegistry.ReadStoryBlock(nil, big.NewInt(characterId))
		if err != nil {
			logger.Errorf("Failed to create get the story block for the character %d: %v", characterId, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		if characterBlock.BlockType != 2 {
			logger.Errorf("Invalid block type for character. type: %d, id: %s", storyBlock.BlockType, characterId)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid character id"))
			return
		}

		characterInfo := &entity.CharacterInfoModel{
			ID:            uuid.New().String(),
			FranchiseId:   franchiseId,
			CharacterId:   &characterId,
			CharacterName: characterBlock.Name,
			OwnerAddress:  characterOwnerStr,
			Backstory:     storyInfo.Content,
			MediaUri:      &characterBlock.MediaUrl,
			Txhash:        &requestBody.TxHash,
		}

		var characterMetaData entity.CharacterMetadata
		_, err = httpClient.Request("GET", characterBlock.MediaUrl, nil, &characterMetaData)
		if err != nil {
			logger.Errorf("Failed to get character metadata from remote storage. url: %s, error: %s", characterBlock.MediaUrl, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		if characterMetaData.ImageUrl != nil {
			characterInfo.ImageUrl = characterMetaData.ImageUrl
		}

		err = characterInfoRepository.CreateCharacter(characterInfo)
		if err != nil {
			logger.Errorf("Failed to create the character: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

func createStoryV2(
	franchiseId int64,
	request *entity.CreateStoryRequestBody,
	web3Gateway gateway.Web3GatewayClient,
) (*string, error) {
	storyMeta := &entity.StoryMetadata{
		Content: *request.Content,
	}
	if len(request.Characters) > 0 {
		storyMeta.Characters = request.Characters
	}
	storyMetaBytes, err := json.Marshal(storyMeta)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the story meta: %v", err)
	}

	resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
		Storage:     web3_gateway.StorageType_ARWEAVE,
		Content:     storyMetaBytes,
		ContentType: "application/json",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload content to web3-gateway: %v", err)
	}

	return &resp.ContentUrl, nil
}
