package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NftTokenRepository interface {
	GetNftByTokenId(tokenId int, collectionAddress string) (*entity.NftTokenModel, error)
	GetNfts(collectionAddresses []string, walletAddress string, offset *int, limit *int) ([]*entity.NftTokenResponse, error)
	GetFilteredNfts(collectionAddresses []string, walletAddress string, filter *entity.Filter, offset *int, limit *int) ([]*entity.NftTokenResponse, int, error)
	GetNftsByOwner(franchiseId int64, walletAddress string) ([]*entity.NftTokenModel, error)
	UpdateNftBackstory(tokenId int, collectionAddress string, backstory *string) (*entity.NftTokenModel, error)
	UpdateNftOwner(tokenId int, collectionAddress string, ownerAddress string) (*entity.NftTokenModel, error)
	CreateNft(nftToken *entity.NftTokenModel) (*entity.NftTokenModel, error)
	UpdateNft(nftToken *entity.NftTokenModel) (*entity.NftTokenModel, error)
	DeleteNft(tokenId int, collectionAddress string) error
}

func NewNftTokenDbImpl(db *gorm.DB) NftTokenRepository {
	return &nftTokenDbImpl{
		db: db,
	}
}

type nftTokenDbImpl struct {
	db *gorm.DB
}

func (n *nftTokenDbImpl) GetNftByTokenId(tokenId int, collectionAddress string) (*entity.NftTokenModel, error) {
	collectionAddress = strings.ToLower(collectionAddress)
	results := &entity.NftTokenModel{}
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

	results := []*entity.NftTokenModel{}
	r := query.Order("collection_address").Order("token_id").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) || len(results) == 0 {
		return []*entity.NftTokenResponse{}, nil
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	response := []*entity.NftTokenResponse{}
	for _, nftModel := range results {
		if nftModel.Image == nil { // When image is not present, it means the nft is not revealed yet.
			continue
		}
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
	if limit != nil && (start+*limit) < end {
		end = start + *limit
	}

	return filteredResults[start:end], totalFiltered, nil
}

func (n *nftTokenDbImpl) GetNftsByOwner(franchiseId int64, walletAddress string) ([]*entity.NftTokenModel, error) {
	walletAddress = strings.ToLower(walletAddress)
	results := []*entity.NftTokenModel{}
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

func (n *nftTokenDbImpl) UpdateNftBackstory(tokenId int, collectionAddress string, backstory *string) (*entity.NftTokenModel, error) {
	collectionAddress = strings.ToLower(collectionAddress)

	nftToken := &entity.NftTokenModel{}
	r := n.db.Model(&nftToken).Clauses(clause.Returning{}).Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).Update("backstory", *backstory)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return nftToken, nil
}

func (n *nftTokenDbImpl) UpdateNftOwner(tokenId int, collectionAddress string, ownerAddress string) (*entity.NftTokenModel, error) {
	collectionAddress = strings.ToLower(collectionAddress)
	ownerAddress = strings.ToLower(ownerAddress)

	nftToken := &entity.NftTokenModel{}
	r := n.db.Model(&nftToken).Clauses(clause.Returning{}).Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).Update("owner_address", ownerAddress)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return nftToken, nil
}

func (n *nftTokenDbImpl) CreateNft(nftToken *entity.NftTokenModel) (*entity.NftTokenModel, error) {
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

func (n *nftTokenDbImpl) UpdateNft(nftToken *entity.NftTokenModel) (*entity.NftTokenModel, error) {
	nftToken.CollectionAddress = strings.ToLower(nftToken.CollectionAddress)
	if nftToken.OwnerAddress != nil {
		ownerAddress := strings.ToLower(*nftToken.OwnerAddress)
		nftToken.OwnerAddress = &ownerAddress
	}

	r := n.db.Model(&nftToken).Updates(nftToken)
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

	r := n.db.Where("token_id = ? and collection_address = ?", tokenId, collectionAddress).Delete(&entity.NftTokenModel{})
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to delete the row from db: %v", r.Error)
	}

	return nil
}
