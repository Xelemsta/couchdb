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

func TestAddDatabase(t *testing.T) {
	err := testutils.Init()
	td.CmpNoError(t, err)

	validDBName := "valid_database_name"
	ctx := context.Background()

	defer func() {
		_, _ = couchdb.PerformRequest(
			ctx,
			http.MethodDelete,
			fmt.Sprintf("%s", validDBName),
			nil,
		)
	}()

	err = couchdb.CreateDatabase(ctx, "1invalid_database_name")
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "illegal_database_name - Name: '1invalid_database_name'. Only lowercase characters (a-z), digits (0-9), and any of the characters _, $, (, ), +, -, and / are allowed. Must begin with a letter.")

	err = couchdb.CreateDatabase(ctx, validDBName)
	td.CmpNoError(t, err)

	err = couchdb.CreateDatabase(ctx, validDBName)
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "file_exists - The database could not be created, the file already exists.")
}
