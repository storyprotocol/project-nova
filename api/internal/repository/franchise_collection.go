package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type FranchiseCollectionRepository interface {
	GetCollectionAddressesByFranchise(franchiseId int64) ([]string, error)
	GetAndLoadFranchiseCollections() error
}

type FranchiseCollectionModel struct {
	ID                string    `gorm:"primaryKey;column:id" json:"id"`
	FranchiseId       int64     `json:"franchiseId"`
	CollectionAddress string    `json:"collectionAddress"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (FranchiseCollectionModel) TableName() string {
	return "franchise_collection"
}

func NewFranchiseCollectionDbImpl(db *gorm.DB) FranchiseCollectionRepository {
	return &franchiseCollectionDbImpl{
		db:                      db,
		franchiseCollectionsMap: make(map[int64][]string),
	}
}

type franchiseCollectionDbImpl struct {
	db                      *gorm.DB
	franchiseCollectionsMap map[int64][]string
}

func (s *franchiseCollectionDbImpl) GetCollectionAddressesByFranchise(franchiseId int64) ([]string, error) {
	// Check cache
	if val, ok := s.franchiseCollectionsMap[franchiseId]; ok {
		return val, nil
	}

	result := []string{}
	r := s.db.Where("franchise_id = ?", franchiseId).Find(&result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	// Update cache
	s.franchiseCollectionsMap[franchiseId] = result

	return result, nil
}

func (s *franchiseCollectionDbImpl) GetAndLoadFranchiseCollections() error {
	result := []FranchiseCollectionModel{}

	r := s.db.Find(&result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return r.Error
	}
	if r.Error != nil {
		return fmt.Errorf("failed to scan db: %v", r.Error)
	}

	// Load the records to local cache
	for _, record := range result {
		s.franchiseCollectionsMap[record.FranchiseId] = append(s.franchiseCollectionsMap[record.FranchiseId], record.CollectionAddress)
	}

	return nil
}
