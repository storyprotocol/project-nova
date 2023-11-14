package v0alpha

type License struct {
	ID          string `json:"licenseId,omitempty"`
	IPAssetId   string `json:"ipAssetId,omitempty"`
	IPOrgId     string `json:"ipOrgId,omitempty"`
	Owner       string `json:"owner,omitempty"`
	MetadataUri string `json:"meataUri,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	TxHash      string `json:"txHash,omitempty"`
}

type GetLicenseResponse struct {
	License *License `json:"license"`
}

type ListLicensesResponse struct {
	Licenses []*License `json:"licenses"`
}
