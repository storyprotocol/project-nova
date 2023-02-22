package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type NftCollectionRepository interface {
	GetCollections(collectionAddresses []string) ([]*NftCollectionModel, error)
	GetAllCollections() ([]*NftCollectionModel, error)
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
	for idx, address := range collectionAddresses {
		collectionAddresses[idx] = strings.ToLower(address)
	}

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
