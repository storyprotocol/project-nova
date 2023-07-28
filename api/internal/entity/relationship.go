package entity

import (
	"time"

	"github.com/google/uuid"
)

type GetRelationshipResp struct {
	ID             string  `json:"id"`
	SourceContract string  `json:"sourceContract"`
	SrcId					 uint64  `json:"srcId"`
	DestContract   string  `json:"destContract"`
	DstId          uint64  `json:"destId"`
	TxHash         string  `json:"txHash"`
	Type           string `json:"type"`
}

type CreateRelationshipRequestBody struct {
	SourceContract *string `json:"sourceContract"`
	SrcId					 uint64  `json:"srcId"`
	DestContract   *string `json:"destContract"`
	DstId          uint64  `json:"destId"`
	TxHash         *string `json:"txHash"`
	Type           *string `json:"type"`
}


type RelationshipModel struct {
	ID             string    `gorm:"primaryKey;column:id" json:"id"`
	SourceContract string    `json:"sourceContract"`
	SrcId          uint64    `json:"srcId"`
	DestContract   string    `json:"destContract"`
	DstId          uint64    `json:"destId"`
	TxHash         string    `json:"txHash"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (RelationshipModel) TableName() string {
	return "relationships"
}

func (r *RelationshipModel) ToGetRelationshipResp() (*GetRelationshipResp, error) {
	resp := &GetRelationshipResp{
		ID:             r.ID,
		SourceContract: r.SourceContract,
		SrcId:          r.SrcId,
		DestContract:   r.DestContract,
		DstId:          r.DstId,
		TxHash:         r.TxHash,
		Type:           r.Type,
	}
	return resp, nil
}

func (c *CreateRelationshipRequestBody) ToRelationshipModel() *RelationshipModel {
	relationshipInfo := &RelationshipModel{
		ID: uuid.New().String(),
	}

	if c.SourceContract != nil && *c.SourceContract != "" {
		relationshipInfo.SourceContract = *c.SourceContract
	}

	relationshipInfo.SrcId = c.SrcId

	if c.DestContract != nil && *c.DestContract != "" {
		relationshipInfo.DestContract = *c.DestContract
	}

	relationshipInfo.DstId = c.DstId

	if c.TxHash != nil && *c.TxHash != "" {
		relationshipInfo.TxHash = *c.TxHash
	}

	if c.Type != nil && *c.Type != "" {
		relationshipInfo.Type = *c.Type
	}

	return relationshipInfo
}
