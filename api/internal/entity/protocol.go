package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/project-nova/backend/pkg/model"
)

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

// TODO: This is a temporary hack since protocol doesn't provide this info.
// Down the road, this should be fetched from the protocol
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
