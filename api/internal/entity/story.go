package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/project-nova/backend/pkg/gateway"
)

type StoryChapterResp struct {
	ChapterNum int    `json:"chapterNum"`
	Title      string `json:"title"`
	Heading    string `json:"heading"`
	CoverUrl   string `json:"coverUrl"`
}

// StoryChapterModel represents the story chapter's model in data storage
type StoryChapterModel struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	StoryId   string    `json:"storyId"`
	SeqNum    int       `json:"seqNum"`
	Title     *string   `json:"title"`
	Heading   *string   `json:"heading"`
	CoverUrl  *string   `json:"coverUrl"`
	ReleaseAt time.Time `json:"releaseAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (StoryChapterModel) TableName() string {
	return "story_chapter"
}

func (s *StoryChapterModel) ToStoryChapterResp() *StoryChapterResp {
	resp := &StoryChapterResp{
		ChapterNum: s.SeqNum,
	}

	if s.Heading != nil {
		resp.Heading = *s.Heading
	}
	if s.Title != nil {
		resp.Title = *s.Title
	}
	if s.CoverUrl != nil {
		resp.CoverUrl = *s.CoverUrl
	}

	return resp
}

func ToStoryChapterModel(request *gateway.CreateStoryChapterRequestBody, storyId string, chapterNum int) *StoryChapterModel {
	storyChapterModel := &StoryChapterModel{
		ID:      uuid.New().String(),
		StoryId: storyId,
		SeqNum:  chapterNum,
	}

	if request.Heading != nil {
		storyChapterModel.Heading = request.Heading
	}
	if request.Title != nil {
		storyChapterModel.Title = request.Title
	}
	if request.CoverUrl != nil {
		storyChapterModel.CoverUrl = request.CoverUrl
	}
	if request.ReleaseAt == nil {
		curr := time.Now().UTC()
		request.ReleaseAt = &curr
	}
	storyChapterModel.ReleaseAt = *request.ReleaseAt

	return storyChapterModel
}
