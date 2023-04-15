package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/project-nova/backend/pkg/model"
)

type Franchise struct {
	Address            string                 `json:"address"`
	Name               string                 `json:"name"`
	VaultAddress       string                 `json:"vaultAddress"`
	CharacterRegistry  string                 `json:"characterRegistry"`
	CharacterContracts []*CharacterCollection `json:"characterContract"`
	StoryRegistry      string                 `json:"storyRegistry"`
	StoryContracts     []*StoryCollection     `json:"storyContract"`
	LicenseRepository  string                 `json:"licenseRepository"`
	LicenseRegistry    string                 `json:"licenseRegistry"`
}

var SingleFranchise = &Franchise{
	Address:           "0x624bdd3b5d4f67fef15880a8f2cb0a0703d6ec0c",
	Name:              "Forced Offline",
	VaultAddress:      "0x624bdd3b5d4f67fef15880a8f2cb0a0703d6ec0c",
	CharacterRegistry: "0x402d3e9c136dfefaae3b04ea086632688951eccd",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0xf90bf1f50b71baae6bedcaf92a3a63d97200382c",
			Name:    "Force Offline Main",
		},
		{
			Address: "0xd7431ef1cbd4b5bcc568def48d4480f4e08d2224",
			Name:    "Forced offline FanFic",
		},
	},
	StoryRegistry: "0x13687aafb5accecb358f207f58614e01b09dec72",
	StoryContracts: []*StoryCollection{
		{
			Address: "0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca",
			IsCanon: true,
		},
		{
			Address: "0x72f822e6b4a752b11b4275aad841b86a3f5266ab",
			IsCanon: false,
		},
	},
}

var NewFranchise = &Franchise{
	Address:           "0xa75daa0095847af19538d833a718b746849ab509",
	Name:              "Forced Offline",
	VaultAddress:      "0xa75daa0095847af19538d833a718b746849ab509",
	CharacterRegistry: "0x33f46c2739c51eda2b19e783830bddea8498b277",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0xa311d9c50bf955ef37b41233a27b263c93d814b3",
			Name:    "Force Offline Main",
		},
		{
			Address: "0x56cef818ae52fed05f6f61170782a49f9a5fd123",
			Name:    "Forced offline FanFic",
		},
	},
	StoryRegistry: "0x73494f208375403015d9ae0a72154c4a67b11552",
	StoryContracts: []*StoryCollection{
		{
			Address: "0x4c9c850042a48920c0ead85b6ad9cb2227b9f63c",
			IsCanon: true,
		},
		{
			Address: "0x296509d7ee30ffbf12707450136d4bf67b91743a",
			IsCanon: false,
		},
	},
	LicenseRepository: "0x48bfe2765f85076ce7d4de390f258dd53ca3db66",
	LicenseRegistry:   "0x1b886b15c5d87c0918b5d93bc3503fe9d866d9be",
}

var FranchiseMap = map[string]*Franchise{
	"0x624bdd3b5d4f67fef15880a8f2cb0a0703d6ec0c": SingleFranchise,
	"0xa75daa0095847af19538d833a718b746849ab509": NewFranchise,
}

var CharacterContractMap = map[string]*CharacterCollection{
	"0xf90bf1f50b71baae6bedcaf92a3a63d97200382c": {
		Address: "0xf90bf1f50b71baae6bedcaf92a3a63d97200382c",
		Name:    "Force Offline Main",
	},
	"0xd7431ef1cbd4b5bcc568def48d4480f4e08d2224": {
		Address: "0xd7431ef1cbd4b5bcc568def48d4480f4e08d2224",
		Name:    "Forced offline FanFic",
	},
	"0xa311d9c50bf955ef37b41233a27b263c93d814b3": {
		Address: "0xf90bf1f50b71baae6bedcaf92a3a63d97200382c",
		Name:    "Force Offline Main",
	},
	"0x56cef818ae52fed05f6f61170782a49f9a5fd123": {
		Address: "0xd7431ef1cbd4b5bcc568def48d4480f4e08d2224",
		Name:    "Forced offline FanFic",
	},
}

var StoryContractMap = map[string]*StoryCollection{
	"0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca": {
		Address: "0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca",
		IsCanon: true,
	},
	"0x72f822e6b4a752b11b4275aad841b86a3f5266ab": {
		Address: "0x72f822e6b4a752b11b4275aad841b86a3f5266ab",
		IsCanon: false,
	},
	"0x4c9c850042a48920c0ead85b6ad9cb2227b9f63c": {
		Address: "0x6a84bcceebcd42ee0d73fa95f75862f635fd96ca",
		IsCanon: true,
	},
	"0x296509d7ee30ffbf12707450136d4bf67b91743a": {
		Address: "0x72f822e6b4a752b11b4275aad841b86a3f5266ab",
		IsCanon: false,
	},
}

var Franchises = []*Franchise{
	SingleFranchise,
	NewFranchise,
	/*
		{
			Address:           "0xWhiteFountain",
			Name:              "White Fountain",
			VaultAddress:      "0xWhiteFountainVault",
			CharacterRegistry: "",
			CharacterContracts: []*CharacterCollection{
				{
					Address: "0xWhiteFountainCharacter",
					Name:    "White Fountain Main",
				},
				{
					Address: "0xWhiteFountainFanFicCharacter",
					Name:    "White Fountain FanFic",
				},
			},
			StoryRegistry: "",
			StoryContracts: []*StoryCollection{
				{
					Address: "0xWhiteFountainStory",
					IsCanon: true,
				},
				{
					Address: "0xWhiteFountainFanfic",
					IsCanon: false,
				},
			},
		},
	*/
}

