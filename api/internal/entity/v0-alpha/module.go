package v0alpha

import "github.com/project-nova/backend/pkg/utils"

type Module struct {
	ID        string  `json:"id,omitempty"`
	IPOrgId   string  `json:"ipOrgId,omitempty"`
	ModuleKey string  `json:"moduleKey,omitempty"`
	Interface string  `json:"interface,omitempty"`
	PreHooks  []*Hook `json:"preHooks,omitempty"`
	PostHooks []*Hook `json:"postHooks,omitempty"`
}

type GetModuleResponse struct {
	Module *Module `json:"module"`
}

type ListModulesRequest struct {
	IpOrgId *string       `json:"ipOrgId"`
	Options *QueryOptions `json:"options"`
}

func (l *ListModulesRequest) Validate() bool {
	if l.IpOrgId != nil && !utils.IsValidAddress(*l.IpOrgId) {
		return false
	}

	return true
}

type ListModulesResponse struct {
	Modules []*Module `json:"modules"`
}

type ModuleTheGraphAlpha struct {
	ID             string `json:"id"`
	IPOrgId        string `json:"ipOrgId"`
	ModuleKey      string `json:"moduleKey"`
	ModuleId       string `json:"moduleId"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
	TxHash         string `json:"transactionHash"`
}

type ModuleTheGraphAlphaResponse struct {
	Modules []*ModuleTheGraphAlpha `json:"moduleRegisterreds"`
}

func (m *ModuleTheGraphAlphaResponse) ToModules() []*Module {
	modules := []*Module{}
	for _, module := range m.Modules {
		modules = append(modules, module.ToModule())
	}

	return modules
}

func (m *ModuleTheGraphAlpha) ToModule() *Module {
	return &Module{
		ID:        m.ModuleId,
		IPOrgId:   m.IPOrgId,
		ModuleKey: m.ModuleKey,
	}
}
