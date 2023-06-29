package handler

import (
	"encoding/base64"
	"encoding/json"
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

		c.JSON(http.StatusOK, &entity.CreateCharacterResp{
			MediaUri: resp.ContentUrl,
		})
	}
}
