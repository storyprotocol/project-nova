package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NftTokenRepository interface {
	GetNftByTokenId(tokenId int, collectionAddress string) (*NftTokenModel, error)
	GetNfts(collectionAddresses []string, walletAddress string, offset *int, limit *int) ([]*entity.NftTokenResponse, error)
	GetFilteredNfts(collectionAddresses []string, walletAddress string, filter *entity.Filter, offset *int, limit *int) ([]*entity.NftTokenResponse, int, error)
	GetNftsByOwner(franchiseId int64, walletAddress string) ([]*NftTokenModel, error)
	UpdateNftBackstory(tokenId int, collectionAddress string, backstory *string) (*NftTokenModel, error)
	UpdateNftOwner(tokenId int, collectionAddress string, ownerAddress string) (*NftTokenModel, error)
	CreateNft(nftToken *NftTokenModel) (*NftTokenModel, error)
	UpdateNft(nftToken *NftTokenModel) (*NftTokenModel, error)
	DeleteNft(tokenId int, collectionAddress string) error
}

type NftTokenModel struct {
	ID                string     `gorm:"primaryKey;column:id" json:"id"`
	CollectionAddress string     `json:"collectionAddress"`
	TokenId           int        `json:"tokenId"`
	OwnerAddress      *string    `json:"ownerAddress"`
	Name              *string    `json:"name"`
	Description       *string    `json:"description"`
	ImageUrl          *string    `json:"imageUrl"`
	Image             *string    `json:"image"`
	AnimationUrl      *string    `json:"animationUrl"`
	Traits            *string    `json:"traits"`
	Backstory         *string    `json:"backstory"`
	OwnerUpdatedAt    *time.Time `json:"ownerUpdatedAt"`
	StoryUpdatedAt    *time.Time `json:"storyUpdatedAt"`
}

func (NftTokenModel) TableName() string {
	return "nft_token"
}

func NewNftTokenDbImpl(db *gorm.DB) NftTokenRepository {
	return &nftTokenDbImpl{
		db: db,
	}
}

func (n *NftTokenModel) ToNftTokenResponse() (*entity.NftTokenResponse, error) {
	if n == nil {
		return nil, fmt.Errorf("input nft token model is nil")
	}

	nftResponse := &entity.NftTokenResponse{
		CollectionAddress: n.CollectionAddress,
		TokenId:           n.TokenId,
		OwnerAddress:      n.OwnerAddress,
		Name:              n.Name,
		Description:       n.Description,
		ImageUrl:          n.ImageUrl,
		Image:             n.Image,
		AnimationUrl:      n.AnimationUrl,
		Backstory:         n.Backstory,
	}

	if n.Traits != nil {
		var traits []*entity.NftTrait
		err := json.Unmarshal([]byte(*n.Traits), &traits)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal traits: %v", err)
		}
		for _, trait := range traits {
			nftTraitResponse := entity.NftTraitResponse(*trait)
			nftResponse.Traits = append(nftResponse.Traits, &nftTraitResponse)
		}
	}

	return nftResponse, nil
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

func (n *nftTokenDbImpl) GetNfts(collectionAddresses []string, walletAddress string, offset *int, limit *int) ([]*entity.NftTokenResponse, error) {
	for idx, address := range collectionAddresses {
		collectionAddresses[idx] = strings.ToLower(address)
	}
	query := n.db.Where("collection_address IN ?", collectionAddresses)

	if walletAddress != "" {
		walletAddress = strings.ToLower(walletAddress)
		query = query.Where("owner_address = ?", walletAddress)
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	results := []*NftTokenModel{}
	r := query.Order("collection_address").Order("token_id").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) || len(results) == 0 {
		return []*entity.NftTokenResponse{}, nil
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	var response []*entity.NftTokenResponse
	for _, nftModel := range results {
		nftResponse, err := nftModel.ToNftTokenResponse()
		if err != nil {
			return nil, fmt.Errorf("failed to convert nft model to nft response: %v", err)
		}
		response = append(response, nftResponse)
	}

	return response, nil
}

func (n *nftTokenDbImpl) GetFilteredNfts(collectionAddresses []string, walletAddress string, filter *entity.Filter, offset *int, limit *int) ([]*entity.NftTokenResponse, int, error) {
	nftTokens, err := n.GetNfts(collectionAddresses, walletAddress, nil, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all nfts: %v", err)
	}

	total := len(nftTokens)
	if total == 0 {
		return []*entity.NftTokenResponse{}, 0, nil
	}

	var filteredResults []*entity.NftTokenResponse
	for _, nft := range nftTokens {
		if nft.Traits == nil {
			continue
		}
		matched := filter.Eval(nft.GetKeyValuesFromTraits())
		if matched {
			filteredResults = append(filteredResults, nft)
		}
	}

	totalFiltered := len(filteredResults)
	if totalFiltered == 0 {
		return []*entity.NftTokenResponse{}, 0, nil
	}

	var start int = 0
	var end int = totalFiltered
	if offset != nil {
		start = *offset
	}
	if limit != nil {
		end = start + *limit
	}

	return filteredResults[start:end], totalFiltered, nil
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

func (n *nftTokenDbImpl) UpdateNft(nftToken *NftTokenModel) (*NftTokenModel, error) {
	nftToken.CollectionAddress = strings.ToLower(nftToken.CollectionAddress)
	if nftToken.OwnerAddress != nil {
		ownerAddress := strings.ToLower(*nftToken.OwnerAddress)
		nftToken.OwnerAddress = &ownerAddress
	}

	r := n.db.Model(&NftTokenModel{}).Where("token_id = ?", nftToken.TokenId).Updates(nftToken)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to update into db: %v", r.Error)
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
