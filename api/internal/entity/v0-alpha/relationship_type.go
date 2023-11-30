package v0alpha

import (
	"strconv"

	"github.com/project-nova/backend/pkg/utils"
)

type RelationshipType struct {
	Type            string `json:"type"`
	IpOrgId         string `json:"ipOrgId"`
	SrcContract     string `json:"srcContract"`
	SrcRelatable    int    `json:"srcRelatable"`
	SrcSubtypesMask int    `json:"srcSubtypesMask"`
	DstContract     string `json:"dstContract"`
	DstRelatable    int    `json:"dstRelatable"`
	DstSubtypesMask int    `json:"dstSubtypesMask"`
	RegisteredAt    string `json:"registeredAt"`
	TxHash          string `json:"txHash"`
}

type GetRelationshipTypeRequest struct {
	IpOrgId string `form:"ipOrgId"`
	RelType string `form:"relType"`
}

type GetRelationshipTypeResponse struct {
	RelationshipType *RelationshipType `json:"relationshipType"`
}

type ListRelationshipTypesRequest struct {
	IpOrgId string        `json:"ipOrgId"`
	Options *QueryOptions `json:"options"`
}

type ListRelationshipTypesResponse struct {
	RelationshipTypes []*RelationshipType `json:"relationshipTypes"`
}

type RelationshipTypeTheGraphAlpha struct {
	ID              string `json:"id"`
	RelType         string `json:"relType"`
	IpOrgId         string `json:"ipOrgId"`
	Src             string `json:"src"`
	SrcRelatable    int    `json:"srcRelatable"`
	SrcSubtypesMask string `json:"srcSubtypesMask"`
	Dst             string `json:"dst"`
	DstRelatable    int    `json:"dstRelatable"`
	DstSubtypesMask string `json:"dstSubtypesMask"`
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  string `json:"blockTimestamp"`
	TxHash          string `json:"transactionHash"`
}

type RelationshipTypeTheGraphAlphaResponse struct {
	RelationshipTypes []*RelationshipTypeTheGraphAlpha `json:"relationshipTypeSets"`
}

func (r *RelationshipTypeTheGraphAlphaResponse) ToRelationshipTypes() []*RelationshipType {
	relationshipTypes := []*RelationshipType{}
	for _, relationshipType := range r.RelationshipTypes {
		relationshipTypes = append(relationshipTypes, relationshipType.ToRelationshipType())
	}

	return relationshipTypes
}

func (r *RelationshipTypeTheGraphAlpha) ToRelationshipType() *RelationshipType {
	srcSubtypesMask, _ := strconv.Atoi(r.SrcSubtypesMask)
	dstSubtypesMask, _ := strconv.Atoi(r.DstSubtypesMask)
	return &RelationshipType{
		Type:            r.RelType,
		IpOrgId:         r.IpOrgId,
		SrcContract:     r.Src,
		SrcRelatable:    r.SrcRelatable,
		SrcSubtypesMask: srcSubtypesMask,
		DstContract:     r.Dst,
		DstRelatable:    r.DstRelatable,
		DstSubtypesMask: dstSubtypesMask,
		RegisteredAt:    utils.TimestampInSecondsToISO8601(r.BlockTimestamp),
		TxHash:          r.TxHash,
	}
}
