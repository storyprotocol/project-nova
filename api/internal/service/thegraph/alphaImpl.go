package thegraph

import (
	"context"
	"fmt"
	"strings"

	"github.com/machinebox/graphql"
	v0alpha "github.com/project-nova/backend/api/internal/entity/v0-alpha"
	"github.com/project-nova/backend/pkg/utils"
)

const (
	QUERY_INTERFACE = "$first: Int, $skip: Int, $orderBy: String, $orderDirection: String"
	QUERY_VALUE     = "first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection"
)

func NewTheGraphServiceAlphaImpl(client *graphql.Client) TheGraphServiceAlpha {
	return &theGraphServiceAlphaImpl{
		client: client,
	}
}

type theGraphServiceAlphaImpl struct {
	client *graphql.Client
}

func (s *theGraphServiceAlphaImpl) GetRelationship(relationshipId string) (*v0alpha.Relationship, error) {
	if !utils.IsValidNumberString(relationshipId) {
		return nil, fmt.Errorf("invalid relationshipId")
	}

	req := graphql.NewRequest(`
		query($relationshipId: String) {
			relationshipCreateds(where: {relationshipId: $relationshipId}) {
				transactionHash
				srcId
				srcAddress
				relationshipId
				relType
				id
				dstId
				dstAddress
				blockTimestamp
				blockNumber
		}
	}`)
	req.Var("relationshipId", relationshipId)

	ctx := context.Background()
	var relationshipsResponse v0alpha.RelationshipTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &relationshipsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the relationships from the graph. error: %v", err)
	}

	if len(relationshipsResponse.Relationships) == 0 {
		return nil, nil
	}

	return relationshipsResponse.Relationships[0].ToRelationship(), nil
}

func (s *theGraphServiceAlphaImpl) GetRelationshipType(relType string, ipOrgId string) (*v0alpha.RelationshipType, error) {
	if relType == "" || ipOrgId == "" {
		return nil, fmt.Errorf("relType and ipOrgId cannot be empty")
	}
	queryInterface := "$relType: String, $ipOrgId: String"
	queryValue := "where: {relType: $relType, ipOrgId: $ipOrgId}"

	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			relationshipTypeSets(%s) {
				blockNumber
				blockTimestamp
				dst
				dstRelatable
				dstSubtypesMask
				id
				ipOrgId
				relType
				src
				srcRelatable
				srcSubtypesMask
				transactionHash
			}
		}
	`, queryInterface, queryValue))
	req.Var("relType", relType)
	req.Var("ipOrgId", ipOrgId)

	ctx := context.Background()
	var relationshipTypesResponse v0alpha.RelationshipTypeTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &relationshipTypesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the relationship types from the graph. error: %v", err)
	}

	if len(relationshipTypesResponse.RelationshipTypes) == 0 {
		return nil, nil
	}

	return relationshipTypesResponse.RelationshipTypes[0].ToRelationshipType(), nil
}

func (s *theGraphServiceAlphaImpl) ListRelationshipTypes(ipOrgId *string, options *TheGraphQueryOptions) ([]*v0alpha.RelationshipType, error) {
	options = s.setQueryOptions(options)
	queryInterface := fmt.Sprintf("$ipOrgId: String, %s", QUERY_INTERFACE)
	queryValue := fmt.Sprintf("where: {ipOrgId: $ipOrgId}, %s", QUERY_VALUE)
	if ipOrgId == nil || *ipOrgId == "" {
		queryInterface = QUERY_INTERFACE
		queryValue = QUERY_VALUE
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			relationshipTypeSets(%s) {
				id
				relType
				ipOrgId
				src
				srcRelatable
				srcSubtypesMask
				dst
				dstRelatable
				dstSubtypesMask
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`, queryInterface, queryValue))
	if ipOrgId != nil {
		req.Var("ipOrgId", *ipOrgId)
	}

	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var relationshipTypesResponse v0alpha.RelationshipTypeTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &relationshipTypesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the relationship types from the graph. error: %v", err)
	}

	return relationshipTypesResponse.ToRelationshipTypes(), nil

}

