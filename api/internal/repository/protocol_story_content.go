package repository

import (
	"errors"
	"fmt"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
)

type ProtocolStoryContentRepository interface {
	GetContentByID(contentId string) (*entity.ProtocolStoryContentModel, error)
	CreateContent(chapter *entity.ProtocolStoryContentModel) error
}

func NewProtocolStoryContentDbImpl(db *gorm.DB) ProtocolStoryContentRepository {
	return &protocolStoryContentDbImpl{
		db: db,
	}
}

type protocolStoryContentDbImpl struct {
	db *gorm.DB
}

func (s *protocolStoryContentDbImpl) GetContentByID(contentId string) (*entity.ProtocolStoryContentModel, error) {
	result := &entity.ProtocolStoryContentModel{}
	r := s.db.Where("id = ?", contentId).Find(result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}

func (s *protocolStoryContentDbImpl) CreateContent(content *entity.ProtocolStoryContentModel) error {
	r := s.db.Create(content)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return nil
}
