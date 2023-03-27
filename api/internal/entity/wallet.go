package entity

type WalletProofResponse struct {
	Proof string `json:"proof"`
}

type SignMessageResponse struct {
	Message string `json:"message"`
}
