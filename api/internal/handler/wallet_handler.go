package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/logger"
	"gorm.io/gorm"
)

type WalletMerkleProofModel struct {
	ID            string `gorm:"primaryKey;column:id"`
	AllowlistId   string
	WalletAddress string
	Proof         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (WalletMerkleProofModel) TableName() string {
	return "wallet_merkle_proof"
}

func NewGetWalletProofHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		allowlistId := c.DefaultQuery("allowlistId", "")
		address := c.Param("walletAddress")

		// Validate address
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		result := &WalletMerkleProofModel{}
		r := db.Where("address = ? and allowlist_id = ?", address, allowlistId).First(&result)
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		if r.Error != nil {
			logger.Errorf("Failed to query db: %v", r.Error)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"proof": result.Proof,
		})

	}
}

func GetWalletNftsHandler(c *gin.Context) {

}
