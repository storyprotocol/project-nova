package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/logger"
)

func NewUpdateNftBackstoryHandler(nftTokenRepository repository.NftTokenRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		type UpdateNftBackstoryRequestBody struct {
			CollectionAddress string `json:"collectionAddress"`
			WalletAddress     string `json:"walletAddress"`
			Backstory         string `json:"backstory"`
		}

		var requestBody UpdateNftBackstoryRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			c.String(http.StatusBadRequest, "invalid request body")
			return
		}

		tokenId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.Errorf("Failed to convert token id: %v", err)
			c.String(http.StatusBadRequest, "token id is invalid")
			return
		}

		nftToken, err := nftTokenRepository.GetNftByTokenId(tokenId, requestBody.CollectionAddress)
		if err != nil {
			logger.Errorf("Failed to get nft by token id: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		if nftToken.OwnerAddress != requestBody.WalletAddress {
			logger.Errorf("The request wallet is not the owner of the nft, owner address: %s, wallet address: %s", nftToken.OwnerAddress, requestBody.WalletAddress)
			c.String(http.StatusForbidden, "The wallet doesn't have permission for this operation")
			return
		}

		nftToken, err = nftTokenRepository.UpdateNftBackstory(tokenId, requestBody.CollectionAddress, &requestBody.Backstory)
		if err != nil {
			logger.Errorf("Failed to update nft backstory: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, nftToken)
	}
}
