package repository

import (
	"errors"
	"fmt"

	"github.com/project-nova/backend/api/internal/entity"
	"gorm.io/gorm"
)

type CharacterInfoRepository interface {
	CreateCharacter(character *entity.CharacterInfoModel) error
	GetCharacter(franchiseAddress string, characterId int64) (*entity.CharacterInfoModel, error)
	GetCharacters(franchiseAddress string) ([]*entity.CharacterInfoModel, error)
}

func NewCharacterInfoDbImpl(db *gorm.DB) CharacterInfoRepository {
	return &characterInfoDbImpl{
		db: db,
	}
}

type characterInfoDbImpl struct {
	db *gorm.DB
}

func (c *characterInfoDbImpl) CreateCharacter(character *entity.CharacterInfoModel) error {
	r := c.db.Create(character)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if r.Error != nil {
		return fmt.Errorf("failed to insert into db: %v", r.Error)
	}

	return nil
}

func (c *characterInfoDbImpl) GetCharacter(franchiseAddress string, characterId int64) (*entity.CharacterInfoModel, error) {
	result := &entity.CharacterInfoModel{}
	r := c.db.Where("franchise_address = ? and character_id = ?", franchiseAddress, characterId).First(result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return result, nil
}

func (c *characterInfoDbImpl) GetCharacters(franchiseAddress string) ([]*entity.CharacterInfoModel, error) {
	results := []*entity.CharacterInfoModel{}
	r := c.db.Where("franchise_address = ?", franchiseAddress).Order("character_id asc").Find(&results)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, r.Error
	}
	if r.Error != nil {
		return nil, fmt.Errorf("failed to query db: %v", r.Error)
	}

	return results, nil
}
