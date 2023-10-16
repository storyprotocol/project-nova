package thegraph

import "github.com/project-nova/backend/api/internal/entity"

type TheGraphServiceKbw interface {
	GetFranchises() ([]*entity.Franchise, error)
	GetFranchise(franchiseId int64) (*entity.Franchise, error)
	GetCharactersByCharacterIds(franchiseId int64, characterIds []string) ([]*entity.CharacterInfoModel, error)
	GetCharacters(franchiseId int64) ([]*entity.CharacterInfoModel, error)
	GetCharacter(franchiseId int64, characterId string) (*entity.CharacterInfoModel, error)
	GetStories(franchiseId int64) ([]*entity.StoryInfoV2Model, error)
	GetStory(franchiseId int64, storyId string) (*entity.StoryInfoV2Model, error)
	GetLicensesByIpAsset(franchiseId int64, ipAssetId string) ([]*entity.License, error)
	GetLicensesByProfile(franchiseId int64, ipAssetId string, walletAddress string) ([]*entity.License, error)
	GetLicense(licenseId int64) (*entity.License, error)
}

type TheGraphServiceMvp interface {
	GetFranchises() ([]*entity.FranchiseMVP, error)
	GetFranchise(franchiseId string) (*entity.FranchiseMVP, error)
	GetIpAssets(franchiseId string) ([]*entity.IpAssetMVP, error)
	GetIpAsset(franchiseId string, ipAssetId string) (*entity.IpAssetMVP, error)
	GetLicenses(franchiseId string, ipAssetId string) ([]*entity.LicenseMVP, error)
	GetLicense(licenseId string) (*entity.LicenseMVP, error)
	GetCollections(franchiseId string) ([]*entity.CollectionMVP, error)
	GetTransactions() ([]*entity.TransactionMVP, error)
	GetTransaction(transactionId string) (*entity.TransactionMVP, error)
}
