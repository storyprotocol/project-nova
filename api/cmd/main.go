package main

import (
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/database"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

type WhitelistWallet struct {
	ID          string `gorm:"primaryKey;column:id"`
	Address     *string
	MerkleProof *string
	CreatedAt   time.Time
}

type StoryNFT struct {
	Name   string
	Symbol string
	Count  uint64
	URI    string
}

func (WhitelistWallet) TableName() string {
	return "whitelist_wallet"
}

func main() {
	r := gin.Default()

	r.GET("/mint/proof", func(c *gin.Context) {
		address := c.DefaultQuery("address", "")

		// Validate address
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		db, err := database.NewGormDB("postgresql://postgres:@localhost:23005/postgres?sslmode=disable")
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		result := &WhitelistWallet{}
		r := db.Where("address = ?", address).First(&result)
		if r.Error != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("address is invalid: %v", r.Error))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"proof": *result.MerkleProof,
		})
	})

	r.GET("/nfts", func(c *gin.Context) {
		privateKey := "a6cd3f393b1cddf8be66e2ff784640adbafbce852267ec1ec000000000000000" // Fake Key
		contractAddress := "0x64432E5A76a93e79be2f7F3F12982059a32Fd794"
		address := c.DefaultQuery("address", "")
		// Validate address
		if address == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("input address is invalid, address: %s", address))
			return
		}

		sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
			PrivateKey: privateKey,
		})
		if err != nil {
			logger.Errorf("Failed to set up sdk: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		content, err := os.ReadFile("./resource/abi/story_pass.json")
		if err != nil {
			logger.Errorf("Failed to read JSON files: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		abi := string(content)
		contract, err := sdk.GetContractFromAbi(contractAddress, abi)
		if err != nil {
			logger.Errorf("Failed to get contract: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balance, err := contract.Call(c, "balanceOf", address)
		if err != nil {
			logger.Errorf("Failed to get balance: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balanceBigInt, ok := balance.(*big.Int)
		if !ok {
			logger.Errorf("Failed to convert balance\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		var response []*StoryNFT

		if balanceBigInt.Uint64() > 0 {
			name, err := contract.Call(c, "name")
			if err != nil {
				logger.Errorf("Failed to get name: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			nameStr, ok := name.(string)
			if !ok {
				logger.Errorf("Failed to convert name to string\n")
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			symbol, err := contract.Call(c, "symbol")
			if err != nil {
				logger.Errorf("Failed to get symbol: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			symbolStr, ok := symbol.(string)
			if !ok {
				logger.Errorf("Failed to convert symbol to string\n")
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			tokenID, err := contract.Call(c, "tokenOfOwnerByIndex", address, 0)
			if err != nil {
				logger.Errorf("Failed to get tokenID: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			tokenURI, err := contract.Call(c, "tokenURI", tokenID)
			if err != nil {
				logger.Errorf("Failed to get tokenURI: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			tokenURIStr, ok := tokenURI.(string)
			if !ok {
				logger.Errorf("Failed to convert tokenURI to string\n")
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			response = append(response, &StoryNFT{
				Name:   nameStr,
				Symbol: symbolStr,
				Count:  balanceBigInt.Uint64(),
				URI:    tokenURIStr,
			})
		}

		c.JSON(http.StatusOK, response)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
