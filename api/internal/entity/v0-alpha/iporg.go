package v0alpha

type IPOrg struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
	Owner       string `json:"owner,omitempty"`
	MetadataUrl string `json:"metadataUrl,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	TxHash      string `json:"txHash,omitempty"`
}

type GetIpOrgResponse struct {
	IPOrg *IPOrg `json:"iporg"`
}

type ListIpOrgsRequest struct {
	Options *QueryOptions `json:"options"`
}

type ListIpOrgsResponse struct {
	IPOrgs []*IPOrg `json:"iporgs"`
}

type IPOrgTheGraphAlpha struct {
	ID             string   `json:"id"`
	Owner          string   `json:"owner"`
	IPOrgId        string   `json:"ipOrgId"`
	Name           string   `json:"name"`
	Symbol         string   `json:"symbol"`
	IpAssetTypes   []string `json:"ipAssetTypes"`
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
		ID:          i.ID,
		Name:        i.Name,
		Symbol:      i.Symbol,
		Owner:       i.Owner,
		MetadataUrl: i.BaseURI,
		CreatedAt:   i.BlockTimestamp,
		TxHash:      i.TxHash,
	}
}
