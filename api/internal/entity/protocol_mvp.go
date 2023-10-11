package entity

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
}{
	Story:     "STORY",
	Character: "CHARACTER",
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
	Franchise:    "FRANCHISE",
	IpAsset:      "IP_ASSET",
	License:      "LICENSE",
	Relationship: "RELATIONSHIP",
	Collection:   "COLLECTION",
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
	Data FranchiseMVP
}

type GetFranchisesResponseMVP struct {
	Data []FranchiseMVP
}

type GetIpAssetResponseMVP struct {
	Data IpAssetMVP
}

type GetIpAssetsResponseMVP struct {
	Data []IpAssetMVP
}

type GetLicenseResponseMVP struct {
	Data LicenseMVP
}

type GetLicensesResponseMVP struct {
	Data []LicenseMVP
}

type GetCollectionsResponseMVP struct {
	Data []CollectionMVP
}

type GetTransactionResponseMVP struct {
	Data TransactionMVP
}

type GetTransactionsResponseMVP struct {
	Data []TransactionMVP
}
