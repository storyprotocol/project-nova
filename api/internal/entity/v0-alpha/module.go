package v0alpha

type Module struct {
	ID        string  `json:"id,omitempty"`
	IPOrgId   string  `json:"ipOrgId,omitempty"`
	Interface string  `json:"interface,omitempty"`
	PreHooks  []*Hook `json:"preHooks,omitempty"`
	PostHooks []*Hook `json:"postHooks,omitempty"`
}

type GetModuleResponse struct {
	Module *Module `json:"module"`
}

type ListModulesResponse struct {
	Modules []*Module `json:"modules"`
}
