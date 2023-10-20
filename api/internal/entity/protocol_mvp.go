package entity

import (
	"strconv"
	"time"
)

type FranchiseMVP struct {
	ID           string `json:"franchiseId,omitempty"`
	Name         string `json:"franchiseName,omitempty"`
	OwnerAddress string `json:"ownerAddress,omitempty"`
	TokenUri     string `json:"tokenUri,omitempty"`
	TxHash       string `json:"txHash,omitempty"`
}

type IpAssetType string

var IpAssetTypes = struct {
	Story     IpAssetType
	Character IpAssetType
	Art       IpAssetType
	Group     IpAssetType
	Location  IpAssetType
}{
	Story:     "Story",
	Character: "Character",
	Art:       "Art",
	Group:     "Group",
	Location:  "Location",
}

type IpAssetMVP struct {
	ID           string      `json:"ipAssetId,omitempty"`
	FranchiseId  string      `json:"franchiseId,omitempty"`
	Name         string      `json:"ipAssetName,omitempty"`
	Type         IpAssetType `json:"ipAssetType,omitempty"`
	OwnerAddress string      `json:"ownerAddress,omitempty"`
	TokenUri     string      `json:"tokenUri,omitempty"`
	Metadata     string      `json:"metadata,omitempty"`
	TxHash       string      `json:"txHash,omitempty"`
}

type LicenseMVP struct {
	ID              string `json:"licenseId,omitempty"`
	IpAssetId       string `json:"ipAssetId,omitempty"`
	FranchiseId     string `json:"franchiseId,omitempty"`
	ParentLicenseId string `json:"parentLicenseId,omitempty"`
	OwnerAddress    string `json:"ownerAddress,omitempty"`
	Uri             string `json:"uri,omitempty"`
	TxHash          string `json:"txHash,omitempty"`
}

type CollectionMVP struct {
	TotalCollected int64  `json:"totalCollected,omitempty"`
	IpAssetId      string `json:"ipAssetId,omitempty"`
	FranchiseId    string `json:"franchiseId,omitempty"`
}

type ResourceType string

var ResourceTypes = struct {
	Franchise    ResourceType
	IpAsset      ResourceType
	License      ResourceType
	Relationship ResourceType
	Collection   ResourceType
}{
	Franchise:    "Franchise",
	IpAsset:      "IPAsset",
	License:      "License",
	Relationship: "Relationship",
	Collection:   "Collection",
}

type TransactionMVP struct {
	ID             string       `json:"txId,omitempty"`
	TxHash         string       `json:"txHash,omitempty"`
	CreatedAt      string       `json:"createdAt,omitempty"`
	CreatorAddress string       `json:"creatorAddress,omitempty"`
	ResourceType   ResourceType `json:"type,omitempty"`
	ResourceId     string       `json:"resourceId,omitempty"`
	FranchiseId    string       `json:"franchiseId,omitempty"`
}

type GetFranchiseResponseMVP struct {
	Data *FranchiseMVP `json:"data"`
}

type GetFranchisesResponseMVP struct {
	Data []*FranchiseMVP `json:"data"`
}

type GetIpAssetResponseMVP struct {
	Data *IpAssetMVP `json:"data"`
}

type GetIpAssetsResponseMVP struct {
	Data []*IpAssetMVP `json:"data"`
}

type GetLicenseResponseMVP struct {
	Data *LicenseMVP `json:"data"`
}

type GetLicensesResponseMVP struct {
	Data []*LicenseMVP `json:"data"`
}

type GetCollectionsResponseMVP struct {
	Data []*CollectionMVP `json:"data"`
}

type GetTransactionResponseMVP struct {
	Data *TransactionMVP `json:"data"`
}

type GetTransactionsResponseMVP struct {
	Data []*TransactionMVP `json:"data"`
}

type FranchisesTheGraphMVPResponse struct {
	Franchises []*FranchiseTheGraphMVP `json:"franchises"`
}

type FranchiseTheGraphMVP struct {
	ID              string `json:"id"`
	FranchiseId     string `json:"franchiseId"`
	Owner           string `json:"owner"`
	IpAssetRegistry string `json:"ipAssetRegistry"`
	Name            string `json:"name"`
	TokenURI        string `json:"tokenURI"`
	TxHash          string `json:"transactionHash"`
}

func (f *FranchiseTheGraphMVP) ToFranchise() *FranchiseMVP {
	franchise := &FranchiseMVP{
		ID:           f.FranchiseId,
		OwnerAddress: f.Owner,
		Name:         f.Name,
		TokenUri:     f.TokenURI,
		TxHash:       f.TxHash,
	}
	return franchise
}

