package entity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CharacterMetadata struct {
	Name      string `json:"name"`
	Backstory string `json:"backstory"`
}

type CreateCharacterResp struct {
	MediaUri string `json:"arweaveURI"`
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
	return resp
}
