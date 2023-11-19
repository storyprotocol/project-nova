package thegraph

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	v0alpha "github.com/project-nova/backend/api/internal/entity/v0-alpha"
)

func NewTheGraphServiceAlphaImpl(client *graphql.Client) TheGraphServiceAlpha {
	return &theGraphServiceAlphaImpl{
		client: client,
	}
}

type theGraphServiceAlphaImpl struct {
	client *graphql.Client
}

func (s *theGraphServiceAlphaImpl) ListRelationships(contract string, tokenId string, options *v0alpha.QueryOptions) ([]*v0alpha.Relationship, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(`
		query($contract: String, $tokenId: String, $first: Int, $skip: Int) {
			relationships(where: {contract: $contract, tokenId: $tokenId}, first: $first, skip: $skip) {
				id
				contract
				tokenId
				owner
				relatedTokenId
				relationshipType
			}
		}
	`)
	req.Var("contract", contract)
	req.Var("tokenId", tokenId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var relationshipsResponse v0alpha.RelationshipTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &relationshipsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the relationships from the graph. error: %v", err)
	}

	return relationshipsResponse.ToRelationships(), nil
}

func (s *theGraphServiceAlphaImpl) ListHooks(moduleId *string, options *v0alpha.QueryOptions) ([]*v0alpha.Hook, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(`
		query($moduleId: String, $first: Int, $skip: Int) {
			hooks(where: {moduleId: $moduleId}, first: $first, skip: $skip) {
				id
				moduleId
				hookType
				hookData
			}
		}
	`)
	req.Var("moduleId", *moduleId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var hooksResponse v0alpha.HookTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &hooksResponse); err != nil {
		return nil, fmt.Errorf("failed to get the hooks from the graph. error: %v", err)
	}

	return hooksResponse.ToHooks(), nil
}

func (s *theGraphServiceAlphaImpl) GetHook(hookId string) (*v0alpha.Hook, error) {
	req := graphql.NewRequest(`
		query($hookId: String) {
			hooks(where: {id: $hookId}) {
				id
				moduleId
				hookType
				hookData
			}
		}
	`)
	req.Var("hookId", hookId)

	ctx := context.Background()
	var hooksResponse v0alpha.HookTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &hooksResponse); err != nil {
		return nil, fmt.Errorf("failed to get the hook from the graph. error: %v", err)
	}

	if len(hooksResponse.Hooks) == 0 {
		return nil, fmt.Errorf("failed to find the hook")
	}

	return hooksResponse.Hooks[0].ToHook(), nil
}

func (s *theGraphServiceAlphaImpl) ListModules(ipOrgId *string, options *v0alpha.QueryOptions) ([]*v0alpha.Module, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(`
		query($ipOrgId: String, $first: Int, $skip: Int) {
			modules(where: {ipOrgId: $ipOrgId}, first: $first, skip: $skip) {
				id
				ipOrgId
				moduleType
				moduleData
			}
		}
	`)
	req.Var("ipOrgId", *ipOrgId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var modulesResponse v0alpha.ModuleTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &modulesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the modules from the graph. error: %v", err)
	}

	return modulesResponse.ToModules(), nil
}

func (s *theGraphServiceAlphaImpl) GetModule(moduleId string) (*v0alpha.Module, error) {
	req := graphql.NewRequest(`
		query($moduleId: String) {
			modules(where: {id: $moduleId}) {
				id
				ipOrgId
				moduleType
				moduleData
			}
		}
	`)
	req.Var("moduleId", moduleId)

	ctx := context.Background()
	var modulesResponse v0alpha.ModuleTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &modulesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the module from the graph. error: %v", err)
	}

	if len(modulesResponse.Modules) == 0 {
		return nil, fmt.Errorf("failed to find the module")
	}

	return modulesResponse.Modules[0].ToModule(), nil
}

