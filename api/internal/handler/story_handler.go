package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/api/internal/repository"
	"github.com/project-nova/backend/pkg/logger"
)

type GetStoryChaptersResp struct {
	ChapterNum int
	Title      string
	Subtitle   string
	CoverUrl   string
}

func NewGetStoryChaptersHandler(
	storyChapterRepo repository.StoryChapterRepository,
	storyInfoRepo repository.StoryInfoRepository,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert franchise id: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("franchise id is invalid, id: %v", franchiseId))
			return
		}

		storyNum, err := strconv.Atoi(c.Param("storyNum"))
		if err != nil {
			logger.Errorf("Failed to convert story num: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("story num is invalid, number: %v", storyNum))
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

func NewGetStoryChapterContentsHandler(storyContentRepo repository.StoryContentRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId, err := strconv.ParseInt(c.DefaultQuery("franchiseId", ""), 10, 64)
		if err != nil {
			logger.Errorf("Failed to convert franchise id: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("franchise id is invalid, id: %v", franchiseId))
			return
		}

		storyNum, err := strconv.Atoi(c.Param("storyNum"))
		if err != nil {
			logger.Errorf("Failed to convert story num: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("story num is invalid, number: %v", storyNum))
			return
		}

		chapterNum, err := strconv.Atoi(c.Param("chapterNum"))
		if err != nil {
			logger.Errorf("Failed to convert chapter num: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("chapter num is invalid, number: %v", chapterNum))
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
