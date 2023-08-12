package service

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/pkg/logger"
)

type TheGraphService interface {
	GetFranchises() ([]*entity.Franchise, error)
	GetFranchise(franchiseId int64) (*entity.Franchise, error)
	GetCharacters(franchiseId int64) ([]*entity.CharacterInfoModel, error)
	GetStories(franchiseId int64) ([]*entity.StoryInfoV2Model, error)
	GetStory(franchiseId int64, storyId int64) (*entity.StoryInfoV2Model, error)
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
	req := graphql.NewRequest(`
    {
		franchiseRegistereds {
			id
			franchiseId
			owner
			ipAssetRegistry
			name
			tokenURI
		}
    }`)

	ctx := context.Background()
	var franchisesResponse entity.FranchisesTheGraphResponse
	if err := c.client.Run(ctx, req, &franchisesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	var franchises []*entity.Franchise
	for _, franchise := range franchisesResponse.FranchisesRegistered {
		franchises = append(franchises, franchise.ToFranchise())
	}
	return franchises, nil
}

func (c *theGraphServiceImpl) GetFranchise(franchiseId int64) (*entity.Franchise, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt) {
		franchiseRegistereds(where: {franchiseId: $franchiseId}) {
		  id
		  franchiseId
		  owner
		  ipAssetRegistry
		  name
		  tokenURI
		}
	}`)

	req.Var("franchiseId", franchiseId)

	ctx := context.Background()
	var franchiseResponse entity.FranchisesTheGraphResponse
	if err := c.client.Run(ctx, req, &franchiseResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}
	if len(franchiseResponse.FranchisesRegistered) == 0 {
		return nil, fmt.Errorf("failed to find the franchise")
	}

	return franchiseResponse.FranchisesRegistered[0].ToFranchise(), nil
}

func (c *theGraphServiceImpl) GetCharacters(franchiseId int64) ([]*entity.CharacterInfoModel, error) {
	req := graphql.NewRequest(`
	query($franchiseId: BigInt) {
		ipassetCreateds(where: { and: [{ ipAssetType: 2 }, { franchiseId: $franchiseId }]}) {
			id
			franchiseId
			ipAssetId
			owner
			name
			mediaUrl
	  	}
	}`)

	req.Var("franchiseId", franchiseId)

	ctx := context.Background()
	var charactersResponse entity.IpAssetsTheGraphResposne
	if err := c.client.Run(ctx, req, &charactersResponse); err != nil {
		return nil, fmt.Errorf("failed to get the characters from the graph. error: %v", err)
	}

	var characters []*entity.CharacterInfoModel
	for _, character := range charactersResponse.IpAssetCreateds {
		characterInfo, err := character.ToCharacterInfo()
		if err != nil {
			logger.Errorf("Failed to convert the graph character response to character info: %v", err)
			continue
		}
		characters = append(characters, characterInfo)
	}
	return characters, nil
}

func (c *theGraphServiceImpl) GetStories(franchiseId int64) ([]*entity.StoryInfoV2Model, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt) {
		ipassetCreateds(where: { and: [{ ipAssetType: 1 }, { franchiseId: $franchiseId }]}) {
			id
			franchiseId
			ipAssetId
			owner
			name
			mediaUrl
    		transactionHash
	  	}
	}`)

	req.Var("franchiseId", franchiseId)

	ctx := context.Background()
	var storiesResponse entity.IpAssetsTheGraphResposne
	if err := c.client.Run(ctx, req, &storiesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the stories from the graph. error: %v", err)
	}

	var stories []*entity.StoryInfoV2Model
	for _, story := range storiesResponse.IpAssetCreateds {
		storyInfo, err := story.ToStoryInfoV2()
		if err != nil {
			logger.Errorf("Failed to convert the graph story response to story info: %v", err)
			continue
		}
		stories = append(stories, storyInfo)
	}
	return stories, nil
}

func (c *theGraphServiceImpl) GetStory(franchiseId int64, storyId int64) (*entity.StoryInfoV2Model, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt, $storyId: BigInt) {
		ipassetCreateds(where: { and: [{ ipAssetType: 1 }, { franchiseId: $franchiseId }, { ipAssetId: $storyId }]}) {
			id
			franchiseId
			ipAssetId
			owner
			name
			mediaUrl
    		transactionHash
	  	}
	}`)

	req.Var("franchiseId", franchiseId)
	req.Var("storyId", storyId)

	ctx := context.Background()
	var storiesResponse entity.IpAssetsTheGraphResposne
	if err := c.client.Run(ctx, req, &storiesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the stories from the graph. error: %v", err)
	}
	if len(storiesResponse.IpAssetCreateds) == 0 {
		return nil, fmt.Errorf("failed to find the story")
	}

	storyInfo, err := storiesResponse.IpAssetCreateds[0].ToStoryInfoV2()
	if err != nil {
		return nil, fmt.Errorf("failed to convert the graph story response to story info: %v", err)
	}

	return storyInfo, nil
}
