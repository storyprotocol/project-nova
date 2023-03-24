package repository

import (
	"errors"
	"fmt"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
)

type StoryChapterRepository interface {
	GetChaptersByID(storyId string) ([]*entity.StoryChapterModel, error)
	GetChapter(storyId string, chapterNum int) (*entity.StoryChapterModel, error)
	CreateChapter(chapter *entity.StoryChapterModel) error
}

func NewStoryChapterDbImpl(db *gorm.DB) StoryChapterRepository {
	return &storyChapterDbImpl{
		db: db,
	}
}

type storyChapterDbImpl struct {
	db *gorm.DB
}

func (s *storyChapterDbImpl) GetChaptersByID(storyId string) ([]*entity.StoryChapterModel, error) {
	results := []*entity.StoryChapterModel{}
	r := s.db.Where("story_id = ?", storyId).Order("seq_num asc").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (s *storyChapterDbImpl) GetChapter(storyId string, chapterNum int) (*entity.StoryChapterModel, error) {
	result := &entity.StoryChapterModel{}
	r := s.db.Where("story_id = ? and seq_num = ?", storyId, chapterNum).First(result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}

func (s *storyChapterDbImpl) CreateChapter(chapter *entity.StoryChapterModel) error {
	r := s.db.Create(chapter)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return nil
}
