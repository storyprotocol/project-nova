package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/logger"
)

func NewGetWalletProofHandler(walletMerkleProofRepo repository.WalletMerkleProofRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		allowlistId := c.DefaultQuery("allowlistId", "")
		address := c.Param("walletAddress")

		// Validate address
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		result, err := walletMerkleProofRepo.GetMerkleProof(address, allowlistId)
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

func NewGetWalletNftsHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		address := c.Param("walletAddress")
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert franchise id: %v", err)
			c.String(http.StatusBadRequest, "franchise id is invalid")
			return
		}

		result, err := nftTokenRepository.GetNftsByOwner(franchiseId, address)
		if err != nil {
			logger.Errorf("Failed to get nfts by owner: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