func (s *theGraphServiceAlphaImpl) ListRelationships(contract string, tokenId string, options *TheGraphQueryOptions) ([]*v0alpha.Relationship, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(fmt.Sprintf(`
		query($contract: String, $tokenId: String, %s) {
			relationshipCreateds(where: {
				or: [
					{ srcAddress: $contract, srcId: $tokenId },
					{ dstAddress: $contract, dstId: $tokenId }
				]
			}, %s) {
				id
				relationshipId
				relType
				srcAddress
				srcId
				dstAddress
				dstId
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`, QUERY_INTERFACE, QUERY_VALUE))
	req.Var("contract", contract)
	req.Var("tokenId", tokenId)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)
	ctx := context.Background()
	var relationshipsResponse v0alpha.RelationshipTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &relationshipsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the relationships from the graph. error: %v", err)
	}

	return relationshipsResponse.ToRelationships(), nil
}

func (s *theGraphServiceAlphaImpl) ListHooks(moduleId *string, options *TheGraphQueryOptions) ([]*v0alpha.Hook, error) {
	options = s.setQueryOptions(options)
	queryInterface := fmt.Sprintf("$moduleId: String, %s", QUERY_INTERFACE)
	queryValue := fmt.Sprintf("where: {moduleId: $moduleId}, %s", QUERY_VALUE)
	if moduleId == nil || *moduleId == "" {
		queryInterface = QUERY_INTERFACE
		queryValue = QUERY_VALUE
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			hookRegistereds(%s) {
				blockNumber
				blockTimestamp
				id
				moduleId
				registryKey
				transactionHash
				type
			}
		}
	`, queryInterface, queryValue))

	if moduleId != nil {
		req.Var("moduleId", *moduleId)
	}
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)
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
			hookRegistereds(where: {id: $hookId}) {
				blockNumber
				blockTimestamp
				id
				moduleId
				registryKey
				transactionHash
				type
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
		return nil, nil
	}

	return hooksResponse.Hooks[0].ToHook(), nil
}

func (s *theGraphServiceAlphaImpl) ListModules(ipOrgId *string, options *TheGraphQueryOptions) ([]*v0alpha.Module, error) {
	options = s.setQueryOptions(options)
	queryInterface := fmt.Sprintf("$ipOrgId: String, %s", QUERY_INTERFACE)
	queryValue := fmt.Sprintf("where: {ipOrgId: $ipOrgId}, %s", QUERY_VALUE)
	if ipOrgId == nil || *ipOrgId == "" {
		queryInterface = QUERY_INTERFACE
		queryValue = QUERY_VALUE
	}
	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			moduleRegisterreds(%s) {
				id
				ipOrgId
				blockNumber
				blockTimestamp
				moduleId
				moduleKey
				transactionHash
			}
		}
	`, queryInterface, queryValue))
	if ipOrgId != nil {
		req.Var("ipOrgId", *ipOrgId)
	}

	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)
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
			moduleRegisterreds(where: {moduleId: $moduleId}) {
				blockNumber
				blockTimestamp
				id
				ipOrgId
				moduleId
				moduleKey
				transactionHash
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
		return nil, nil
	}

	return modulesResponse.Modules[0].ToModule(), nil
}

func (s *theGraphServiceAlphaImpl) ListIPOrgs(options *TheGraphQueryOptions) ([]*v0alpha.IPOrg, error) {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			iporgRegistereds(%s) {
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
	`, QUERY_INTERFACE, QUERY_VALUE))
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)
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
		return nil, nil
	}

	return ipOrgsResponse.IporgRegistereds[0].ToIPOrg(), nil
}

func (s *theGraphServiceAlphaImpl) ListIPAssets(iporgId *string, options *TheGraphQueryOptions) ([]*v0alpha.IPAsset, error) {
	options = s.setQueryOptions(options)

	queryInterface := fmt.Sprintf("$ipOrgId: String, %s", QUERY_INTERFACE)
	queryValue := fmt.Sprintf("where: {ipOrgId: $ipOrgId}, %s", QUERY_VALUE)
	if iporgId == nil || *iporgId == "" {
		queryInterface = QUERY_INTERFACE
		queryValue = QUERY_VALUE
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			ipassetRegistereds(%s) {
				blockNumber
				blockTimestamp
				contentHash
				id
				ipAssetId
				ipAssetTypeValue
				ipAssetTypeIndex
				ipOrgAssetId
				ipOrgId
				mediaUrl
				name
				owner
				transactionHash
			}
		}
	`, queryInterface, queryValue))
	if iporgId != nil {
		req.Var("ipOrgId", iporgId)
	}
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var ipAssetsResponse v0alpha.IpAssetTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &ipAssetsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}
	return ipAssetsResponse.ToIPAssets(), nil
}

