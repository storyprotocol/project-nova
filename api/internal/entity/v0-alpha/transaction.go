package v0alpha

import "github.com/project-nova/backend/pkg/utils"

type Transaction struct {
	ID           string       `json:"id,omitempty"`
	TxHash       string       `json:"txHash,omitempty"`
	IPOrgId      string       `json:"ipOrgId,omitempty"`
	ResourceId   string       `json:"resourceId,omitempty"`
	ResourceType ResourceType `json:"resourceType,omitempty"`
	ActionType   ActionType   `json:"actionType,omitempty"`
	Initiator    string       `json:"initiator,omitempty"`
	CreatedAt    string       `json:"createdAt,omitempty"`
}

type ResourceType string

var ResourceTypes = struct {
	IPOrg            ResourceType
	IPAsset          ResourceType
	License          ResourceType
	Relationship     ResourceType
	RelationshipType ResourceType
	Module           ResourceType
	Hook             ResourceType
	Dispute          ResourceType
}{
	IPOrg:            "IPOrg",
	IPAsset:          "IPAsset",
	License:          "License",
	Relationship:     "Relationship",
	RelationshipType: "RelationshipType",
	Module:           "Module",
	Hook:             "Hook",
	Dispute:          "Dispute",
}

type ActionType string

var ActionTypes = struct {
	Create     ActionType
	Register   ActionType
	Unregister ActionType
	Configure  ActionType
}{
	Create:     "Create",
	Register:   "Register",
	Unregister: "Unregister",
	Configure:  "Configure",
}

type GetTransactionResponse struct {
	Transaction *Transaction `json:"transaction"`
}

type ListTransactionsRequest struct {
	IpOrgId *string       `json:"ipOrgId"`
	Options *QueryOptions `json:"options"`
}

func (l *ListTransactionsRequest) Validate() bool {
	if l.IpOrgId != nil && !utils.IsValidAddress(*l.IpOrgId) {
		return false
	}

	return true
}

type ListTransactionsResponse struct {
	Transactions []*Transaction `json:"transactions"`
}

type TransactionTheGraphAlpha struct {
	ID             string `json:"id"`
	Initiator      string `json:"initiator"`
	IpOrgId        string `json:"ipOrgId"`
	ResourceType   string `json:"resourceType"`
	ResourceId     string `json:"resourceId"`
	ActionType     string `json:"actionType"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
	TxHash         string `json:"transactionHash"`
}

type TransactionTheGraphAlphaResponse struct {
	Transactions []*TransactionTheGraphAlpha `json:"transactions"`
}

func (t *TransactionTheGraphAlphaResponse) ToTransactions() []*Transaction {
	transactions := make([]*Transaction, len(t.Transactions))
	for i, transaction := range t.Transactions {
		transactions[i] = transaction.ToTransaction()
	}
	return transactions
}

func (t *TransactionTheGraphAlpha) ToTransaction() *Transaction {
	return &Transaction{
		ID:           t.ID,
		TxHash:       t.TxHash,
		IPOrgId:      t.IpOrgId,
		ResourceId:   t.ResourceId,
		ResourceType: ResourceType(t.ResourceType),
		ActionType:   ActionType(t.ActionType),
		Initiator:    t.Initiator,
		CreatedAt:    utils.TimestampInSecondsToISO8601(t.BlockTimestamp),
	}
}