type LicensesTheGraphMVPResponse struct {
	Licenses []*LicenseTheGraphMVP `json:"licenses"`
}

type LicenseTheGraphMVP struct {
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
	ID              string `json:"id"`
	LicenseId       string `json:"licenseId"`
	FranchiseId     string `json:"franchiseId"`
	IpAssetId       string `json:"ipAssetId"`
	OwnerAddress    string `json:"licenseOwner"`
	ParentLicenseId string `json:"parentLicenseId"`
	Uri             string `json:"uri"`
	TxHash          string `json:"transactionHash"`
}

func (l *LicenseTheGraphMVP) ToLicense() *LicenseMVP {
	license := &LicenseMVP{
		ID:              l.LicenseId,
		FranchiseId:     l.FranchiseId,
		IpAssetId:       l.IpAssetId,
		OwnerAddress:    l.OwnerAddress,
		ParentLicenseId: l.ParentLicenseId,
		Uri:             l.Uri,
		TxHash:          l.TxHash,
	}
	return license
}

type IpAssetsTheGraphMVPResposne struct {
	IpAssets []*IpAssetTheGraphMVP `json:"ipassets"`
}

type IpAssetTheGraphMVP struct {
	/*
			id: Bytes!
		    franchiseId: BigInt! # uint256
		    owner: Bytes! # address
		    ipAssetRegistry: Bytes! # address
		    name: String!
		    description: String!
		    ipAssetId: BigInt! # uint256
		    ipAssetType: IpAssetType!
		    mediaUrl: String!
		    blockNumber: BigInt!
		    blockTimestamp: BigInt!
		    transactionHash: Bytes!
	*/
	ID          string `json:"id"`
	FranchiseId string `json:"franchiseId"`
	IpAssetId   string `json:"ipAssetId"`
	IpAssetType string `json:"ipAssetType"`
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	MediaUrl    string `json:"mediaUrl"`
	TxHash      string `json:"transactionHash"`
}

func (f *IpAssetTheGraphMVP) ToIpAsset() *IpAssetMVP {
	ipAsset := &IpAssetMVP{
		ID:           f.IpAssetId,
		Type:         IpAssetType(f.IpAssetType),
		Name:         f.Name,
		FranchiseId:  f.FranchiseId,
		OwnerAddress: f.Owner,
		TokenUri:     f.MediaUrl,
		TxHash:       f.TxHash,
	}
	return ipAsset
}

type CollectionsTheGraphMVPResposne struct {
	Collections []*CollectionTheGraphMVP `json:"collections"`
}

type CollectionTheGraphMVP struct {
	/*
	  id: Bytes! # address
	  franchiseId: BigInt! # uint256
	  ipAssetId: BigInt! # uint256
	  totalCollected: BigInt! # uint256
	*/
	ID             string `json:"id"`
	FranchiseId    string `json:"franchiseId"`
	IpAssetId      string `json:"ipAssetId"`
	TotalCollected string `json:"totalCollected"`
}

func (f *CollectionTheGraphMVP) ToCollection() (*CollectionMVP, error) {
	totalCollected, err := strconv.ParseInt(f.TotalCollected, 10, 64)
	if err != nil {
		return nil, err
	}
	collection := &CollectionMVP{
		IpAssetId:      f.IpAssetId,
		FranchiseId:    f.FranchiseId,
		TotalCollected: totalCollected,
	}
	return collection, nil
}

type TransactionsTheGraphMVPResposne struct {
	Transactions []*TransactionTheGraphMVP `json:"transactions"`
}

type TransactionTheGraphMVP struct {
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
	ID             string `json:"id"`
	Owner          string `json:"owner"`
	FranchiseId    string `json:"franchiseId"`
	ResourceId     string `json:"resourceId"`
	ResourceType   string `json:"resourceType"`
	TxHash         string `json:"transactionHash"`
	BlockTimestamp string `json:"blockTimestamp"`
}

func (f *TransactionTheGraphMVP) ToTransaction() (*TransactionMVP, error) {
	transaction := &TransactionMVP{
		ID:             f.ID,
		TxHash:         f.TxHash,
		FranchiseId:    f.FranchiseId,
		ResourceId:     f.ResourceId,
		ResourceType:   ResourceType(f.ResourceType),
		CreatorAddress: f.Owner,
		CreatedAt:      f.BlockTimestamp,
	}
	blockTimeInt64, err := strconv.ParseInt(f.BlockTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(blockTimeInt64, 0)
	transaction.CreatedAt = tm.UTC().Format("2006-01-02T15:04:05.999Z")
	return transaction, nil
}
