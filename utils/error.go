package utils

// AppError represents an application error with a status code and message
type AppError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}
