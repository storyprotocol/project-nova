package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StoryChapterRepository interface {
	GetChaptersByID(id string) ([]*StoryChapterModel, error)
}

type StoryChapterModel struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	StoryId   string    `json:"storyId"`
	SeqNum    int       `json:"seqNum"`
	Title     string    `json:"title"`
	CoverUrl  string    `json:"coverUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (StoryChapterModel) TableName() string {
	return "story_chapter"
}

func NewStoryChapterDbImpl(db *gorm.DB) StoryChapterRepository {
	return &storyChapterDbImpl{
		db: db,
	}
}

type storyChapterDbImpl struct {
	db *gorm.DB
}

func (s *storyChapterDbImpl) GetChaptersByID(id string) ([]*StoryChapterModel, error) {
	results := []*StoryChapterModel{}
	r := s.db.Where("story_id = ?", id).Order("seq_num asc").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}
