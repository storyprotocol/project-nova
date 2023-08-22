package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/service"
	xhttp "github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
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
		franchisesFinal := []*entity.Franchise{}
		for _, franchise := range franchises {
			var franchiseMetaData entity.FranchiseMetadata
			_, err = httpClient.Request("GET", franchise.TokenUri, nil, &franchiseMetaData)
			if err != nil {
				logger.Errorf("Failed to get story metadata from remote storage. url: %s, error: %v", franchise.TokenUri, err)
				continue
			}
			franchise.BannerUrl = franchiseMetaData.BannerUrl
			franchise.ImageUrl = franchiseMetaData.ImageUrl

			franchise.Metrics.StoryCount = 424
			franchise.Metrics.LicenseCount = 210
			franchise.Metrics.Revenue = "$60K"

			franchisesFinal = append(franchisesFinal, franchise)
		}

		c.JSON(http.StatusOK, franchisesFinal)
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
		charactersFinal := []*entity.CharacterInfoModel{}
		for _, character := range characters {
			var characterMetaData entity.CharacterMetadata
			_, err = httpClient.Request("GET", *character.MediaUri, nil, &characterMetaData)
			if err != nil {
				logger.Errorf("Failed to get character metadata from remote storage. url: %s, error: %v", *character.MediaUri, err)
				continue
			}
			character.ImageUrl = characterMetaData.ImageUrl
			charactersFinal = append(charactersFinal, character)
		}

		c.JSON(http.StatusOK, charactersFinal)
	}
}

// GET /character/:characterId
func NewGetCharacterHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.DefaultQuery("franchiseId", ""))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}
		_, err = strconv.ParseInt(c.Param("characterId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid character Id: %s", c.Param("characterId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid character Id"))
			return
		}

		// 1. call the graph service to get all the characters
		character, err := graphService.GetCharacter(franchiseId, c.Param("characterId"))
		if err != nil {
			logger.Errorf("Failed to get the character: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		// 2. Get the metadata from arweave
		var characterMetaData entity.CharacterMetadata
		_, err = httpClient.Request("GET", *character.MediaUri, nil, &characterMetaData)
		if err != nil {
			logger.Errorf("Failed to get character metadata from remote storage. url: %s, error: %v", *character.MediaUri, err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}
		character.ImageUrl = characterMetaData.ImageUrl

		c.JSON(http.StatusOK, character)
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
		storiesFinal := []*entity.StoryInfoV2Model{}
		for _, story := range stories {
			var storyMetaData entity.StoryMetadata
			_, err = httpClient.Request("GET", *story.MediaUri, nil, &storyMetaData)
			if err != nil {
				logger.Errorf("Failed to get story metadata from remote storage. url: %s, error: %v", *story.MediaUri, err)
				continue
			}
			story.Content = &storyMetaData.Content
			storiesFinal = append(storiesFinal, story)
		}

		c.JSON(http.StatusOK, storiesFinal)
	}
}

// GET /story/:storyId
func NewGetStoryHandlerKbw(graphService service.TheGraphService, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := strconv.ParseInt(c.Param("storyId"), 10, 64)
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
		story, err := graphService.GetStory(franchiseId, c.Param("storyId"))
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

// GET /license
func NewGetLicensesHandlerKbw(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := strconv.ParseInt(c.DefaultQuery("ipAssetId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Invalid ip asset id: %s", c.DefaultQuery("ipAssetId", ""))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid ip asset id"))
			return
		}
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Invalid franchise id: %s", c.DefaultQuery("franchiseId", ""))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid franchise id"))
			return
		}

		walletAddress := c.DefaultQuery("walletAddress", "")

		var licenses []*entity.License
		if walletAddress != "" {
			walletAddress, err = utils.SanitizeAddress(walletAddress)
			if err != nil {
				logger.Errorf("Invalid wallet address: %s", walletAddress)
				c.JSON(http.StatusBadRequest, ErrorMessage("Invalid wallet address"))
				return
			}
			licenses, err = graphService.GetLicensesByProfile(franchiseId, c.DefaultQuery("ipAssetId", ""), walletAddress)
			if err != nil {
				logger.Errorf("Failed to get licenses by profile from the graph service. error: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
		} else {
			licenses, err = graphService.GetLicensesByIpAsset(franchiseId, c.DefaultQuery("ipAssetId", ""))
			if err != nil {
				logger.Errorf("Failed to get licenses by IP Asset from the graph service. error: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
				return
			}
		}

		c.JSON(http.StatusOK, licenses)
	}
}

// GET /license/:licenseId
func NewGetLicenseHandlerKbw(graphService service.TheGraphService) func(c *gin.Context) {
	return func(c *gin.Context) {
		licenseId, err := strconv.ParseInt(c.Param("licenseId"), 10, 64)
		if err != nil {
			logger.Errorf("Invalid license id: %s", c.Param("licenseId"))
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid license id"))
			return
		}

		license, err := graphService.GetLicense(licenseId)
		if err != nil {
			logger.Errorf("Failed to get license by id from the graph service. error: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, license)
	}
}
