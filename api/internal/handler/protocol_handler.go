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

// GET /franchise/:franchiseAddress
func NewGetFranchiseCollectionsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.SingleFranchise)
	}
}

// GET /character/:franchiseAddress/:collectionAddress
func NewGetCharactersHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Characters)
	}
}

// GET /character/:franchiseAddress/:collectionAddress/:tokenId
func NewGetCharacterHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.SingleCharacter)
	}
}

// GET /character/:franchiseAddress/:collectionAddress/:tokenId/collectors
func NewGetCollectorsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Collectors)
	}
}

// Get /story/:franchiseAddress/:collectionAddress
func NewGetStoriesHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.Stories)
	}
}

// Get /story/:franchiseAddress/:collectionAddress/:tokenId
func NewGetStoryHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, entity.SingleStory)
	}
}

// GET
func NewGetDerivativesHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
