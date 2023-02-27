package entity

import (
	"encoding/json"
	"fmt"

	"github.com/project-nova/backend/api/internal/repository"
)

type NftTokenResponse struct {
	CollectionAddress string              `json:"collectionAddress"`
	TokenId           int                 `json:"tokenId"`
	OwnerAddress      *string             `json:"ownerAddress"`
	Name              *string             `json:"name"`
	Description       *string             `json:"description"`
	ImageUrl          *string             `json:"imageUrl"`
	Image             *string             `json:"image"` // Onchain image
	AnimationUrl      *string             `json:"animationUrl"`
	Traits            []*NftTraitResponse `json:"traits"`
	Backstory         *string             `json:"backstory"`
}

func NewNftTokenResponseFrom(nftModel *repository.NftTokenModel) (*NftTokenResponse, error) {
	if nftModel == nil {
		return nil, fmt.Errorf("input nft token model is nil")
	}

	nftResponse := &NftTokenResponse{
		CollectionAddress: nftModel.CollectionAddress,
		TokenId:           nftModel.TokenId,
		OwnerAddress:      nftModel.OwnerAddress,
		Name:              nftModel.Name,
		Description:       nftModel.Description,
		ImageUrl:          nftModel.ImageUrl,
		Image:             nftModel.Image,
		AnimationUrl:      nftModel.AnimationUrl,
		Backstory:         nftModel.Backstory,
	}

	if nftModel.Traits != nil {
		var traits []*NftTrait
		err := json.Unmarshal([]byte(*nftModel.Traits), &traits)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal traits: %v", err)
		}
		for _, trait := range traits {
			nftTraitResponse := NftTraitResponse(*trait)
			nftResponse.Traits = append(nftResponse.Traits, &nftTraitResponse)
		}
	}

	return nftResponse, nil
}

func (n *NftTokenResponse) GetKeyValuesFromTraits() []KeyValue {
	var keyValues []KeyValue
	for _, trait := range n.Traits {
		keyValues = append(keyValues, KeyValue(trait))
	}
	return keyValues
}

type NftOnchainMeta struct {
	Name         *string     `json:"name"`
	Description  *string     `json:"decription"`
	Image        *string     `json:"image"`
	AnimationUrl *string     `json:"animation_url"`
	Attributes   []*NftTrait `json:"attributes"`
}

// For onchain representation: snake case
type NftTrait struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

// For API response: camel case
type NftTraitResponse struct {
	TraitType string      `json:"traitType"`
	Value     interface{} `json:"value"`
}

func (n *NftTraitResponse) GetKey() string {
	return n.TraitType
}

func (n *NftTraitResponse) GetValue() interface{} {
	return n.Value
}
