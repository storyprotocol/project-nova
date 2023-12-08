package v0alpha

import (
	"github.com/project-nova/backend/pkg/utils"
)

type Param struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type IpOrgLicenseParam struct {
	IPOrgId        string  `json:"ipOrgId"`
	FrameworkId    string  `json:"frameworkId"`
	URL            string  `json:"url"`
	LicensorConfig int     `json:"licensorConfig"`
	Params         []Param `json:"params"`
	CreatedAt      string  `json:"createdAt"`
	TxHash         string  `json:"txHash"`
}

type ListLicenseParamsRequest struct {
	IpOrgId string        `json:"ipOrgId"`
	Options *QueryOptions `json:"options"`
}

func (l *ListLicenseParamsRequest) Validate() bool {
	return utils.IsValidAddress(l.IpOrgId)
}

type ListLicenseParamsResponse struct {
	LicenseParams []*IpOrgLicenseParam `json:"licenseParams"`
}

type IpOrgLicenseParamTheGraph struct {
	ID              string   `json:"id"`
	FrameworkId     string   `json:"frameworkId"`
	URL             string   `json:"url"`
	LicensorConfig  int      `json:"licensorConfig"`
	ParamTags       []string `json:"paramTags"`
	ParamValues     []string `json:"paramValues"`
	BlockNumber     string   `json:"blockNumber"`
	BlockTimestamp  string   `json:"blockTimestamp"`
	TransactionHash string   `json:"transactionHash"`
}

type IpOrgLicenseParamTheGraphResponse struct {
	IpOrgLicenseParams []*IpOrgLicenseParamTheGraph `json:"iporgLicenseParams"`
}

func (l *IpOrgLicenseParamTheGraphResponse) ToLicenseParams() []*IpOrgLicenseParam {
	var licenseParams []*IpOrgLicenseParam
	for _, licenseParam := range l.IpOrgLicenseParams {
		convertedLicenseParam := licenseParam.ToLicenseParams()
		licenseParams = append(licenseParams, convertedLicenseParam)
	}
	return licenseParams
}

func (l *IpOrgLicenseParamTheGraph) ToLicenseParams() *IpOrgLicenseParam {
	licensorConfig := int(l.LicensorConfig)

	licenseParam := &IpOrgLicenseParam{
		IPOrgId:        l.ID,
		FrameworkId:    l.FrameworkId,
		URL:            l.URL,
		LicensorConfig: licensorConfig,
		CreatedAt:      l.BlockTimestamp,
		TxHash:         l.TransactionHash,
	}

	for i, tag := range l.ParamTags {
		param := Param{
			Tag:   tag,
			Value: l.ParamValues[i],
		}
		licenseParam.Params = append(licenseParam.Params, param)
	}

	return licenseParam
}
