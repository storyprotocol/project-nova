package v0alpha

type Module struct {
	ID        string  `json:"id,omitempty"`
	IPOrgId   string  `json:"ipOrgId,omitempty"`
	Interface string  `json:"interface,omitempty"`
	PreHooks  []*Hook `json:"preHooks,omitempty"`
	PostHooks []*Hook `json:"postHooks,omitempty"`
}

/*
{
	id: string
	ipOrgId: string
	interface: string
	preHooks: Hook[]
	postHooks: Hook[]
	registeredAt: string // ISO 8601
	txHash: string
}
*/
