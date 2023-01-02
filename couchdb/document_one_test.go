package couchdb_test

import (
	"context"
	"couchdb/couchdb"
	"couchdb/testutils"
	"fmt"
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestAddDocument(t *testing.T) {
	err := testutils.Init()
	td.CmpNoError(t, err)

	dbName := "db1"
	ctx := context.Background()
	id := "my_id"

	defer func() {
		_, _ = couchdb.PerformRequest(
			ctx,
			http.MethodDelete,
			fmt.Sprintf("%s", dbName),
			nil,
		)

		_, _ = couchdb.PerformRequest(
			ctx,
			http.MethodDelete,
			fmt.Sprintf("%s/%s", dbName, id),
			nil,
		)
	}()

	err = couchdb.CreateDocument(ctx, id, nil, "my_database")
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "document is nil")

	err = couchdb.CreateDocument(ctx, id, `{"test":"test"}`, dbName)
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "not_found - Database does not exist.")

	err = couchdb.CreateDatabase(ctx, dbName)
	td.CmpNoError(t, err)

	err = couchdb.CreateDocument(ctx, id, `{"test":"test"}`, dbName)
	td.CmpNoError(t, err)
}
