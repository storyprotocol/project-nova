package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/pkg/s3"
)

type PlatformHandlerInterface interface {
	RequestFileUpload(c *gin.Context)
	ConfirmFileUpload(c *gin.Context)
}

type PlatformProtocolHandler struct {
	s3Client s3.S3Client
}

func NewPlatformProtocolHandler(s3Client s3.S3Client) PlatformHandlerInterface {
	return &PlatformProtocolHandler{
		s3Client: s3Client,
	}
}

func (ph *PlatformProtocolHandler) RequestFileUpload(c *gin.Context) {
	uuidString := uuid.New().String()
	signedUrl, err := ph.s3Client.RequestPreSignedUrl("story-file-upload", uuidString)
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
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
