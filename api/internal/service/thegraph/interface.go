package thegraph

import (
	v0alpha "github.com/project-nova/backend/api/internal/entity/v0-alpha"
)

type TheGraphServiceAlpha interface {
	GetRelationship(relationshipId string) (*v0alpha.Relationship, error)
	ListRelationships(contract string, tokenId string, options *TheGraphQueryOptions) ([]*v0alpha.Relationship, error)
	ListHooks(moduleId *string, options *TheGraphQueryOptions) ([]*v0alpha.Hook, error)
	GetHook(hookId string) (*v0alpha.Hook, error)
	ListModules(ipOrgId *string, options *TheGraphQueryOptions) ([]*v0alpha.Module, error)
	GetModule(moduleId string) (*v0alpha.Module, error)
	ListIPOrgs(options *TheGraphQueryOptions) ([]*v0alpha.IPOrg, error)
	GetIPOrg(iporgId string) (*v0alpha.IPOrg, error)
	ListIPAssets(iporgId *string, options *TheGraphQueryOptions) ([]*v0alpha.IPAsset, error)
	GetIPAsset(ipAssetId string) (*v0alpha.IPAsset, error)
	ListTransactions(ipOrgId *string, options *TheGraphQueryOptions) ([]*v0alpha.Transaction, error)
	GetTransaction(transactionId string) (*v0alpha.Transaction, error)
	ListLicenses(iporgId *string, ipAssetId *string, options *TheGraphQueryOptions) ([]*v0alpha.License, error)
	GetLicense(licenseId string) (*v0alpha.License, error)
	GetRelationshipType(relType string, ipOrgId string) (*v0alpha.RelationshipType, error)
	ListRelationshipTypes(ipOrgId *string, options *TheGraphQueryOptions) ([]*v0alpha.RelationshipType, error)
	ListLicenseParams(ipOrgId string, options *TheGraphQueryOptions) ([]*v0alpha.IpOrgLicenseParam, error)
}

type TheGraphQueryOptions struct {
	First          int
	Skip           int
	OrderBy        string
	OrderDirection string
}

func FromRequestQueryOptions(options *v0alpha.QueryOptions) *TheGraphQueryOptions {
	if options == nil {
		return &TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	if options.Pagination.Limit == 0 {
		options.Pagination.Limit = 100
	}

	return &TheGraphQueryOptions{
		First: options.Pagination.Limit,
		Skip:  options.Pagination.Offset,
	}
}
