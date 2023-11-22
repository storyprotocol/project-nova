package v0alpha

import "github.com/project-nova/backend/pkg/utils"

type Hook struct {
	ID           string `json:"id,omitempty"`
	ModuleId     string `json:"moduleId,omitempty"`
	Interface    string `json:"interface,omitempty"`
	HookType     int64  `json:"hookType"`
	RegistryKey  string `json:"registryKey,omitempty"`
	RegisteredAt string `json:"registeredAt,omitempty"`
	TxHash       string `json:"txHash,omitempty"`
}

type GetHookResponse struct {
	Hook *Hook `json:"hook"`
}

type ListHooksRequest struct {
	ModuleId *string       `json:"moduleId"`
	Options  *QueryOptions `json:"options"`
}

type ListHookRequest struct {
	ModuleId *string       `json:"moduleId"`
	Options  *QueryOptions `json:"options"`
}

type ListHooksResponse struct {
	Hooks []*Hook `json:"hooks"`
}

type HookTheGraphAlpha struct {
	ID             string `json:"id"`
	ModuleId       string `json:"moduleId"`
	Type           int64  `json:"type"`
	RegistryKey    string `json:"registryKey"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
	TxHash         string `json:"transactionHash"`
}

type HookTheGraphAlphaResponse struct {
	Hooks []*HookTheGraphAlpha `json:"hookRegistereds"`
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
		ModuleId:     h.ModuleId,
		HookType:     h.Type,
		RegistryKey:  h.RegistryKey,
		RegisteredAt: utils.TimestampInSecondsToISO8601(h.BlockTimestamp),
		TxHash:       h.TxHash,
	}
}
