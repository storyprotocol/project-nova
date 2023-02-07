package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NftTokenRepository interface {
	GetNftByTokenId(tokenId int, collectionAddress string) (*NftTokenModel, error)
	GetNftsByOwner(franchiseId int64, walletAddress string) ([]*NftTokenModel, error)
	UpdateNftBackstory(tokenId int, collectionAddress string, backstory *string) (*NftTokenModel, error)
	CreateNft(nftToken *NftTokenModel) (*NftTokenModel, error)
}

type NftTokenModel struct {
	ID                string `gorm:"primaryKey;column:id"`
	CollectionAddress string
	TokenId           int
	FranchiseId       int64
	OwnerAddress      *string
	Name              *string
	Description       *string
	ImageUrl          *string
	Image             *string
	AnimationUrl      *string
	Traits            *string
	Backstory         *string
	OwnerUpdatedAt    *time.Time
	StoryUpdatedAt    *time.Time
}

func (NftTokenModel) TableName() string {
	return "nft_token"
}

func NewNftTokenDbImpl(db *gorm.DB) NftTokenRepository {
	return &nftTokenDbImpl{
		db: db,
	}
}

type nftTokenDbImpl struct {
	db *gorm.DB
}

func (n *nftTokenDbImpl) GetNftByTokenId(tokenId int, collectionAddress string) (*NftTokenModel, error) {
	results := &NftTokenModel{}
	r := n.db.Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).First(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (n *nftTokenDbImpl) GetNftsByOwner(franchiseId int64, walletAddress string) ([]*NftTokenModel, error) {
	results := []*NftTokenModel{}
	r := n.db.Where("franchise_id = ? and owner_address = ?", franchiseId, walletAddress).
		Order("collection_address").Order("token_id").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (n *nftTokenDbImpl) UpdateNftBackstory(tokenId int, collectionAddress string, backstory *string) (*NftTokenModel, error) {
	nftToken := &NftTokenModel{}
	r := n.db.Model(&nftToken).Clauses(clause.Returning{}).Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).Update("backstory", *backstory)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return nftToken, nil
}

func (n *nftTokenDbImpl) CreateNft(nftToken *NftTokenModel) (*NftTokenModel, error) {
	r := n.db.Create(nftToken)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return nftToken, nil
}
