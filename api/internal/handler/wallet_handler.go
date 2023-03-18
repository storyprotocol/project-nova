package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/constant"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

// NewGetWalletProofHandler get wallet's proof
// Doc: To Be Added
func NewGetWalletProofHandler(walletMerkleProofRepo repository.WalletMerkleProofRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		allowlistId := c.DefaultQuery("allowlistId", "")
		address := c.Param("walletAddress")

		if !utils.IsValidAddress(address) {
			logger.Errorf("Invalid wallet address: %s", address)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid wallet addresss"})
			return
		}

		if !utils.IsValidUUID(allowlistId) {
			logger.Errorf("Invalid allowlist id: %s", allowlistId)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid allowlistId"})
			return
		}

		result, err := walletMerkleProofRepo.GetMerkleProof(address, allowlistId)
		if err == gorm.ErrRecordNotFound {
			logger.Errorf("Failed to get wallet merkle proof: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": constant.ErrorNotOnWhitelist})
			return
		}
		if err != nil {
			logger.Errorf("Failed to get wallet merkle proof: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"proof": result.Proof,
		})
	}
}

// NewGetWalletSignMessageHandler creates a handler to return wallet signing message
func NewGetWalletSignMessageHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		address := c.Param("walletAddress")
		if !utils.IsValidAddress(address) {
			logger.Errorf("Invalid wallet address: %s", address)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid wallet addresss"})
			return
		}

		nonce := xid.New().String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(constant.BackstorySignMessage, nonce),
		})
	}
}

func NewAdminCreateWalletProofHandler(walletMerkleProofRepo repository.WalletMerkleProofRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		allowlistId := c.DefaultQuery("allowlistId", "")
		proof := c.DefaultQuery("proof", "")
		address := c.Param("walletAddress")

		if !utils.IsValidAddress(address) {
			logger.Errorf("Invalid wallet address: %s", address)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid wallet addresss"})
			return
		}

		if !utils.IsValidUUID(allowlistId) {
			logger.Errorf("Invalid allowlist id: %s", allowlistId)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid allowlistId"})
			return
		}

		if proof == "" {
			logger.Error("Proof is not presented")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid proof"})
			return
		}

		decodedProof, err := base64.StdEncoding.DecodeString(proof)
		if err != nil {
			logger.Error("failed to decode proof")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid proof"})
			return
		}

		_, err = walletMerkleProofRepo.CreateMerkleProof(address, allowlistId, string(decodedProof))
		if err != nil {
			logger.Errorf("Failed to get wallet merkle proof: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("Successfully created the wallet proof for address: %s", address))
	}
}
