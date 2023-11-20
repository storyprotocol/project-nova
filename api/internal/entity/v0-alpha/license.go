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

type ListLicensesRequest struct {
	IpOrgId   *string       `json:"ipOrgId"`
	IpAssetId *string       `json:"ipAssetId"`
	Options   *QueryOptions `json:"options"`
}

type ListLicensesResponse struct {
	Licenses []*License `json:"licenses"`
}

type LicenseRegistryTheGraphAlpha struct {
	ID              string   `json:"id"`
	LicenseId       string   `json:"licenseId"`
	IsCommercial    bool     `json:"isCommercial"`
	Status          int      `json:"status"`
	Licensor        string   `json:"licensor"`
	Revoker         string   `json:"revoker"`
	IpOrgId         string   `json:"ipOrgId"`
	LicenseeType    int      `json:"licenseeType"`
	IpAssetId       string   `json:"ipAssetId"`
	ParentLicenseId string   `json:"parentLicenseId"`
	TermIds         []string `json:"termIds"`
	TermsData       []string `json:"termsData"`
	BlockNumber     string   `json:"blockNumber"`
	BlockTimestamp  string   `json:"blockTimestamp"`
	TxHash          string   `json:"transactionHash"`
}

type LicenseTheGraphAlphaResponse struct {
	LicenseRegistereds []*LicenseRegistryTheGraphAlpha `json:"licenseRegistereds"`
}

func (l *LicenseTheGraphAlphaResponse) ToLicenses() []*License {
	licenses := []*License{}
	for _, license := range l.LicenseRegistereds {
		licenses = append(licenses, license.ToLicense())
	}

	return licenses
}

func (l *LicenseRegistryTheGraphAlpha) ToLicense() *License {
	return &License{
		ID:          l.ID,
		IPAssetId:   l.IpAssetId,
		IPOrgId:     l.IpOrgId,
		Owner:       l.Licensor,
		MetadataUri: l.TermsData[0],
		CreatedAt:   l.BlockTimestamp,
		TxHash:      l.TxHash,
	}
}