func (s *theGraphServiceAlphaImpl) ListIPOrgs(options *v0alpha.QueryOptions) ([]*v0alpha.IPOrg, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(`
		query($first: Int, $skip: Int) {
			iporgRegistereds(first: $first, skip: $skip) {
				id
				owner
				ipOrgId
				name
				symbol
				ipAssetTypes
				baseURI
				contractURI
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var ipOrgsResponse v0alpha.IpOrgTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &ipOrgsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	return ipOrgsResponse.ToIPOrgs(), nil
}

func (s *theGraphServiceAlphaImpl) GetIPOrg(iporgId string) (*v0alpha.IPOrg, error) {
	req := graphql.NewRequest(`
		query($iporgId: BigInt) {
			iporgRegistereds(where: {id: $iporgId}) {
				id
				owner
				ipOrgId
				name
				symbol
				ipAssetTypes
				baseURI
				contractURI
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`)
	req.Var("iporgId", iporgId)

	ctx := context.Background()
	var ipOrgsResponse v0alpha.IpOrgTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &ipOrgsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	if len(ipOrgsResponse.IporgRegistereds) == 0 {
		return nil, fmt.Errorf("failed to find the iporg")
	}

	return ipOrgsResponse.IporgRegistereds[0].ToIPOrg(), nil
}

func (s *theGraphServiceAlphaImpl) ListIPAssets(iporgId *string, options *v0alpha.QueryOptions) ([]*v0alpha.IPAsset, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(`
		query($iporgId: BigInt, $first: Int, $skip: Int) {
			ipassetRegistereds(where: {ipOrgId: $iporgId}, first: $first, skip: $skip) {
				id
				ipOrgId
				ipAssetId
				ipOrgAssetId
				owner
				name
				ipAssetType
				contentHash
				mediaUrl
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`)
	req.Var("iporgId", iporgId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var ipAssetsResponse v0alpha.IpAssetTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &ipAssetsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}
	return ipAssetsResponse.ToIPAssets(), nil
}

func (s *theGraphServiceAlphaImpl) GetIPAsset(iporgId string, ipAssetId string) (*v0alpha.IPAsset, error) {
	req := graphql.NewRequest(`
		query($iporgId: BigInt, $ipAssetId: BigInt) {
			iporgRegistereds(where: {id: $ipAssetId, ipOrgId: $iporgId}) {
				id
				ipAssetId
				ipOrgId
				ipOrgAssetId
				owner
				name
				ipAssetType
				contentHash
				mediaUrl
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`)
	req.Var("iporgId", iporgId)
	req.Var("ipAssetId", ipAssetId)

	ctx := context.Background()
	var ipAssetsResponse v0alpha.IpAssetTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &ipAssetsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	if len(ipAssetsResponse.IpassetRegistereds) == 0 {
		return nil, fmt.Errorf("failed to find the iporg")
	}

	return ipAssetsResponse.IpassetRegistereds[0].ToIPAsset(), nil
}

func (s *theGraphServiceAlphaImpl) ListTransactions(ipOrgId *string, options *v0alpha.QueryOptions) ([]*v0alpha.Transaction, error) {
	// TODO(Rex): support optional iporgid
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(`
		query($iporgId: BigInt, $first: Int, $skip: Int) {
			transactions(where: {ipOrgId: $iporgId}, first: $first, skip: $skip) {
				id
				ipOrgId
				ipAssetId
				transactionHash
				blockNumber
				blockTimestamp
			}
		}
	`)
	req.Var("iporgId", ipOrgId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var transactionsResponse v0alpha.TransactionTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &transactionsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the transactions from the graph. error: %v", err)
	}

	return transactionsResponse.ToTransactions(), nil
}

func (s *theGraphServiceAlphaImpl) GetTransaction(transactionId string) (*v0alpha.Transaction, error) {
	req := graphql.NewRequest(`

		query($transactionId: BigInt) {
			transactions(where: {id: $transactionId}) {
				id
				ipOrgId
				ipAssetId
				transactionHash
				blockNumber
				blockTimestamp
			}
		}
	`)
	req.Var("transactionId", transactionId)

	ctx := context.Background()
	var transactionsResponse v0alpha.TransactionTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &transactionsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the transactions from the graph. error: %v", err)
	}

	if len(transactionsResponse.Transactions) == 0 {
		return nil, fmt.Errorf("failed to find the transaction")
	}

	return transactionsResponse.Transactions[0].ToTransaction(), nil
}

func (s *theGraphServiceAlphaImpl) ListLicenses(iporgId *string, ipAssetId *string, options *v0alpha.QueryOptions) ([]*v0alpha.License, error) {
	options = s.setQueryOptions(options)

	req := graphql.NewRequest(`
		query($iporgId: String, $ipAssetId: String, $first: Int, $skip: Int) {
			licenses(where: {ipOrgId: $iporgId, ipAssetId: $ipAssetId}, first: $first, skip: $skip) {
				id
				ipOrgId
				ipAssetId
				licenseType
				issuedAt
				expiredAt
			}
		}
	`)
	req.Var("iporgId", iporgId)
	req.Var("ipAssetId", ipAssetId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)

	ctx := context.Background()
	var licensesResponse v0alpha.LicenseTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &licensesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the licenses from the graph. error: %v", err)
	}

	return licensesResponse.ToLicenses(), nil
}

func (s *theGraphServiceAlphaImpl) GetLicense(licenseId string) (*v0alpha.License, error) {
	req := graphql.NewRequest(`
		query($licenseId: String) {
			license(id: $licenseId) {
				id
				ipOrgId
				ipAssetId
				licenseType
				issuedAt
				expiredAt
			}
		}
	`)
	req.Var("licenseId", licenseId)

	ctx := context.Background()
	var licenseResponse v0alpha.LicenseTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &licenseResponse); err != nil {
		return nil, fmt.Errorf("failed to get the license from the graph. error: %v", err)
	}

	if len(licenseResponse.LicenseRegistereds) == 0 {
		return nil, fmt.Errorf("license not found")
	}

	return licenseResponse.LicenseRegistereds[0].ToLicense(), nil
}

func (s *theGraphServiceAlphaImpl) setQueryOptions(options *v0alpha.QueryOptions) *v0alpha.QueryOptions {
	if options == nil {
		options = &v0alpha.QueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	if options.First == 0 {
		options.First = 100
	}

	return options
}
