package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type NftTokensResponse struct {
	Total int                 `json:"total"`
	Data  []*NftTokenResponse `json:"data"`
}

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

// NftTokenModel represents the nft token's model in data storage
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

func (n *NftTokenModel) ToNftTokenResponse() (*NftTokenResponse, error) {
	if n == nil {
		return nil, fmt.Errorf("input nft token model is nil")
	}

	nftResponse := &NftTokenResponse{
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
		var traits []*NftTrait
		err := json.Unmarshal([]byte(*n.Traits), &traits)
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

// UpdateNftBackstoryRequestBody is the request body for UpdateNftBackstoryRequest
type UpdateNftBackstoryRequestBody struct {
	CollectionAddress string `json:"collectionAddress"`
	WalletAddress     string `json:"walletAddress"`
	Message           string `json:"message"`
	Backstory         string `json:"backstory"`
	Signature         string `json:"signature"`
}
