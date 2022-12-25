package couchdb

import "fmt"

// CouchdbError handles a HTTP response error
// coming from a couchdb request.
type Error struct {
	Origin     error
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error"`
	Reason     string `json:"reason"`
}

// Error implements error interface.
func (e *Error) Error() string {
	errorStr := fmt.Sprintf("%s - %s", e.ErrorCode, e.Reason)
	if e.Origin != nil {
		errorStr += " - " + e.Origin.Error()
	}
	return errorStr
}
