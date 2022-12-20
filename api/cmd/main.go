package main

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/database"
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
		privateKey := "a6cd3f393b1cddf8be66e2ff784640adbafbce852267ec1ec0e22eb741232000"

		sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
			PrivateKey: privateKey,
		})
		if err != nil {
			fmt.Printf("Failed to set up sdk: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		// You can replace your own contract address here
		contractAddress := "0xbcF76d7B52D6edef3D6Ec009D0371F06895B4E6A"

		// Add your contract ABI here
		abi := `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"claimed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"components":[{"internalType":"string","name":"passId","type":"string"},{"internalType":"string","name":"grade","type":"string"},{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"background","type":"string"},{"internalType":"string","name":"badges","type":"string"},{"internalType":"string","name":"offChainData","type":"string"}],"internalType":"struct StoryPass.TokenURIParams","name":"params","type":"tuple"}],"name":"constructTokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"grades","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"merkleRoot","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32[]","name":"merkleProof","type":"bytes32[]"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"mintPass","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"offChainDS","outputs":[{"internalType":"contract IOffChainDataSource","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"nft","type":"address"},{"internalType":"string","name":"badge","type":"string"}],"name":"register","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"registeredBadges","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"registeredNfts","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"resetRegisteredNfts","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId_","type":"uint256"},{"internalType":"uint256","name":"grade_","type":"uint256"}],"name":"setGrade","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"merkleRoot_","type":"bytes32"}],"name":"setMerkleRoot","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"contract IOffChainDataSource","name":"offChainDS_","type":"address"}],"name":"setOffChainDataSource","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"tokenGrades","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId_","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
		// Now you have a contract instance ready to go
		contract, err := sdk.GetContractFromAbi(contractAddress, abi)
		if err != nil {
			fmt.Printf("Failed to get contract: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balance, err := contract.Call(c, "balanceOf", "0xA8A97aBb6ABaD0c04321bF2afA1a2E2639b371e7")
		if err != nil {
			fmt.Printf("Failed to get balance: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balanceBigInt, ok := balance.(*big.Int)
		if !ok {
			fmt.Print("Failed to convert balance\n")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		var response []*StoryNFT

		if balanceBigInt.Uint64() > 0 {
			name, err := contract.Call(c, "name")
			if err != nil {
				fmt.Printf("Failed to get name: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			nameStr, ok := name.(string)
			if !ok {
				fmt.Print("Failed to convert name to string\n")
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			symbol, err := contract.Call(c, "symbol")
			if err != nil {
				fmt.Printf("Failed to get symbol: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			symbolStr, ok := symbol.(string)
			if !ok {
				fmt.Print("Failed to convert symbol to string\n")
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			tokenURI, err := contract.Call(c, "tokenURI", 2)
			if err != nil {
				fmt.Printf("Failed to get tokenURI: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			tokenURIStr, ok := tokenURI.(string)
			if !ok {
				fmt.Print("Failed to convert tokenURI to string\n")
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
		/*
			decimal, err := contract.Call(c, "decimals")
			if err != nil {
				fmt.Printf("Failed to get decimal: %v \n", err)
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			decimalUint8, ok := decimal.(uint8)
			if !ok {
				fmt.Print("Failed to convert decimal\n")
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": utils.ToDecimal(balanceBigInt, int(decimalUint8)).String(),
			})
		*/
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
