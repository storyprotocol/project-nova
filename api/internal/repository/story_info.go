package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StoryInfoRepository interface {
	GetStoryByFranchise(franchiseId int64, storyNum int) (*StoryInfoModel, error)
}

type StoryInfoModel struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	FranchiseId int64     `json:"franchiseId"`
	SeqNum      int       `json:"seqNum"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	CoverUrl    string    `json:"coverUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (StoryInfoModel) TableName() string {
	return "story_info"
}

func NewStoryInfoDbImpl(db *gorm.DB) StoryInfoRepository {
	return &storyInfoDbImpl{
		db: db,
	}
}

type storyInfoDbImpl struct {
	db *gorm.DB
}

func (s *storyInfoDbImpl) GetStoryByFranchise(franchiseId int64, storyNum int) (*StoryInfoModel, error) {
	result := &StoryInfoModel{}
	r := s.db.Where("franchise_id = ? and seq_num = ?", franchiseId, storyNum).First(&result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}
