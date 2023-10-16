package thegraph

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/project-nova/backend/api/internal/entity"
)

func NewTheGraphServiceMvpImpl(client *graphql.Client) TheGraphServiceMvp {
	return &theGraphServiceMvpImpl{
		client: client,
	}
}

type theGraphServiceMvpImpl struct {
	client *graphql.Client
}

func (c *theGraphServiceMvpImpl) GetFranchises() ([]*entity.FranchiseMVP, error) {
	req := graphql.NewRequest(`
    {
		franchises {
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
	var franchisesResponse entity.FranchisesTheGraphMVPResponse
	if err := c.client.Run(ctx, req, &franchisesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}

	franchises := []*entity.FranchiseMVP{}
	for _, franchise := range franchisesResponse.Franchises {
		franchises = append(franchises, franchise.ToFranchise())
	}

	return franchises, nil
}

func (c *theGraphServiceMvpImpl) GetFranchise(franchiseId string) (*entity.FranchiseMVP, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt) {
		franchises(where: {franchiseId: $franchiseId}) {
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
	var franchiseResponse entity.FranchisesTheGraphMVPResponse
	if err := c.client.Run(ctx, req, &franchiseResponse); err != nil {
		return nil, fmt.Errorf("failed to get the franchises from the graph. error: %v", err)
	}
	if len(franchiseResponse.Franchises) == 0 {
		return nil, fmt.Errorf("failed to find the franchise")
	}

	return franchiseResponse.Franchises[0].ToFranchise(), nil
}

func (c *theGraphServiceMvpImpl) GetIpAssets(franchiseId string) ([]*entity.IpAssetMVP, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt) {
		ipassets(where: {franchiseId: $franchiseId}) {
			id
			franchiseId
			ipAssetId
			ipAssetType
			owner
			name
			mediaUrl
			transactionHash
	  	}
	}`)

	req.Var("franchiseId", franchiseId)

	ctx := context.Background()
	var ipAssetsResponse entity.IpAssetsTheGraphMVPResposne
	if err := c.client.Run(ctx, req, &ipAssetsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the ip assets from the graph. error: %v", err)
	}

	ipAssets := []*entity.IpAssetMVP{}
	for _, ipAsset := range ipAssetsResponse.IpAssets {
		ipAssets = append(ipAssets, ipAsset.ToIpAsset())
	}

	return ipAssets, nil
}

func (c *theGraphServiceMvpImpl) GetIpAsset(franchiseId string, ipAssetId string) (*entity.IpAssetMVP, error) {
	/*
		ID          string `json:"id"`
		FranchiseId string `json:"franchiseId"`
		IpAssetId   string `json:"ipAssetId"`
		IpAssetType string `json:"ipAssetType"`
		Owner       string `json:"owner"`
		Name        string `json:"name"`
		MediaUrl    string `json:"mediaUrl"`
		TxHash      string `json:"transactionHash"`
	*/
	req := graphql.NewRequest(`
    query($franchiseId: BigInt, $ipAssetId: BigInt) {
		ipassets(where: { and: [{ franchiseId: $franchiseId }, { ipAssetId: $ipAssetId }]}) {
			id
			franchiseId
			ipAssetId
			ipAssetType
			owner
			name
			mediaUrl
			transactionHash
	  	}
	}`)

	req.Var("franchiseId", franchiseId)
	req.Var("ipAssetId", ipAssetId)

	ctx := context.Background()
	var ipAssetsResponse entity.IpAssetsTheGraphMVPResposne
	if err := c.client.Run(ctx, req, &ipAssetsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the ip assets from the graph. error: %v", err)
	}
	if len(ipAssetsResponse.IpAssets) == 0 {
		return nil, fmt.Errorf("failed to find the ip asset")
	}

	return ipAssetsResponse.IpAssets[0].ToIpAsset(), nil
}

func (c *theGraphServiceMvpImpl) GetLicenses(franchiseId string, ipAssetId string) ([]*entity.LicenseMVP, error) {
	req := graphql.NewRequest(`
    query($franchiseId: BigInt, $ipAssetId: BigInt) {
		licenses(where: { and: [{ franchiseId: $franchiseId }, { ipAssetId: $ipAssetId }]}) {
			id
			licenseId
			franchiseId
			ipAssetId
			licenseOwner
			parentLicenseId
			uri
			transactionHash
	  	}
	}`)

	req.Var("franchiseId", franchiseId)
	req.Var("ipAssetId", ipAssetId)

	ctx := context.Background()
	var licensesResponse entity.LicensesTheGraphMVPResponse
	if err := c.client.Run(ctx, req, &licensesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the licenses from the graph. error: %v", err)
	}

	licenses := []*entity.LicenseMVP{}
	for _, license := range licensesResponse.Licenses {
		licenseModel := license.ToLicense()
		licenses = append(licenses, licenseModel)
	}

	return licenses, nil
}

