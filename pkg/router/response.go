package router

// ErrorResponse is the object we use when returning errors
type ErrorResponse struct {
	Error string `json:"error"`
}
