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
