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

type ListIpAssetsResponse struct {
	IPAssets []*IPAsset `json:"ipassets"`
}
