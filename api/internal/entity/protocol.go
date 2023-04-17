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

type CharacterCollection struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StoryCollection struct {
	Address string `json:"address"`
	IsCanon bool   `json:"isCanon"`
}

var FirstFranchise = &Franchise{
	Address:           "0xa75daa0095847af19538d833a718b746849ab509",
	Name:              "First Franchise",
	VaultAddress:      "0xa75daa0095847af19538d833a718b746849ab509",
	CharacterRegistry: "0x33f46c2739c51eda2b19e783830bddea8498b277",
	CharacterContracts: []*CharacterCollection{
		{
			Address: "0xa311d9c50bf955ef37b41233a27b263c93d814b3",
			Name:    "First Franchise Main",
		},
		{
			Address: "0x56cef818ae52fed05f6f61170782a49f9a5fd123",
			Name:    "First Franchise FanFic",
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
	"0xa75daa0095847af19538d833a718b746849ab509": FirstFranchise,
}

var Franchises = []*Franchise{
	FirstFranchise,
}

type ContractType string

var ContractTypes = struct {
	Character ContractType
	Story     ContractType
}{
	Character: "character",
	Story:     "story",
}

type ContractInfo struct {
	Type    ContractType
	IsCanon bool
}

var ContractInfoMap = map[string]*ContractInfo{
	"0xa311d9c50bf955ef37b41233a27b263c93d814b3": {
		Type:    ContractTypes.Character,
		IsCanon: true,
	},
	"0x56cef818ae52fed05f6f61170782a49f9a5fd123": {
		Type:    ContractTypes.Character,
		IsCanon: false,
	},
	"0x4c9c850042a48920c0ead85b6ad9cb2227b9f63c": {
		Type:    ContractTypes.Story,
		IsCanon: true,
	},
	"0x296509d7ee30ffbf12707450136d4bf67b91743a": {
		Type:    ContractTypes.Story,
		IsCanon: false,
	},
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

type Collector struct {
	Address string
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
	Fee      int64  `json:"fee"`
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
