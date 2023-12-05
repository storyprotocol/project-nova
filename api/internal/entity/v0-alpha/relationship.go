package v0alpha

import "github.com/project-nova/backend/pkg/utils"

type Relationship struct {
	ID           string `json:"id,omitempty"`
	Type         string `json:"type,omitempty"`
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

func (l *ListRelationshipRequest) Validate() bool {
	if !utils.IsValidAddress(l.Contract) {
		return false
	}

	if !utils.IsValidNumberString(l.TokenId) {
		return false
	}

	return true
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

func (r *RelationshipTheGraphAlpha) ToRelationship() *Relationship {
	return &Relationship{
		ID:           r.RelationshipId,
		Type:         r.RelationshipType,
		SrcContract:  r.SrcAddress,
		SrcTokenId:   r.SrcId,
		DstContract:  r.DstAddress,
		DstTokenId:   r.DstId,
		RegisteredAt: utils.TimestampInSecondsToISO8601(r.BlockTimestamp),
		TxHash:       r.TxHash,
	}
}
