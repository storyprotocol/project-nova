package http

type Error struct {
	Message string `json:"message"`
}

// Error implements the built-in error interface type.
func (e Error) Error() string {
	return e.Message
}
