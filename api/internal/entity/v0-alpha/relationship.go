package v0alpha

type Relationship struct {
	ID           string `json:"id,omitempty"`
	Type         string `json:"type,omitempty"`
	TypeId       string `json:"typeId,omitempty"`
	SrcContract  string `json:"srcContract,omitempty"`
	SrcTokenId   string `json:"srcTokenId,omitempty"`
	SrcName      string `json:"srcName,omitempty"`
	DstContract  string `json:"dstContract,omitempty"`
	DstTokenId   string `json:"dstTokenId,omitempty"`
	DstName      string `json:"dstName,omitempty"`
	TTL          *int64 `json:"ttl,omitempty"`
	RegisteredAt string `json:"registeredAt,omitempty"`
	TxHash       string `json:"txHash,omitempty"`
}
