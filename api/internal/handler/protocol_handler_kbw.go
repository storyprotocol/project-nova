package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/service"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/logger"
)

// GET /franchise
func NewGetFranchisesHandlerKbw(graphClient service.TheGraphService, web3Gateway gateway.Web3GatewayClient) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// GET /franchise/:franchiseId
func NewGetFranchiseHandlerKbw(graphClient service.TheGraphService, web3Gateway gateway.Web3GatewayClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}
	}
}

// GET /character
func NewGetCharactersHandlerKbw(graphClient service.TheGraphService, web3Gateway gateway.Web3GatewayClient) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// GET /story
func NewGetStoriesHandlerKbw(graphClient service.TheGraphService, web3Gateway gateway.Web3GatewayClient) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

// GET /story/:storyId
func NewGetStoryHandlerKbw(graphClient service.TheGraphService, web3Gateway gateway.Web3GatewayClient) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
