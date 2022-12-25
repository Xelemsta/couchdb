package couchdb

import (
	"context"
	"fmt"
	"net/http"
)

// CreateDatabase creates a couchdb database
func CreateDatabase(ctx context.Context, name string) error {
	_, err := PerformRequest(ctx, http.MethodPut, fmt.Sprintf("%s", name), nil)
	return err
}
