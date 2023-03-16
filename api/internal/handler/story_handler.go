package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/logger"
)

// NewGetStoryChaptersHandler creates a handler to handle GET /story/:franchiseId/:storyNum request.
// Doc: (To Be Added)
func NewGetStoryChaptersHandler(
	storyChapterRepo repository.StoryChapterRepository,
	storyInfoRepo repository.StoryInfoRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert franchise id: %v", err)
			c.String(http.StatusBadRequest, "franchise id is invalid")
			return
		}

		storyNum, err := strconv.Atoi(c.Param("storyNum"))
		if err != nil {
			logger.Errorf("Failed to convert story num: %v", err)
			c.String(http.StatusBadRequest, "story num is invalid")
			return
		}

		storyInfoResult, err := storyInfoRepo.GetStoryByFranchise(franchiseId, storyNum)
		if err != nil {
			logger.Errorf("Failed to get story info: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		storyId := storyInfoResult.ID

		storyChapters, err := storyChapterRepo.GetChaptersByID(storyId)
		if err != nil {
			logger.Errorf("Failed to get story chapters: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		var resp []*entity.StoryChapterResp
		for _, v := range storyChapters {
			if v.ReleaseAt.Before(time.Now()) {
				resp = append(resp, v.ToStoryChapterResp())
			}

		}

		c.JSON(http.StatusOK, resp)
	}
}

// NewGetStoryChapterContentsHandler creates the handler to handle /story/:franchiseId/:storyNum/:chapterNum request.
// Doc: (To Be Added)
func NewGetStoryChapterContentsHandler(
	storyContentRepo repository.StoryContentRepository,
	storyChapterRepo repository.StoryChapterRepository,
	storyInfoRepo repository.StoryInfoRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.Param("franchiseId"), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert franchise id: %v", err)
			c.String(http.StatusBadRequest, "franchise id is invalid")
			return
		}

		storyNum, err := strconv.Atoi(c.Param("storyNum"))
		if err != nil {
			logger.Errorf("Failed to convert story num: %v", err)
			c.String(http.StatusBadRequest, "story num is invalid")
			return
		}

		chapterNum, err := strconv.Atoi(c.Param("chapterNum"))
		if err != nil {
			logger.Errorf("Failed to convert chapter num: %v", err)
			c.String(http.StatusBadRequest, "chapter num is invalid")
			return
		}

		storyInfoResult, err := storyInfoRepo.GetStoryByFranchise(franchiseId, storyNum)
		if err != nil {
			logger.Errorf("Failed to get story info: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		storyId := storyInfoResult.ID

		storyChapter, err := storyChapterRepo.GetChapter(storyId, chapterNum)
		if err != nil {
			logger.Errorf("Failed to get story chapter: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		if storyChapter.ReleaseAt.After(time.Now()) {
			logger.Error("Denied chapter request, chapter content is not released yet")
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		storyContents, err := storyContentRepo.GetContentByChapter(franchiseId, storyNum, chapterNum)
		if err != nil {
			logger.Errorf("Failed to get story contents: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, storyContents)
	}
}
