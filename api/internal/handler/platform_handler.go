package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/s3"
	"github.com/project-nova/backend/proto/v1/web3_gateway"
)

type PlatformHandlerInterface interface {
	RequestFileUpload(c *gin.Context)
	ConfirmFileUpload(c *gin.Context)
}

type PlatformProtocolHandler struct {
	s3Client     s3.S3Client
	s3BucketName string
	web3Gateway  gateway.Web3GatewayClient
}

func NewPlatformProtocolHandler(s3Client s3.S3Client, s3FileUploadBucketName string, web3Gateway gateway.Web3GatewayClient) PlatformHandlerInterface {
	return &PlatformProtocolHandler{
		s3Client:     s3Client,
		s3BucketName: s3FileUploadBucketName,
		web3Gateway:  web3Gateway,
	}
}

func (ph *PlatformProtocolHandler) RequestFileUpload(c *gin.Context) {
	uuidString := uuid.New().String()
	signedUrl, err := ph.s3Client.RequestPreSignedUrl(ph.s3BucketName, uuidString)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &entity.FileUploadResp{
		URI: signedUrl,
	})
}

func (ph *PlatformProtocolHandler) ConfirmFileUpload(c *gin.Context) {
	var requestBody entity.UploadFileConfirmRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
		return
	}

	resp, err := ph.web3Gateway.UploadContent(&web3_gateway.UploadContentReq{
		Storage:  web3_gateway.StorageType_ARWEAVE,
		S3Bucket: ph.s3BucketName,
		S3Key:    requestBody.Filename,
	})

	if err != nil {
		logger.Errorf("Failed to upload content to web3-gateway: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("failed to upload content to web3-gateway"))
		return
	}

	c.JSON(http.StatusOK, &entity.FileUploadResp{
		URI: resp.ContentUrl,
	})
}
