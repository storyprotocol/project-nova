package repository

import (
	"errors"
	"fmt"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
)

type StoryInfoV2Repository interface {
	GetStory(franchiseId int64, storyId int64) (*entity.StoryInfoV2Model, error)
	GetStories(franchiseId int64) ([]*entity.StoryInfoV2Model, error)
	CreateStory(story *entity.StoryInfoV2Model) error
}

func NewStoryInfoV2DbImpl(db *gorm.DB) StoryInfoV2Repository {
	return &storyInfoV2DbImpl{
		db: db,
	}
}

type storyInfoV2DbImpl struct {
	db *gorm.DB
}

func (s *storyInfoV2DbImpl) GetStory(franchiseId int64, storyId int64) (*entity.StoryInfoV2Model, error) {
	result := &entity.StoryInfoV2Model{}
	r := s.db.Where("franchise_id = ? and story_id = ?", franchiseId, storyId).First(&result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}

func (c *storyInfoV2DbImpl) GetStories(franchiseId int64) ([]*entity.StoryInfoV2Model, error) {
	results := []*entity.StoryInfoV2Model{}
	r := c.db.Where("franchise_id = ?", franchiseId).Order("story_id asc").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (c *storyInfoV2DbImpl) CreateStory(story *entity.StoryInfoV2Model) error {
	r := c.db.Create(story)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return nil
}
