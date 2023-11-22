package v0alpha

type IPOrg struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Symbol       string   `json:"symbol,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	BaseUri      string   `json:"baseUri,omitempty"`
	ContractUri  string   `json:"contractUri,omitempty"`
	IPAssetTypes []string `json:"ipAssetTypes"`
	CreatedAt    string   `json:"createdAt,omitempty"`
	TxHash       string   `json:"txHash,omitempty"`
}

type GetIpOrgResponse struct {
	IPOrg *IPOrg `json:"ipOrg"`
}

type ListIpOrgsRequest struct {
	Options *QueryOptions `json:"options"`
}

type ListIpOrgsResponse struct {
	IPOrgs []*IPOrg `json:"ipOrgs"`
}

type IPOrgTheGraphAlpha struct {
	ID             string   `json:"id"`
	Owner          string   `json:"owner"`
	IPOrgId        string   `json:"ipOrgId"`
	Name           string   `json:"name"`
	Symbol         string   `json:"symbol"`
	IPAssetTypes   []string `json:"ipAssetTypes"`
	BaseURI        string   `json:"baseURI"`
	ContractURI    string   `json:"contractURI"`
	BlockNumber    string   `json:"blockNumber"`
	BlockTimestamp string   `json:"blockTimestamp"`
	TxHash         string   `json:"transactionHash"`
}

type IpOrgTheGraphAlphaResponse struct {
	IporgRegistereds []*IPOrgTheGraphAlpha `json:"iporgRegistereds"`
}

func (i *IpOrgTheGraphAlphaResponse) ToIPOrgs() []*IPOrg {
	iporgs := []*IPOrg{}
	for _, iporg := range i.IporgRegistereds {
		iporgs = append(iporgs, iporg.ToIPOrg())
	}

	return iporgs
}

func (i *IPOrgTheGraphAlpha) ToIPOrg() *IPOrg {
	return &IPOrg{
		ID:           i.ID,
		Name:         i.Name,
		Symbol:       i.Symbol,
		Owner:        i.Owner,
		BaseUri:      i.BaseURI,
		ContractUri:  i.ContractURI,
		IPAssetTypes: i.IPAssetTypes,
		CreatedAt:    i.BlockTimestamp,
		TxHash:       i.TxHash,
	}
}
