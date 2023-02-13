package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NftTokenRepository interface {
	GetNftByTokenId(tokenId int, collectionAddress string) (*NftTokenModel, error)
	GetNfts(collectionAddresses []string, walletAddress string) ([]*NftTokenModel, error)
	GetNftsByOwner(franchiseId int64, walletAddress string) ([]*NftTokenModel, error)
	UpdateNftBackstory(tokenId int, collectionAddress string, backstory *string) (*NftTokenModel, error)
	UpdateNftOwner(tokenId int, collectionAddress string, ownerAddress string) (*NftTokenModel, error)
	CreateNft(nftToken *NftTokenModel) (*NftTokenModel, error)
	DeleteNft(tokenId int, collectionAddress string) error
}

type NftTokenModel struct {
	ID                string `gorm:"primaryKey;column:id"`
	CollectionAddress string
	TokenId           int
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
	collectionAddress = strings.ToLower(collectionAddress)
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

func (n *nftTokenDbImpl) GetNfts(collectionAddresses []string, walletAddress string) ([]*NftTokenModel, error) {
	for idx, address := range collectionAddresses {
		collectionAddresses[idx] = strings.ToLower(address)
	}
	query := n.db.Where("collection_address IN ?", collectionAddresses)

	if walletAddress != "" {
		walletAddress = strings.ToLower(walletAddress)
		query = query.Where("owner_address = ?", walletAddress)
	}

	results := []*NftTokenModel{}
	r := query.Order("collection_address").Order("token_id").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}

func (n *nftTokenDbImpl) GetNftsByOwner(franchiseId int64, walletAddress string) ([]*NftTokenModel, error) {
	walletAddress = strings.ToLower(walletAddress)
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
	collectionAddress = strings.ToLower(collectionAddress)

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

func (n *nftTokenDbImpl) UpdateNftOwner(tokenId int, collectionAddress string, ownerAddress string) (*NftTokenModel, error) {
	collectionAddress = strings.ToLower(collectionAddress)
	ownerAddress = strings.ToLower(ownerAddress)

	nftToken := &NftTokenModel{}
	r := n.db.Model(&nftToken).Clauses(clause.Returning{}).Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).Update("owner_address", ownerAddress)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return nftToken, nil
}

func (n *nftTokenDbImpl) CreateNft(nftToken *NftTokenModel) (*NftTokenModel, error) {
	nftToken.CollectionAddress = strings.ToLower(nftToken.CollectionAddress)
	if nftToken.OwnerAddress != nil {
		ownerAddress := strings.ToLower(*nftToken.OwnerAddress)
		nftToken.OwnerAddress = &ownerAddress
	}

	r := n.db.Create(nftToken)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return nftToken, nil
}

func (n *nftTokenDbImpl) DeleteNft(tokenId int, collectionAddress string) error {
	collectionAddress = strings.ToLower(collectionAddress)

	r := n.db.Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).Delete(&NftTokenModel{})
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to delete the row from db: %v", r.Error)
	}

	return nil
}
