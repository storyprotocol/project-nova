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

type GetRelationshipResponse struct {
	Relationship *Relationship `json:"relationship"`
}

type ListRelationshipRequest struct {
	Contract string        `json:"contract"`
	TokenId  string        `json:"tokenId"`
	Options  *QueryOptions `json:"options"`
}

type ListRelationshipsResponse struct {
	Relationships []*Relationship `json:"relationships"`
}

type RelationshipTheGraphAlpha struct {
	ID               string `json:"id"`
	RelationshipId   string `json:"relationshipId"`
	RelationshipType string `json:"relType"`
	SrcAddress       string `json:"srcAddress"`
	SrcId            string `json:"srcId"`
	DstAddress       string `json:"dstAddress"`
	DstId            string `json:"dstId"`
	BlockNumber      string `json:"blockNumber"`
	BlockTimestamp   string `json:"blockTimestamp"`
	TxHash           string `json:"transactionHash"`
}

type RelationshipTypeTheGraphAlpha struct {
	ID               string `json:"id"`
	RelationshipId   string `json:"relationshipId"`
	RelationshipType string `json:"relType"`
	SrcAddress       string `json:"srcAddress"`
	SrcId            string `json:"srcId"`
	DstAddress       string `json:"dstAddress"`
	DstId            string `json:"dstId"`
	BlockNumber      string `json:"blockNumber"`
	BlockTimestamp   string `json:"blockTimestamp"`
	TxHash           string `json:"transactionHash"`
}

type RelationshipTheGraphAlphaResponse struct {
	Relationships []*RelationshipTheGraphAlpha `json:"relationshipCreateds"`
}

func (r *RelationshipTheGraphAlphaResponse) ToRelationships() []*Relationship {
	relationships := []*Relationship{}
	for _, relationship := range r.Relationships {
		relationships = append(relationships, relationship.ToRelationship())
	}

	return relationships
}

type RelationshipTypeTheGraphAlphaResponse struct {
	RelationshipTypes []*RelationshipTypeTheGraphAlpha `json:"relationshipTypes"`
}

func (r *RelationshipTheGraphAlpha) ToRelationship() *Relationship {
	return &Relationship{
		ID:           r.ID,
		Type:         r.RelationshipType,
		TypeId:       r.RelationshipId,
		SrcContract:  r.SrcAddress,
		SrcTokenId:   r.SrcId,
		DstContract:  r.DstAddress,
		DstTokenId:   r.DstId,
		RegisteredAt: r.BlockTimestamp,
		TxHash:       r.TxHash,
	}
}
