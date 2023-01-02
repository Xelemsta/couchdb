package handler_test

import (
	"couchdb/handler"
	"fmt"
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestError(t *testing.T) {
	err := handler.APIError{
		Origin:  fmt.Errorf(`my_error_origin`),
		Code:    http.StatusBadRequest,
		Message: "my_msg",
	}
	td.Cmp(t, err.Error(), "400 - my_msg - my_error_origin")

	err = handler.APIError{
		Origin:  nil,
		Code:    http.StatusBadRequest,
		Message: "my_msg",
	}
	td.Cmp(t, err.Error(), "400 - my_msg")

}