func (c *theGraphServiceMvpImpl) GetLicense(licenseId string) (*entity.LicenseMVP, error) {
	/*
	  id: Bytes!
	  licenseId: BigInt! # uint256
	  franchiseId: BigInt! # uint256
	  ipAssetId: BigInt! # uint256
	  parentLicenseId: BigInt! # uint256
	  licenseOwner: Bytes! # address
	  uri: String! # string
	  revoker: Bytes! # address
	  blockNumber: BigInt!
	  blockTimestamp: BigInt!
	  transactionHash: Bytes!
	*/
	req := graphql.NewRequest(`
    query($licenseId: BigInt) {
		licenses(where: {licenseId: $licenseId}) {
			id
			licenseId
			franchiseId
			ipAssetId
			licenseOwner
			parentLicenseId
			uri
			transactionHash
	  	}
	}`)

	req.Var("licenseId", licenseId)

	ctx := context.Background()
	var licensesResponse entity.LicensesTheGraphMVPResponse
	if err := c.client.Run(ctx, req, &licensesResponse); err != nil {
		return nil, fmt.Errorf("failed to get the licenses from the graph. error: %v", err)
	}
	if len(licensesResponse.Licenses) == 0 {
		return nil, fmt.Errorf("failed to find the license")
	}

	return licensesResponse.Licenses[0].ToLicense(), nil
}

func (c *theGraphServiceMvpImpl) GetCollections(franchiseId string) ([]*entity.CollectionMVP, error) {
	/*
	  id: Bytes! # address
	  franchiseId: BigInt! # uint256
	  ipAssetId: BigInt! # uint256
	  totalCollected: BigInt! # uint256
	*/
	req := graphql.NewRequest(`
    query($franchiseId: BigInt) {
		collections(where: {franchiseId: $franchiseId}) {
			id
			franchiseId
			ipAssetId
			totalCollected	
	  	}
	}`)

	req.Var("franchiseId", franchiseId)

	ctx := context.Background()
	var collectionsResponse entity.CollectionsTheGraphMVPResposne
	if err := c.client.Run(ctx, req, &collectionsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the collections from the graph. error: %v", err)
	}

	collections := []*entity.CollectionMVP{}
	for _, collection := range collectionsResponse.Collections {
		collectionModel, err := collection.ToCollection()
		if err != nil {
			return nil, fmt.Errorf("failed to convert the graph collection to collection. error: %v", err)
		}
		collections = append(collections, collectionModel)
	}

	return collections, nil
}

func (c *theGraphServiceMvpImpl) GetTransactions() ([]*entity.TransactionMVP, error) {
	/*
		id: Bytes!
		owner: Bytes! # address
		franchiseId: BigInt! # uint256
		resourceType: ResourceType!
		resourceId: BigInt! # uint256
		blockNumber: BigInt!
		blockTimestamp: BigInt!
		transactionHash: Bytes!
	*/
	req := graphql.NewRequest(`
	{
		transactions(orderBy: blockTimestamp, orderDirection: desc) {
			id
			owner
			franchiseId
			resourceId
			resourceType
			transactionHash
			blockTimestamp
	  	}
	}`)

	ctx := context.Background()
	var transactionsResponse entity.TransactionsTheGraphMVPResposne
	if err := c.client.Run(ctx, req, &transactionsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the transactions from the graph. error: %v", err)
	}

	transactions := []*entity.TransactionMVP{}
	for _, transaction := range transactionsResponse.Transactions {
		transactionModel, err := transaction.ToTransaction()
		if err != nil {
			return nil, fmt.Errorf("failed to convert the graph transaction to transaction. error: %v", err)
		}
		transactions = append(transactions, transactionModel)
	}

	return transactions, nil
}

func (c *theGraphServiceMvpImpl) GetTransaction(transactionId string) (*entity.TransactionMVP, error) {
	/*
		id: Bytes!
		owner: Bytes! # address
		franchiseId: BigInt! # uint256
		resourceType: ResourceType!
		resourceId: BigInt! # uint256
		blockNumber: BigInt!
		blockTimestamp: BigInt!
		transactionHash: Bytes!
	*/
	req := graphql.NewRequest(`
    query($transactionId: BigInt) {
		transactions(where: {id: $transactionId}) {
			id
			owner
			franchiseId
			resourceId
			resourceType
			transactionHash
			blockTimestamp
	  	}
	}`)

	req.Var("transactionId", transactionId)

	ctx := context.Background()
	var transactionsResponse entity.TransactionsTheGraphMVPResposne
	if err := c.client.Run(ctx, req, &transactionsResponse); err != nil {
		return nil, fmt.Errorf("failed to get the transactions from the graph. error: %v", err)
	}
	if len(transactionsResponse.Transactions) == 0 {
		return nil, fmt.Errorf("failed to find the transaction")
	}

	return transactionsResponse.Transactions[0].ToTransaction()
}