type CharacterCollection struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StoryCollection struct {
	Address string `json:"address"`
	IsCanon bool   `json:"isCanon"`
}

type Character struct {
	CollectionAddress string              `json:"collectionAddress"`
	TokenId           int                 `json:"tokenId"`
	AuthorAddress     string              `json:"authorAddress"`
	OwnerAddress      string              `json:"ownerAddress"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	ImageUrl          string              `json:"imageUrl"`
	Traits            []*NftTraitResponse `json:"traits"`
}

var SingleCharacter = &Character{
	CollectionAddress: "0xForcedOfflineCharacter",
	TokenId:           1,
	AuthorAddress:     "0xJosephEvans",
	OwnerAddress:      "jason.eth",
	Name:              "Rayze",
	Description:       "Rayze is the main character of force offline",
	ImageUrl:          "https://ipfs.io/ipfs/QmNvRAgJbWgcCvY2cJgvqvfgu6SwRKr5zsbZufm3owpaEc/images/145.png",
	Traits: []*NftTraitResponse{
		{
			TraitType: "hairColor",
			Value:     "red",
		},
	},
}

var Characters = []*Character{
	SingleCharacter,
	{
		CollectionAddress: "0xForcedOfflineCharacter",
		TokenId:           2,
		AuthorAddress:     "0xJosephEvans",
		OwnerAddress:      "allen.eth",
		Name:              "Myza",
		Description:       "Myza is the girlfriend of Rayze",
		ImageUrl:          "https://ipfs.io/ipfs/QmNvRAgJbWgcCvY2cJgvqvfgu6SwRKr5zsbZufm3owpaEc/images/11.png",
		Traits: []*NftTraitResponse{
			{
				TraitType: "hairColor",
				Value:     "purple",
			},
		},
	},
}

type Story struct {
	CollectionAddress string       `json:"collectionAddress"`
	TokenId           int          `json:"tokenId"`
	AuthorAddress     []string     `json:"authorAddress"`
	OwnerAddress      string       `json:"ownerAddress"`
	Title             string       `json:"title"`
	ContentUrl        string       `json:"contentUrl"`
	ImageUrl          string       `json:"imageUrl"`
	Characters        []*Character `json:"characters"`
	IsCanon           bool         `json:"isCanon"`
}

var SingleStory = &Story{
	CollectionAddress: "0xForcedOfflineStory",
	TokenId:           1,
	AuthorAddress:     []string{"0xJosephEvans", "jasonlevy.eth"},
	OwnerAddress:      "0xJosephEvans",
	Title:             "Forced Offline",
	ContentUrl:        "https://stag.api.storyprotocol.net/v1/story/1/1/",
	Characters: []*Character{
		{
			CollectionAddress: "0xForcedOfflineCharacter",
			TokenId:           1,
			AuthorAddress:     "0xJosephEvans",
			OwnerAddress:      "jason.eth",
			Name:              "Rayze",
			Description:       "Rayze is the main character of force offline",
			Traits: []*NftTraitResponse{
				{
					TraitType: "hairColor",
					Value:     "red",
				},
			},
		},
		{
			CollectionAddress: "0xForcedOfflineCharacter",
			TokenId:           2,
			AuthorAddress:     "0xJosephEvans",
			OwnerAddress:      "allen.eth",
			Name:              "Myza",
			Description:       "Myza is the girlfriend of Rayze",
			Traits: []*NftTraitResponse{
				{
					TraitType: "hairColor",
					Value:     "purple",
				},
			},
		},
	},
	IsCanon: true,
}

var Stories = []*Story{
	SingleStory,
}

type Collector struct {
	Address string
}

var Collectors = []*Collector{
	{
		Address: "allen.eth",
	},
	{
		Address: "leo.eth",
	},
}

type CharacterNftOnchainMeta struct {
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	Image       *string     `json:"image"`
	Author      *string     `json:"author"`
	Attributes  []*NftTrait `json:"attributes"`
}

type StoryNftOnchainMeta struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
}

type ProtocolStoryContentModel struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	ContentJson string    `json:"contentJson"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (ProtocolStoryContentModel) TableName() string {
	return "story_content"
}

func (p *ProtocolStoryContentModel) ToStoryContentModel() (*model.StoryContentModel, error) {
	var model *model.StoryContentModel
	err := json.Unmarshal([]byte(p.ContentJson), &model)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal content json: %v", err)
	}

	return model, nil
}

type UploadProtocolStoryRequestBody struct {
	Text string `json:"text"`
}

func (u *UploadProtocolStoryRequestBody) ToProtocolContentModel() (*ProtocolStoryContentModel, error) {
	model := &model.StoryContentModel{
		Content: []*model.StorySectionModel{
			{
				Type: "paragraph",
				Data: []*model.StoryMediaModel{
					{
						Type:    "text",
						Content: u.Text,
					},
				},
			},
		},
	}

	modelBytes, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the story model: %v", err)
	}

	return &ProtocolStoryContentModel{
		ID:          uuid.New().String(),
		ContentJson: string(modelBytes),
	}, nil
}

type LicenseResponse struct {
	Right *LicenseInfo `json:"right"`
	Nfts  []*NftInfo   `json:"nfts"`
}

type LicenseInfo struct {
	Type     string `json:"type"`
	Term     string `json:"term"`
	Fee      string `json:"fee"`
	Currency string `json:"currency"`
}

type NftInfo struct {
	Address string `json:"address"`
	TokenId int    `json:"tokenId"`
}

var LicenseRightsMap = map[uint8]string{
	0: "Unset",
	1: "TheatricalRelease",
	2: "HomeVideo",
	3: "Streaming",
	4: "TV",
	5: "Merchandising",
	6: "ComicBook",
	7: "Game",
	8: "SpinOff",
}
