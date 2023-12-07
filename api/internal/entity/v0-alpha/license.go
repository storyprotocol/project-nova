package v0alpha

import (
	"github.com/project-nova/backend/pkg/utils"
)

type License struct {
	ID                      string `json:"id,omitempty"`
	IsReciprocal            bool   `json:"isReciprocal"`
	DerivativeNeedsApproval bool   `json:"derivativeNeedsApproval"`
	Status                  int    `json:"status"`
	Licensor                string `json:"licensor"`
	Revoker                 string `json:"revoker"`
	IpOrgId                 string `json:"ipOrgId"`
	IpAssetId               string `json:"ipAssetId"`
	ParentLicenseId         string `json:"parentLicenseId"`
	CreatedAt               string `json:"createdAt"`
	TxHash                  string `json:"txHash"`
}

type GetLicenseResponse struct {
	License *License `json:"license"`
}

type ListLicensesRequest struct {
	IpOrgId   *string       `json:"ipOrgId"`
	IpAssetId *string       `json:"ipAssetId"`
	Options   *QueryOptions `json:"options"`
}

func (l *ListLicensesRequest) Validate() bool {
	if l.IpOrgId != nil && !utils.IsValidAddress(l.IpOrgId) {
		return false
	}

	if l.IpAssetId != nil && !utils.IsValidAddress(l.IpAssetId) {
		return false
	}

	return true
}

type ListLicensesResponse struct {
	Licenses []*License `json:"licenses"`
}

type LicenseRegistryTheGraphAlpha struct {
	ID                      string `json:"id"`
	LicenseId               string `json:"licenseId"`
	Status                  int    `json:"status"`
	IsReciprocal            bool   `json:"isReciprocal"`
	DerivativeNeedsApproval bool   `json:"derivativeNeedsApproval"`
	Revoker                 string `json:"revoker"`
	Licensor                string `json:"licensor"`
	IpOrgId                 string `json:"ipOrgId"`
	IpAssetId               string `json:"ipAssetId"`
	ParentLicenseId         string `json:"parentLicenseId"`
	BlockNumber             string `json:"blockNumber"`
	BlockTimestamp          string `json:"blockTimestamp"`
	TxHash                  string `json:"transactionHash"`
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
		ID:                      l.LicenseId,
		IsReciprocal:            l.IsReciprocal,
		DerivativeNeedsApproval: l.DerivativeNeedsApproval,
		Status:                  l.Status,
		Licensor:                l.Licensor,
		Revoker:                 l.Revoker,
		IpOrgId:                 l.IpOrgId,
		IpAssetId:               l.IpAssetId,
		ParentLicenseId:         l.ParentLicenseId,
		CreatedAt:               utils.TimestampInSecondsToISO8601(l.BlockTimestamp),
		TxHash:                  l.TxHash,
	}
}
