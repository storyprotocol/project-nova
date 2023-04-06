package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
)

// GET /franchise
func NewGetFranchisesHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Franchises)
	}
}

// GET /character/:franchiseAddress/:collectionAddress
func NewGetCharactersHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Characters)
	}
}

// GET
func NewGetCollectorsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Collectors)
	}
}

// GET
func NewGetStoriesHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Stories)
	}
}

// GET
func NewGetDerivativesHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
