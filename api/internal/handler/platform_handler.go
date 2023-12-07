package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/s3"
	"github.com/project-nova/backend/pkg/utils"
	"github.com/project-nova/backend/proto/v1/web3_gateway"
	"gorm.io/gorm"
)

const (
	walletSignInMessage = "Welcome to Story Protocol!\n\nClick to sign in and accept the Story Protocol Terms of Service (https://storyprotocol.xyz/tos).\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nWallet address:\n%s\n\nNonce:\n%s"
)

type PlatformHandlerInterface interface {
	RequestFileUpload(c *gin.Context)
	ConfirmFileUpload(c *gin.Context)
	RequestWalletSignIn(c *gin.Context)
	VerifyWalletSignIn(c *gin.Context)
}

type PlatformProtocolHandler struct {
	s3Client                 s3.S3Client
	s3BucketName             string
	web3Gateway              gateway.Web3GatewayClient
	walletSignInfoRepository repository.WalletSignInfoRepository
}

func NewPlatformProtocolHandler(
	s3Client s3.S3Client,
	s3FileUploadBucketName string,
	web3Gateway gateway.Web3GatewayClient,
	walletSignInfoRepository repository.WalletSignInfoRepository,
) PlatformHandlerInterface {
	return &PlatformProtocolHandler{
		s3Client:                 s3Client,
		s3BucketName:             s3FileUploadBucketName,
		web3Gateway:              web3Gateway,
		walletSignInfoRepository: walletSignInfoRepository,
	}
}

func (ph *PlatformProtocolHandler) RequestFileUpload(c *gin.Context) {
	uuidString := uuid.New().String()
	signedUrl, err := ph.s3Client.RequestPreSignedUrl(ph.s3BucketName, uuidString)
	if err != nil {
		logger.Errorf("Failed to request pre-signed url: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage("failed to request pre-signed url"))
		return
	}

	c.JSON(http.StatusOK, &entity.UploadFileRequestResp{
		Url: signedUrl,
		Key: uuidString,
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
		S3Key:    requestBody.Key,
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

func (ph *PlatformProtocolHandler) RequestWalletSignIn(c *gin.Context) {
	var requestBody entity.SignInWalletRequest
	if err := c.BindQuery(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
		return
	}

	if !utils.IsValidAddress(requestBody.WalletAddress) {
		logger.Errorf("Invalid wallet address: %s", requestBody.WalletAddress)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid wallet address"))
		return
	}

	signInfo, err := ph.walletSignInfoRepository.GetWalletNonce(requestBody.WalletAddress)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			nonce := uuid.New().String()
			signInfo, err = ph.walletSignInfoRepository.CreateNewWalletNonce(requestBody.WalletAddress, nonce)
			if err != nil {
				logger.Errorf("Failed to create new wallet nonce: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("failed to create new wallet nonce"))
				return
			}
		} else {
			logger.Errorf("Failed to get wallet nonce: %v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage("failed to get wallet nonce"))
			return
		}
	}

	message := fmt.Sprintf(walletSignInMessage, strings.ToLower(requestBody.WalletAddress), signInfo.Nonce)
	c.JSON(http.StatusOK, &entity.SigninWalletResponse{
		SigningMessage: message,
	})
}

func (ph *PlatformProtocolHandler) VerifyWalletSignIn(c *gin.Context) {
	var requestBody entity.VerifyWalletSignInRequest
	if err := c.BindJSON(&requestBody); err != nil {
		logger.Errorf("Failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid request body"))
		return
	}

	if !utils.IsValidAddress(requestBody.WalletAddress) {
		logger.Errorf("Invalid wallet address: %s", requestBody.WalletAddress)
		c.JSON(http.StatusBadRequest, ErrorMessage("invalid wallet address"))
		return
	}

	signInfo, err := ph.walletSignInfoRepository.GetWalletNonce(requestBody.WalletAddress)
	if err != nil {
		logger.Errorf("Failed to get wallet nonce: %v", err)
		c.JSON(http.StatusBadRequest, ErrorMessage("failed to get wallet nonce"))
		return
	}

	originalMessage := fmt.Sprintf(walletSignInMessage, strings.ToLower(requestBody.WalletAddress), signInfo.Nonce)
	// Recover the public key from the signature
	publicKey, err := crypto.SigToPub([]byte(originalMessage), []byte(requestBody.Signature))
	if err != nil {
		logger.Errorf("Failed to recover public key: %v", err)
		c.JSON(http.StatusUnauthorized, ErrorMessage("Signature verification failed"))
		return
	}

	recoveredAddress := crypto.PubkeyToAddress(*publicKey)

	// compare the address
	if recoveredAddress.Hex() != signInfo.WalletAddress {
		c.JSON(http.StatusUnauthorized, ErrorMessage("Signature verification failed"))
		return
	}

	c.JSON(http.StatusOK, nil)
}
