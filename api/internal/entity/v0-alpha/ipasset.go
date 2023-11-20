package v0alpha

type IPAsset struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Type        IPAssetType `json:"type,omitempty"`
	IPOrgId     string      `json:"ipOrgId,omitempty"`
	Owner       string      `json:"owner,omitempty"`
	MetadataUrl string      `json:"metadataUrl,omitempty"`
	ContentHash []byte      `json:"contentHash,omitempty"`
	Data        []byte      `json:"data,omitempty"`
	CreatedAt   string      `json:"createdAt,omitempty"`
	TxHash      string      `json:"txHash,omitempty"`
}

type IPAssetType string

var IpAssetTypes = struct {
	Story     IPAssetType
	Character IPAssetType
	Art       IPAssetType
	Item      IPAssetType
	Location  IPAssetType
}{
	Story:     "Story",
	Character: "Character",
	Art:       "Art",
	Item:      "Item",
	Location:  "Location",
}

type GetIpAssetResponse struct {
	IPAsset *IPAsset `json:"ipasset"`
}

type ListIpAssetsRequest struct {
	Options *QueryOptions `json:"options"`
}

type ListIpAssetsResponse struct {
	IPAssets []*IPAsset `json:"ipassets"`
}

type IPAssetTheGraphAlpha struct {
	ID             string `json:"id"`
	IpAssetId      string `json:"ipAssetId"`
	IpOrgId        string `json:"ipOrgId"`
	IpOrgAssetId   string `json:"ipOrgAssetId"`
	Owner          string `json:"owner"`
	Name           string `json:"name"`
	IpAssetType    string `json:"ipAssetType"`
	ContentHash    string `json:"contentHash"`
	MediaUrl       string `json:"mediaUrl"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
	TxHash         string `json:"transactionHash"`
}

type IpAssetTheGraphAlphaResponse struct {
	IpassetRegistereds []*IPAssetTheGraphAlpha `json:"ipassetRegistereds"`
}

func (i *IpAssetTheGraphAlphaResponse) ToIPAssets() []*IPAsset {
	ipassets := []*IPAsset{}
	for _, ipasset := range i.IpassetRegistereds {
		ipassets = append(ipassets, ipasset.ToIPAsset())
	}

	return ipassets
}

func (i *IPAssetTheGraphAlpha) ToIPAsset() *IPAsset {
	return &IPAsset{
		ID:          i.ID,
		Name:        i.Name,
		Type:        IPAssetType(i.IpAssetType),
		IPOrgId:     i.IpOrgId,
		Owner:       i.Owner,
		MetadataUrl: i.MediaUrl,
		ContentHash: []byte(i.ContentHash),
		CreatedAt:   i.BlockTimestamp,
		TxHash:      i.TxHash,
	}
}
