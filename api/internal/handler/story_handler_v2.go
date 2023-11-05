package handler

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/gateway"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
	"github.com/project-nova/backend/proto/v1/web3_gateway"
)

// GET /story/:franchiseId/:storyId/:chapterId
func NewGetStoryContentHandlerV2(
	contentRepo repository.ProtocolStoryContentRepository,
	httpClient xhttp.Client,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. verify addresses
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseId"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("storyId"))
		if err != nil {
			logger.Errorf("Invalid story address: %s", c.Param("storyId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story address"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("chapterId"))
		if err != nil {
			logger.Errorf("Invalid chapter id: %s", c.Param("chapterId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid chapter id"))
			return
		}

		// 2. call db to get content uri
		content, err := contentRepo.GetContentByAddresses(franchiseAddress, collectionAddress, tokenId)
		if err != nil {
			logger.Errorf("Failed to read content from database: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 3. call the uri to get the content
		var result entity.ContentV2
		_, err = httpClient.Request("GET", *content.ContentUri, nil, &result)
		if err != nil {
			logger.Errorf("Failed to read content from remote storage: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// GET /story/:franchiseId/:storyId/:chapterId
func NewAdminUploadStoryContentHandlerV2(
	contentRepo repository.ProtocolStoryContentRepository,
	web3Gateway gateway.Web3GatewayClient,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. verify addresses
		franchiseAddress, err := utils.SanitizeAddress(c.Param("franchiseId"))
		if err != nil {
			logger.Errorf("Invalid franchise address: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise address"))
			return
		}

		collectionAddress, err := utils.SanitizeAddress(c.Param("storyId"))
		if err != nil {
			logger.Errorf("Invalid story address: %s", c.Param("storyId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story address"))
			return
		}

		tokenId, err := strconv.Atoi(c.Param("chapterId"))
		if err != nil {
			logger.Errorf("Invalid chapter id: %s", c.Param("chapterId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid chapter id"))
			return
		}

		var requestBody struct {
			Content string `json:"content"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
			return
		}

		// 2. call web3 gateway to upload content
		contentBase64 := base64.StdEncoding.EncodeToString([]byte(requestBody.Content))
		resp, err := web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
			Storage:     web3_gateway.StorageType_ARWEAVE,
			Content:     []byte(contentBase64),
			ContentType: "text/markdown",
			Tags: []*web3_gateway.Tag{
				{
					Name:  "franchise",
					Value: franchiseAddress,
				},
				{
					Name:  "story",
					Value: collectionAddress,
				},
				{
					Name:  "chapter",
					Value: strconv.Itoa(tokenId),
				},
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

		// 3. store content uri and other data in db
		err = contentRepo.CreateContent(&entity.ProtocolStoryContentModel{
			ID:                uuid.New().String(),
			FranchiseAddress:  &franchiseAddress,
			CollectionAddress: &collectionAddress,
			TokenId:           &tokenId,
			ContentJson:       requestBody.Content,
			ContentUri:        &resp.ContentUrl,
		})
		if err != nil {
			logger.Errorf("Failed to create content in the database: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"contentUrl": resp.ContentUrl,
		})
	}
}
