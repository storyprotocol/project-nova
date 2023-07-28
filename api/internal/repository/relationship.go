package repository

import (
	"errors"
	"fmt"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
)

type RelationshipRepository interface {
	CreateRelationship(relationship *entity.RelationshipModel) error
	GetRelationship(srcId string, dstId string) (*entity.RelationshipModel, error)
	GetRelationships() ([]*entity.RelationshipModel, error)
	GetRelationshipsByNodeId(id string) ([]*entity.RelationshipModel, error)
}

func NewRelationshipDbImpl(db *gorm.DB) RelationshipRepository {
	return &relationshipDbImpl{
		db: db,
	}
}

type relationshipDbImpl struct {
	db *gorm.DB
}

func (r *relationshipDbImpl) CreateRelationship(relationship *entity.RelationshipModel) error {
	result := r.db.Debug().Create(relationship)

	if result.Error != nil {
		return fmt.Errorf("failed to insert into db: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("failed to create relationship: %v", gorm.ErrRecordNotFound)
	}
	return nil
}

func (r *relationshipDbImpl) GetRelationship(srcId string, dstId string) (*entity.RelationshipModel, error) {
	result := &entity.RelationshipModel{}
	query := r.db.Where("src_id = ? and dst_id = ?", srcId, dstId).First(result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil, query.Error
	}
	if query.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", query.Error)
	}

	return result, nil
}

func (r *relationshipDbImpl) GetRelationships() ([]*entity.RelationshipModel, error) {
	results := []*entity.RelationshipModel{}
	query := r.db.Find(&results)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil, query.Error
	}
	if query.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", query.Error)
	}

	return results, nil
}

func (r *relationshipDbImpl) GetRelationshipsByNodeId(id string) ([]*entity.RelationshipModel, error) {
	results := []*entity.RelationshipModel{}
	query := r.db.Where("src_id = ? OR dst_id = ?", id, id).Find(&results)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil, query.Error
	}
	if query.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", query.Error)
	}

	return results, nil
}
