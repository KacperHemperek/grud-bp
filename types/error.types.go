package types

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e ApiError) Error() string {
	return e.Message
}
