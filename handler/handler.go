package handler

import (
	"net/http"

	"log"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		switch e := err.(type) {
		case *APIError:
			log.Print(e.Error())
			http.Error(w, e.Error(), e.Code)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
