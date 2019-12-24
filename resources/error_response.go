package resources

// ErrorResponse represents a response when there was an error handling the request
type ErrorResponse struct {
	Message string `json:"message"`
}
