package couchdb_test

import (
	"couchdb/couchdb"
	"fmt"
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestError(t *testing.T) {
	err := couchdb.Error{
		Origin:     fmt.Errorf(`my_error_origin`),
		StatusCode: http.StatusBadRequest,
		ErrorCode:  "my_error_code",
		Reason:     "my_reason",
	}
	td.Cmp(t, err.Error(), "my_error_code - my_reason - my_error_origin")

	err = couchdb.Error{
		Origin:     nil,
		StatusCode: http.StatusBadRequest,
		ErrorCode:  "my_error_code",
		Reason:     "my_reason",
	}
	td.Cmp(t, err.Error(), "my_error_code - my_reason")

}
