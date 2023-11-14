package v0alpha

type IPAsset struct {
	ID          string      `json:"id,omitempty"`
	Type        IPAssetType `json:"type,omitempty"`
	IPOrgId     string      `json:"ipOrgId,omitempty"`
	Owner       string      `json:"owner,omitempty"`
	MetadataUrl string      `json:"metadataUrl,omitempty"`
	Hash        []byte      `json:"hash,omitempty"`
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
