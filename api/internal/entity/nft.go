package entity

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
