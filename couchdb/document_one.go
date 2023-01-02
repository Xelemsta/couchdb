package couchdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateDocument creates a document
func CreateDocument(ctx context.Context, id string, document interface{}, databaseName string) error {
	if document == nil {
		return fmt.Errorf(`document is nil`)
	}
	jsonDocument, err := json.Marshal(document)
	if err != nil {
		return err
	}
	_, err = PerformRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s/%s", databaseName, id),
		bytes.NewBuffer(jsonDocument),
	)
	return err
}
