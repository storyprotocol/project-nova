package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/logger"
)

type GetStoryChaptersResp struct {
	ChapterNum int    `json:"chapterNum"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	CoverUrl   string `json:"coverUrl"`
}

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

		storyChapterResults, err := storyChapterRepo.GetChaptersByID(storyId)
		if err != nil {
			logger.Errorf("Failed to get story chapters: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		var resp []*GetStoryChaptersResp
		for _, v := range storyChapterResults {
			resp = append(resp, &GetStoryChaptersResp{
				ChapterNum: v.SeqNum,
				Title:      v.Title,
				CoverUrl:   v.CoverUrl,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

// NewGetStoryChapterContentsHandler creates the handler to handle /story/:franchiseId/:storyNum/:chapterNum request.
// Doc: (To Be Added)
func NewGetStoryChapterContentsHandler(storyContentRepo repository.StoryContentRepository) func(c *gin.Context) {
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

		storyContents, err := storyContentRepo.GetContentByChapter(franchiseId, storyNum, chapterNum)
		if err != nil {
			logger.Errorf("Failed to get story contents: %v", err)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.JSON(http.StatusOK, storyContents)
	}
}
