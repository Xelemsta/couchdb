package couchdb_test

import (
	"context"
	"couchdb/couchdb"
	"couchdb/testutils"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestBulkDocuments(t *testing.T) {
	err := testutils.Init()
	td.CmpNoError(t, err)

	dbName := "db1"
	ctx := context.Background()
	documents := []interface{}{}

	defer func() {
		_, _ = couchdb.PerformRequest(
			ctx,
			http.MethodDelete,
			fmt.Sprintf("%s", dbName),
			nil,
		)

		for _, doc := range documents {
			docMap := doc.(map[string]string)
			_, _ = couchdb.PerformRequest(
				ctx,
				http.MethodDelete,
				fmt.Sprintf("%s/%s", dbName, docMap["_id"]),
				nil,
			)
		}
	}()

	for i := 1; i < 205; i++ {
		strIndex := strconv.Itoa(i + 1)
		documents = append(documents, map[string]string{
			"test": "toto" + strIndex,
			"_id":  strIndex,
		})
	}

	err = couchdb.CreateDocuments(ctx, nil, dbName)
	td.CmpNoError(t, err)

	err = couchdb.CreateDatabase(ctx, dbName)
	td.CmpNoError(t, err)

	err = couchdb.CreateDocuments(ctx, documents, dbName)
	td.CmpNoError(t, err)

	for _, doc := range documents {
		docMap := doc.(map[string]string)
		_, err := couchdb.PerformRequest(
			ctx,
			http.MethodGet,
			fmt.Sprintf("%s/%s", dbName, docMap["_id"]),
			nil,
		)
		td.CmpNoError(t, err, fmt.Sprintf("not found %s", docMap["_id"]))
	}
}
