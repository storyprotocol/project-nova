package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type NftCollectionRepository interface {
	GetCollections(collectionAddresses []string) ([]*NftCollectionModel, error)
	GetAllCollections() ([]*NftCollectionModel, error)
	UpdateCollectionAddress(oldAddress string, newAddress string) error
}

type NftCollectionModel struct {
	ID                string    `gorm:"primaryKey;column:id" json:"id"`
	CollectionAddress string    `json:"collectionAddress"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	TotalCap          int       `json:"totalCap"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (NftCollectionModel) TableName() string {
	return "nft_collection"
}

func NewNftCollectionDbImpl(db *gorm.DB) NftCollectionRepository {
	return &nftCollectionDbImpl{
		db: db,
	}
}

type nftCollectionDbImpl struct {
	db *gorm.DB
}

func (s *nftCollectionDbImpl) GetCollections(collectionAddresses []string) ([]*NftCollectionModel, error) {
	results := []*NftCollectionModel{}
	r := s.db.Where("collection_address IN ?", collectionAddresses).Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (s *nftCollectionDbImpl) GetAllCollections() ([]*NftCollectionModel, error) {
	results := []*NftCollectionModel{}
	r := s.db.Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (s *nftCollectionDbImpl) UpdateCollectionAddress(oldAddress string, newAddress string) error {
	r := s.db.Model(&NftCollectionModel{}).Where("collection_address = ?", oldAddress).
		Update("collection_address", newAddress)
	if r.RowsAffected == 0 {
		return fmt.Errorf("no rows are affected")
	}
	if r.Error != nil {
		return fmt.Errorf("failed to update db: %v", r.Error)
	}
	return nil
}
