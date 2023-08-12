package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/service"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
)

// GET /franchise
func NewGetFranchisesHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. call the graph service to get all the franchises
		franchises, err := graphService.GetFranchises()
		if err != nil {
			logger.Errorf("Failed to get franchises: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 2. Get the metadata from arweave
		for _, franchise := range franchises {
			var franchiseMetaData entity.FranchiseMetadata
			_, err = httpClient.Request("GET", franchise.TokenUri, nil, &franchiseMetaData)
			if err != nil {
				logger.Errorf("Failed to get story metadata from remote storage. url: %s, error: %v", franchise.TokenUri, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
			franchise.BannerUrl = franchiseMetaData.BannerUrl
			franchise.ImageUrl = franchiseMetaData.ImageUrl

			franchise.Metrics.StoryCount = 424
			franchise.Metrics.LicenseCount = 210
			franchise.Metrics.Revenue = "$60K"
		}

		c.JSON(http.StatusOK, franchises)
	}
}

// GET /franchise/:franchiseId
func NewGetFranchiseHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.Param("franchiseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		// 1. call the graph service to get all the franchises
		franchise, err := graphService.GetFranchise(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get franchises: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 2. Get the metadata from arweave
		var franchiseMetaData entity.FranchiseMetadata
		_, err = httpClient.Request("GET", franchise.TokenUri, nil, &franchiseMetaData)
		if err != nil {
			logger.Errorf("Failed to get franchise metadata from remote storage. url: %s, error: %v", franchise.TokenUri, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		franchise.BannerUrl = franchiseMetaData.BannerUrl
		franchise.ImageUrl = franchiseMetaData.ImageUrl

		franchise.Metrics.StoryCount = 424
		franchise.Metrics.LicenseCount = 210
		franchise.Metrics.Revenue = "$60K"

		c.JSON(http.StatusOK, franchise)
	}
}

// GET /character
func NewGetCharactersHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.DefaultQuery("franchiseId", ""))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		// 1. call the graph service to get all the characters
		characters, err := graphService.GetCharacters(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get characters: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 2. Get the metadata from arweave
		for _, character := range characters {
			var characterMetaData entity.CharacterMetadata
			_, err = httpClient.Request("GET", *character.MediaUri, nil, &characterMetaData)
			if err != nil {
				logger.Errorf("Failed to get character metadata from remote storage. url: %s, error: %v", *character.MediaUri, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
			character.ImageUrl = characterMetaData.ImageUrl
		}

		c.JSON(http.StatusOK, characters)
	}
}

// GET /story
func NewGetStoriesHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.DefaultQuery("franchiseId", ""))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		// 1. call the graph service to get all the stories
		stories, err := graphService.GetStories(franchiseId)
		if err != nil {
			logger.Errorf("Failed to get stories: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 2. Get the metadata from arweave
		for _, story := range stories {
			var storyMetaData entity.StoryMetadata
			_, err = httpClient.Request("GET", *story.MediaUri, nil, &storyMetaData)
			if err != nil {
				logger.Errorf("Failed to get story metadata from remote storage. url: %s, error: %v", *story.MediaUri, err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
			story.Content = &storyMetaData.Content
		}

		c.JSON(http.StatusOK, stories)
	}
}

// GET /story/:storyId
func NewGetStoryHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		storyId, err := strconv.ParseInt(c.Param("storyId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid story id: %s", c.Param("storyId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid story id"))
			return
		}
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.DefaultQuery("franchiseId", ""))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		// 1. call the graph service to get all the stories
		story, err := graphService.GetStory(franchiseId, storyId)
		if err != nil {
			logger.Errorf("Failed to get story: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 2. Get the metadata from arweave
		var storyMetaData entity.StoryMetadata
		_, err = httpClient.Request("GET", *story.MediaUri, nil, &storyMetaData)
		if err != nil {
			logger.Errorf("Failed to get story metadata from remote storage. url: %s, error: %v", *story.MediaUri, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		story.Content = &storyMetaData.Content

		c.JSON(http.StatusOK, story)
	}
}