func (s *theGraphServiceAlphaImpl) GetIPAsset(ipAssetId string) (*v0alpha.IPAsset, error) {
	req := graphql.NewRequest(`
		query($ipAssetId: String) {
			ipassetRegistereds(where: {ipAssetId: $ipAssetId}) {
				id
				ipAssetId
				ipOrgId
				ipOrgAssetId
				owner
				name
				ipAssetTypeValue
				ipAssetTypeIndex
				contentHash
				mediaUrl
				blockNumber
				blockTimestamp
				transactionHash
			}
		}
	`)
	req.Var("ipAssetId", ipAssetId)

	ctx := context.Background()
	var ipAssetsResponse v0alpha.IpAssetTheGraphAlphaResponse
	if err := s.client.Run(ctx, req, &ipAssetsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	if len(ipAssetsResponse.IpassetRegistereds) == 0 {
		return nil, nil
	}

	return ipAssetsResponse.IpassetRegistereds[0].ToIPAsset(), nil
}

func (s *theGraphServiceAlphaImpl) ListTransactions(ipOrgId *string, options *TheGraphQueryOptions) ([]*v0alpha.Transaction, error) {
	options = s.setQueryOptions(options)
	queryInterface := fmt.Sprintf("$ipOrgId: String, %s", QUERY_INTERFACE)
	queryValue := fmt.Sprintf("where: {ipOrgId: $ipOrgId}, %s", QUERY_VALUE)
	if ipOrgId == nil || *ipOrgId == "" {
		queryInterface = QUERY_INTERFACE
		queryValue = QUERY_VALUE
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			transactions(%s) {
				actionType
				blockNumber
				blockTimestamp
				id
				initiator
				ipOrgId
				resourceId
				resourceType
				transactionHash
			}
		}
	`, queryInterface, queryValue))
	if ipOrgId != nil {
		req.Var("ipOrgId", ipOrgId)
	}
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)
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
				actionType
				blockNumber
				blockTimestamp
				id
				initiator
				ipOrgId
				resourceId
				resourceType
				transactionHash
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
		return nil, nil
	}

	return transactionsResponse.Transactions[0].ToTransaction(), nil
}

func (s *theGraphServiceAlphaImpl) ListLicenses(ipOrgId *string, ipAssetId *string, options *TheGraphQueryOptions) ([]*v0alpha.License, error) {
	options = s.setQueryOptions(options)
	queryInterface := QUERY_INTERFACE
	queryValue := QUERY_VALUE
	whereClause := []string{}
	if ipOrgId != nil && *ipOrgId != "" {
		queryInterface += ", $ipOrgId: String"
		whereClause = append(whereClause, "ipOrgId: $ipOrgId")
	}
	if ipAssetId != nil && *ipAssetId != "" {
		queryInterface += ", $ipAssetId: String"
		whereClause = append(whereClause, "ipAssetId: $ipAssetId")
	}
	if len(whereClause) > 0 {
		queryValue = fmt.Sprintf("where: {%s}, %s", strings.Join(whereClause, ","), queryValue)
	}

	req := graphql.NewRequest(fmt.Sprintf(`
		query(%s) {
			licenseRegisterreds(%s) {
				transactionHash
				termsData
				termIds
				status
				revoker
				parentLicenseId
				licensor
				licenseeType
				licenseId
				isCommercial
				ipOrgId
				ipAssetId
				id
				blockTimestamp
				blockNumber
			}
		}
	`, queryInterface, queryValue))
	if ipOrgId != nil && *ipOrgId != "" {
		req.Var("ipOrgId", ipOrgId)
	}
	if ipAssetId != nil && *ipAssetId != "" {
		req.Var("ipAssetId", ipAssetId)
	}
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

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
			licenseRegisterreds(where: {licenseId: $licenseId}) {
				blockNumber
				blockTimestamp
				id
				ipAssetId
				ipOrgId
				licenseId
				isCommercial
				licenseeType
				licensor
				parentLicenseId
				revoker
				status
				termIds
				termsData
				transactionHash
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
		return nil, nil
	}

	return licenseResponse.LicenseRegistereds[0].ToLicense(), nil
}

func (s *theGraphServiceAlphaImpl) setQueryOptions(options *TheGraphQueryOptions) *TheGraphQueryOptions {
	if options == nil {
		options = &TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	if options.First == 0 {
		options.First = 100
	}

	options.OrderBy = "blockTimestamp"
	options.OrderDirection = "desc"

	return options
}
