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

type ListIpOrgsResponse struct {
	IPOrgs []*IPOrg `json:"iporgs"`
}
