package v0alpha

type Hook struct {
	ID           string `json:"id,omitempty"`
	ModuleId     string `json:"moduleId,omitempty"`
	Interface    string `json:"interface,omitempty"`
	RegisteredAt string `json:"registeredAt,omitempty"`
	TxHash       string `json:"txHash,omitempty"`
}

/*
{
	id: string
	moduleId: string
	interface: string
	createdAt: string // ISO 8601
}
*/
