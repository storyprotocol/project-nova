package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-nova/backend/pkg/logger"
	"gorm.io/gorm"
)

type StoryInfoModel struct {
	ID          string `gorm:"primaryKey;column:id"`
	FranchiseId int64
	SeqNum      int
	Title       string
	Subtitle    string
	CoverUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (StoryInfoModel) TableName() string {
	return "story_info"
}

type StoryChapterModel struct {
	ID        string `gorm:"primaryKey;column:id"`
	StoryId   string
	SeqNum    int
	Title     string
	CoverUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (StoryChapterModel) TableName() string {
	return "story_chapter"
}

type GetStoryChaptersResp struct {
	ChapterNum int
	Title      string
	Subtitle   string
	CoverUrl   string
}

func NewGetStoryChaptersHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		franchiseId := c.DefaultQuery("franchiseId", "")
		storyNum := c.Param("storyNum")

		// Validate address
		if franchiseId == "" {
			c.String(http.StatusBadRequest, fmt.Sprint("franchise id is not specified"))
			return
		}

		storyInfoResult := &StoryInfoModel{}
		r := db.Where("franchise_id = ? and seq_num = ?", franchiseId, storyNum).First(&storyInfoResult)
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		if r.Error != nil {
			logger.Errorf("Failed to query db: %v", r.Error)
			c.String(http.StatusInternalServerError, "Internal server error")
			return
		}

		storyId := storyInfoResult.ID
		storyChapterResults := []*StoryChapterModel{}
		r = db.Where("story_id = ?", storyId).Order("seq_num asc").Find(&storyChapterResults)
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		if r.Error != nil {
			logger.Errorf("Failed to query db: %v", r.Error)
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

func NewGetStoryChapterContentsHandler(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
