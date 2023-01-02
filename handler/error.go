package handler

import "fmt"

// APIError handles API response error
type APIError struct {
	Origin  error
	Message string
	Code    int
}

// Error implement error interface
func (e *APIError) Error() string {
	errorStr := fmt.Sprintf(`%d - %s`, e.Code, e.Message)
	if e.Origin != nil {
		errorStr += " - " + e.Origin.Error()
	}
	return errorStr
}
