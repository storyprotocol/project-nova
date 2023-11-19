package v0alpha

type Hook struct {
	ID           string `json:"id,omitempty"`
	ModuleId     string `json:"moduleId,omitempty"`
	Interface    string `json:"interface,omitempty"`
	RegisteredAt string `json:"registeredAt,omitempty"`
	TxHash       string `json:"txHash,omitempty"`
}

type GetHookResponse struct {
	Hook *Hook `json:"hook"`
}

type ListHooksResponse struct {
	Hooks []*Hook `json:"hooks"`
}

type HookTheGraphAlpha struct {
	ID             string `json:"id"`
	HookId         string `json:"hookId"`
	ModuleKey      string `json:"moduleKey"`
	HookType       string `json:"hookType"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
	TxHash         string `json:"transactionHash"`
}

type HookTheGraphAlphaResponse struct {
	Hooks []*HookTheGraphAlpha `json:"hooks"`
}

func (h *HookTheGraphAlphaResponse) ToHooks() []*Hook {
	hooks := []*Hook{}
	for _, hook := range h.Hooks {
		hooks = append(hooks, hook.ToHook())
	}

	return hooks
}

func (h *HookTheGraphAlpha) ToHook() *Hook {
	return &Hook{
		ID:           h.ID,
		ModuleId:     h.ModuleKey,
		Interface:    h.HookType,
		RegisteredAt: h.BlockTimestamp,
		TxHash:       h.TxHash,
	}
}
