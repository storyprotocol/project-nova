package entity

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CharacterMetadata struct {
	Name      string  `json:"name"`
	ImageUrl  *string `json:"image"`
	Backstory string  `json:"backstory"`
}

type CreateCharacterResp struct {
	CharacterMediaUri string `json:"characterUrl"`
	BackstoryMediaUri string `json:"backstoryUrl"`
}

type GetCharactersResp struct {
	Name         string  `json:"name"`
	OwnerAddress string  `json:"owner"`
	CharacterId  *int64  `json:"tokenId"`
	ImageUrl     *string `json:"img"`
}

type GetCharacterResp struct {
	Name         string  `json:"name"`
	OwnerAddress string  `json:"owner"`
	CharacterId  *int64  `json:"tokenId"`
	ImageUrl     *string `json:"img"`
	Backstory    *string `json:"backstory"`
	MediaUri     *string `json:"arweaveURI"`
	Txhash       *string `json:"txhash"`
}

type CreateCharacterRequestBody struct {
	CharacterName *string `json:"name"`
	ImageUrl      *string `json:"img"`
	OwnerAddress  *string `json:"owner"`
	Backstory     *string `json:"backstory"`
}

func (c *CreateCharacterRequestBody) Validate() error {
	if c.CharacterName == nil {
		return fmt.Errorf("character name is empty")
	}
	if c.ImageUrl == nil {
		return fmt.Errorf("image url is empty")
	}
	if c.OwnerAddress == nil {
		return fmt.Errorf("owner address is empty")
	}
	if c.Backstory == nil {
		return fmt.Errorf("backstory is empty")
	}
	return nil
}

func (c *CreateCharacterRequestBody) ToCharacterMetadata() *CharacterMetadata {
	characterMeta := &CharacterMetadata{}
	if c.CharacterName != nil {
		characterMeta.Name = *c.CharacterName
	}
	if c.Backstory != nil {
		characterMeta.Backstory = *c.Backstory
	}
	if c.ImageUrl != nil {
		characterMeta.ImageUrl = c.ImageUrl
	}

	return characterMeta
}

func (c *CreateCharacterRequestBody) ToCharacterInfoModel() *CharacterInfoModel {
	characterInfo := &CharacterInfoModel{
		ID: uuid.New().String(),
	}
	// Temporary: Add character id for FE testing
	characterId := int64(rand.Uint32())
	characterInfo.CharacterId = &characterId

	if c.CharacterName != nil {
		characterInfo.CharacterName = *c.CharacterName
	}
	if c.ImageUrl != nil {
		characterInfo.ImageUrl = c.ImageUrl
	}
	if c.OwnerAddress != nil {
		characterInfo.OwnerAddress = *c.OwnerAddress
	}
	if c.Backstory != nil {
		characterInfo.Backstory = c.Backstory
	}

	return characterInfo
}

type CharacterInfoModel struct {
	ID            string    `gorm:"primaryKey;column:id" json:"id"`
	FranchiseId   int64     `json:"franchiseId"`
	CharacterId   *int64    `json:"characterId"`
	CharacterName string    `json:"characterName"`
	OwnerAddress  string    `json:"ownerAddress"`
	ImageUrl      *string   `json:"imageUrl"`
	Backstory     *string   `json:"backstory"`
	MediaUri      *string   `json:"mediaUri"`
	Txhash        *string   `json:"txhash"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (CharacterInfoModel) TableName() string {
	return "character_info"
}

func (c *CharacterInfoModel) ToGetCharactersResp() (*GetCharactersResp, error) {
	if c.CharacterId == nil {
		return nil, fmt.Errorf("character id is nil")
	}
	resp := &GetCharactersResp{
		Name:         c.CharacterName,
		OwnerAddress: c.OwnerAddress,
		CharacterId:  c.CharacterId,
	}
	if c.ImageUrl != nil {
		resp.ImageUrl = c.ImageUrl
	}
	return resp, nil
}

func (c *CharacterInfoModel) ToGetCharacterResp() *GetCharacterResp {
	resp := &GetCharacterResp{
		Name:         c.CharacterName,
		OwnerAddress: c.OwnerAddress,
	}
	if c.CharacterId != nil {
		resp.CharacterId = c.CharacterId
	}
	if c.ImageUrl != nil {
		resp.ImageUrl = c.ImageUrl
	}
	if c.Backstory != nil {
		resp.Backstory = c.Backstory
	}
	if c.MediaUri != nil {
		resp.MediaUri = c.MediaUri
	}
	if c.Txhash != nil {
		resp.Txhash = c.Txhash
	}

	return resp
}

type IpAssetsTheGraphResposne struct {
	IpAssetCreateds []*IpAssetTheGraph `json:"ipassetCreateds"`
}

type IpAssetTheGraph struct {
	ID          string `json:"id"`
	FranchiseId string `json:"franchiseId"`
	IpAssetId   string `json:"ipAssetId"`
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	MediaUrl    string `json:"mediaUrl"`
	TxHash      string `json:"transactionHash"`
}

func (f *IpAssetTheGraph) ToCharacterInfo() (*CharacterInfoModel, error) {
	franchiseId, err := strconv.ParseInt(f.FranchiseId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. franchise id: %s is not a valid int64. %v", f.FranchiseId, err)
	}

	ipAssetId, err := strconv.ParseInt(f.IpAssetId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. ipAssetId: %s is not a valid int64. %v", f.IpAssetId, err)
	}

	character := &CharacterInfoModel{
		FranchiseId:   franchiseId,
		CharacterId:   &ipAssetId,
		OwnerAddress:  f.Owner,
		CharacterName: f.Name,
		MediaUri:      &f.MediaUrl,
		Txhash:        &f.TxHash,
	}
	return character, err
}
