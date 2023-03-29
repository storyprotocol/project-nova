package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type NftAllowlistRepository interface {
	UpdateCollectionAddress(oldAddress string, newAddress string) error
}

type NftAllowlistModel struct {
	ID                string    `gorm:"primaryKey;column:id" json:"id"`
	Type              string    `json:"type"`
	CollectionAddress string    `json:"collectionAddress"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (NftAllowlistModel) TableName() string {
	return "nft_allowlist"
}

func NewNftAllowlistDbImpl(db *gorm.DB) NftAllowlistRepository {
	return &nftAllowlistDbImpl{
		db: db,
	}
}

type nftAllowlistDbImpl struct {
	db *gorm.DB
}

func (n *nftAllowlistDbImpl) UpdateCollectionAddress(oldAddress string, newAddress string) error {
	r := n.db.Model(&NftAllowlistModel{}).Where("collection_address = ?", oldAddress).
		Update("collection_address", newAddress)
	if r.RowsAffected == 0 {
		return fmt.Errorf("no rows are affected")
	}
	if r.Error != nil {
		return fmt.Errorf("failed to update db: %v", r.Error)
	}
	return nil
}
