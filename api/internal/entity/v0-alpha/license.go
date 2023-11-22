package v0alpha

import (
	"github.com/project-nova/backend/pkg/utils"
)

type License struct {
	ID              string   `json:"id,omitempty"`
	IsCommercial    bool     `json:"isCommercial,omitempty"`
	Status          int      `json:"status,omitempty"`
	Licensor        string   `json:"licensor,omitempty"`
	Revoker         string   `json:"revoker,omitempty"`
	IPOrgId         string   `json:"ipOrgId,omitempty"`
	LicenseeType    int      `json:"licenseeType,omitempty"`
	IPAssetId       string   `json:"ipAssetId,omitempty"`
	ParentLicenseId string   `json:"parentLicenseId,omitempty"`
	TermIds         []string `json:"termIds,omitempty"`
	TermsData       []string `json:"termsData,omitempty"`
	CreatedAt       string   `json:"createdAt,omitempty"`
	TxHash          string   `json:"txHash,omitempty"`
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
	Licenses []*License `json:"licenseRegisterreds"`
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
	LicenseRegistereds []*LicenseRegistryTheGraphAlpha `json:"licenseRegisterreds"`
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
		ID:              l.LicenseId,
		IsCommercial:    l.IsCommercial,
		Status:          l.Status,
		Licensor:        l.Licensor,
		Revoker:         l.Revoker,
		IPOrgId:         l.IpOrgId,
		LicenseeType:    l.LicenseeType,
		IPAssetId:       l.IpAssetId,
		ParentLicenseId: l.ParentLicenseId,
		TermIds:         l.TermIds,
		TermsData:       l.TermsData,
		CreatedAt:       utils.TimestampInSecondsToISO8601(l.BlockTimestamp),
		TxHash:          l.TxHash,
	}
}
