package v0alpha

import "github.com/project-nova/backend/pkg/utils"

type IPAsset struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        int64  `json:"type,omitempty"`
	IPOrgId     string `json:"ipOrgId,omitempty"`
	Owner       string `json:"owner,omitempty"`
	MediaUrl    string `json:"mediaUrl,omitempty"`
	ContentHash []byte `json:"contentHash,omitempty"`
	Data        []byte `json:"data,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	TxHash      string `json:"txHash,omitempty"`
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
	IPAsset *IPAsset `json:"ipAssets"`
}

type ListIpAssetsRequest struct {
	IpOrgId string        `json:"ipOrgId"`
	Options *QueryOptions `json:"options"`
}

type ListIpAssetsResponse struct {
	IPAssets []*IPAsset `json:"ipAssets"`
}

type IPAssetTheGraphAlpha struct {
	ID             string `json:"id"`
	IPAssetId      string `json:"ipAssetId"`
	IPOrgId        string `json:"ipOrgId"`
	IPOrgAssetId   string `json:"ipOrgAssetId"`
	Owner          string `json:"owner"`
	Name           string `json:"name"`
	IPAssetType    int64  `json:"ipAssetType"`
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
		ID:          i.IPAssetId,
		Name:        i.Name,
		Type:        i.IPAssetType,
		IPOrgId:     i.IPOrgId,
		Owner:       i.Owner,
		MediaUrl:    i.MediaUrl,
		ContentHash: []byte(i.ContentHash),
		CreatedAt:   utils.TimestampInSecondsToISO8601(i.BlockTimestamp),
		TxHash:      i.TxHash,
	}
}
