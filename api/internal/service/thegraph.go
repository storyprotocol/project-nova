package service

import (
	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/entity"
)

type TheGraphService interface {
	GetFranchises() ([]*entity.FranchiseInfoModel, error)
	GetFranchise(franchiseId int64) (*entity.FranchiseInfoModel, error)
	GetCharacters(franchiseId int64) ([]*entity.CharacterInfoModel, error)
	GetStories() ([]*entity.StoryInfoV2Model, error)
	GetStory(franchiseId int64) (*entity.StoryInfoV2Model, error)
}

func NewTheGraphServiceImpl(client *graphql.Client) TheGraphService {
	return &theGraphServiceImpl{
		client: client,
	}
}

type theGraphServiceImpl struct {
	client *graphql.Client
}

func (c *theGraphServiceImpl) GetFranchises() ([]*entity.Franchise, error) {
	return nil, nil
}

func (c *theGraphServiceImpl) GetFranchise(franchiseId int64) (*entity.Franchise, error) {
	return nil, nil
}

func (c *theGraphServiceImpl) GetCharacters(franchiseId int64) ([]*entity.CharacterInfoModel, error) {
	return nil, nil
}

func (c *theGraphServiceImpl) GetStories() ([]*entity.StoryInfoV2Model, error) {
	return nil, nil
}

func (c *theGraphServiceImpl) GetStory(franchiseId int64) (*entity.StoryInfoV2Model, error) {
	return nil, nil
}
