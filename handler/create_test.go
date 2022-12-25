package handler_test

import (
	"couchdb/handler"
	"couchdb/testutils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestCreateBadMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/create", strings.NewReader(`[]`))
	w := httptest.NewRecorder()
	err := handler.Create(w, req)
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "405 - method not allowed")
}

func TestCreateInvalidJSONDocuments(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(`[bad_json']`))
	w := httptest.NewRecorder()
	err := handler.Create(w, req)
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "400 - invalid json documents - invalid character 'b' looking for beginning of value")
}

func TestCreateNoDocumentToProcess(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(`[]`))
	w := httptest.NewRecorder()
	err := handler.Create(w, req)
	td.CmpError(t, err)
	td.Cmp(t, err.Error(), "400 - no documents to process")
}

func TestCreateOK(t *testing.T) {
	err := testutils.Init()
	td.CmpNoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(`[{"test1":"test1"}, {"test2":"test2"}]`))
	w := httptest.NewRecorder()
	err = handler.Create(w, req)
	td.CmpNoError(t, err)
}
