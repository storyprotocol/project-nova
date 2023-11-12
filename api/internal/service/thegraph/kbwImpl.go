package thegraph

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/entity"
	"github.com/project-nova/backend/pkg/logger"
)

func NewTheGraphServiceKbwImpl(client *graphql.Client) TheGraphServiceKbw {
	return &theGraphServiceKbwImpl{
		client: client,
	}
}

type theGraphServiceKbwImpl struct {
	client *graphql.Client
}

func (c *theGraphServiceKbwImpl) GetFranchises() ([]*entity.Franchise, error) {
	req := graphql.NewRequest(`
    {
		franchiseRegistereds {
			id
			franchiseId
			owner
			ipAssetRegistry
			name
			tokenURI
			transactionHash
		}
    }`)

	ctx := context.Background()
	var franchisesResponse entity.FranchisesTheGraphResponse
	if err := c.client.Run(ctx, req, &franchisesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	franchises := []*entity.Franchise{}
	for _, franchise := range franchisesResponse.FranchisesRegistered {
		franchises = append(franchises, franchise.ToFranchise())
	}

	return franchises, nil
}

func (c *theGraphServiceKbwImpl) GetFranchise(franchiseId int64) (*entity.Franchise, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt) {
		franchiseRegistereds(where: {franchiseId: $franchiseId}) {
		  id
		  franchiseId
		  owner
		  ipAssetRegistry
		  name
		  tokenURI
		  transactionHash
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

func (c *theGraphServiceKbwImpl) GetCharactersByCharacterIds(franchiseId int64, characterIds []string) ([]*entity.CharacterInfoModel, error) {
	req := graphql.NewRequest(`
	query($franchiseId: BigInt, $characterIds: [String]) {
		ipassetCreateds(where: { and: [{ ipAssetType: 2 }, { franchiseId: $franchiseId }, { ipAssetId_in: $characterIds }]}) {
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
	req.Var("characterIds", characterIds)

	ctx := context.Background()
	var charactersResponse entity.IpAssetsTheGraphResposne
	if err := c.client.Run(ctx, req, &charactersResponse); err != nil {
		return nil, fmt.Errorf("failed to get the characters from the graph. error: %v", err)
	}

	characters := []*entity.CharacterInfoModel{}
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

func (c *theGraphServiceKbwImpl) GetCharacters(franchiseId int64) ([]*entity.CharacterInfoModel, error) {
	req := graphql.NewRequest(`
	query($franchiseId: BigInt) {
		ipassetCreateds(where: { and: [{ ipAssetType: 2 }, { franchiseId: $franchiseId }]}) {
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
	var charactersResponse entity.IpAssetsTheGraphResposne
	if err := c.client.Run(ctx, req, &charactersResponse); err != nil {
		return nil, fmt.Errorf("failed to get the characters from the graph. error: %v", err)
	}

	characters := []*entity.CharacterInfoModel{}
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

func (c *theGraphServiceKbwImpl) GetCharacter(franchiseId int64, characterId string) (*entity.CharacterInfoModel, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt, $characterId: BigInt) {
		ipassetCreateds(where: { and: [{ ipAssetType: 2 }, { franchiseId: $franchiseId }, { ipAssetId: $characterId }]}) {
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
	req.Var("characterId", characterId)

	ctx := context.Background()
	var charactersResponse entity.IpAssetsTheGraphResposne
	if err := c.client.Run(ctx, req, &charactersResponse); err != nil {
		return nil, fmt.Errorf("failed to get the characters from the graph. error: %v", err)
	}
	if len(charactersResponse.IpAssetCreateds) == 0 {
		return nil, fmt.Errorf("failed to find the character")
	}

	characterInfo, err := charactersResponse.IpAssetCreateds[0].ToCharacterInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to convert the graph character response to character info: %v", err)
	}

	return characterInfo, nil
}

func (c *theGraphServiceKbwImpl) GetStories(franchiseId int64) ([]*entity.StoryInfoV2Model, error) {
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

	stories := []*entity.StoryInfoV2Model{}
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

func (c *theGraphServiceKbwImpl) GetStory(franchiseId int64, storyId string) (*entity.StoryInfoV2Model, error) {
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

func (c *theGraphServiceKbwImpl) GetLicensesByIpAsset(franchiseId int64, ipAssetId string) ([]*entity.License, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt, $ipAssetId: BigInt) {
		licenseGranteds(where: { and: [{ franchiseId: $franchiseId }, { storyId: $ipAssetId }]}) {
			id
			licenseId
			franchiseId
			storyId
			storyName
			owner
			scope
			duration
			right
			imageUrl
			collectionAddress
	  	}
	}`)

	req.Var("franchiseId", franchiseId)
	req.Var("ipAssetId", ipAssetId)

	ctx := context.Background()
	var licensesResponse entity.LicensesTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the licenses from the graph. error: %v", err)
	}

	licenses := []*entity.License{}
	for _, license := range licensesResponse.LicensesGranted {
		licenseModel, err := license.ToLicense()
		if err != nil {
			logger.Errorf("Failed to convert the graph license response to license model: %v", err)
			continue
		}
		licenses = append(licenses, licenseModel)
	}
	return licenses, nil
}

func (c *theGraphServiceKbwImpl) GetLicensesByProfile(franchiseId int64, ipAssetId string, walletAddress string) ([]*entity.License, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt, $ipAssetId: BigInt, $walletAddress: String) {
		licenseGranteds(where: { and: [{ franchiseId: $franchiseId }, { storyId: $ipAssetId }, { owner: $walletAddress }]}) {
			id
			licenseId
			franchiseId
			storyId
			storyName
			owner
			scope
			duration
			right
			imageUrl
			collectionAddress
	  	}
	}`)

	req.Var("franchiseId", franchiseId)
	req.Var("ipAssetId", ipAssetId)
	req.Var("walletAddress", walletAddress)

	ctx := context.Background()
	var licensesResponse entity.LicensesTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the licenses from the graph. error: %v", err)
	}

	licenses := []*entity.License{}
	for _, license := range licensesResponse.LicensesGranted {
		licenseModel, err := license.ToLicense()
		if err != nil {
			logger.Errorf("Failed to convert the graph license response to license model: %v", err)
			continue
		}
		licenses = append(licenses, licenseModel)
	}
	return licenses, nil
}

func (c *theGraphServiceKbwImpl) GetLicense(licenseId int64) (*entity.License, error) {
	req := graphql.NewRequest(`
    query($licenseId: BigInt) {
		licenseGranteds(where: {licenseId: $licenseId}) {
			id
			licenseId
			franchiseId
			storyId
			storyName
			owner
			scope
			duration
			right
			imageUrl
			collectionAddress
	  	}
	}`)

	req.Var("licenseId", licenseId)

	ctx := context.Background()
	var licensesResponse entity.LicensesTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the licenses from the graph. error: %v", err)
	}
	if len(licensesResponse.LicensesGranted) == 0 {
		return nil, fmt.Errorf("failed to find the license")
	}

	licenseModel, err := licensesResponse.LicensesGranted[0].ToLicense()
	if err != nil {
		return nil, fmt.Errorf("failed to convert the graph license response to license model: %v", err)
	}

	return licenseModel, nil
}