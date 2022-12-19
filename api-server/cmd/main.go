package main

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/database"
	"github.com/project-nova/backend/pkg/utils"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

type WhitelistWallet struct {
	ID          string `gorm:"primaryKey;column:id"`
	Address     *string
	MerkleProof *string
	CreatedAt   time.Time
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

	r.GET("/membership", func(c *gin.Context) {
		privateKey := "a6cd3f393b1cddf8be66e2ff784640adbafbce852267ec1ec0e22eb741232db6"

		sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
			PrivateKey: privateKey,
		})
		if err != nil {
			fmt.Printf("Failed to set up sdk: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		// You can replace your own contract address here
		contractAddress := "0x7af963cF6D228E564e2A0aA0DdBF06210B38615D"

		// Add your contract ABI here
		abi := `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"mint","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"account","type":"address"}],"name":"addMinter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"renounceMinter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"isMinter","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"name","type":"string"},{"name":"symbol","type":"string"},{"name":"decimals","type":"uint8"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"account","type":"address"}],"name":"MinterAdded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"account","type":"address"}],"name":"MinterRemoved","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`

		// Now you have a contract instance ready to go
		contract, err := sdk.GetContractFromAbi(contractAddress, abi)
		if err != nil {
			fmt.Printf("Failed to get contract: %v \n", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		balance, err := contract.Call(c, "balanceOf", "0xc81f0c814ed6184ff6ef2e9584a75d48197a9caa")
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
		//0x000000000000000000000000af14a532b76d6812e8d036a08be92c6dd6839a48,0xaa776f660f43011909b215fd43887ee32baa2e60fa5d4349a97f03ab8e9d4206
		c.JSON(http.StatusOK, gin.H{
			"message": utils.ToDecimal(balanceBigInt, int(decimalUint8)).String(),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
