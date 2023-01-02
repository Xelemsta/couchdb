package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"couchdb/couchdb"
)

// Create handles "create" endpoint logic
// which creates a database and stores
// provided documents
func CreateHandler() http.Handler {
	return appHandler(Create)
}

func Create(w http.ResponseWriter, r *http.Request) error {
	log.Print("starting")

	ctx := r.Context()

	if r.Method != http.MethodPost {
		return apiError(&APIError{
			Code:    http.StatusMethodNotAllowed,
			Origin:  nil,
			Message: "method not allowed",
		})
	}

	var documents []interface{}
	err := json.NewDecoder(r.Body).Decode(&documents)
	if err != nil {
		return apiError(&APIError{
			Code:    http.StatusBadRequest,
			Origin:  err,
			Message: "invalid json documents",
		})
	}

	nbOfDocs := len(documents)
	if nbOfDocs == 0 {
		log.Print(`no documents to handle`)
		return apiError(&APIError{
			Code:    http.StatusBadRequest,
			Origin:  nil,
			Message: "no documents to process",
		})
	}
	log.Printf(`found %d documents to push`, nbOfDocs)

	uniqueDBName := "a" + strconv.FormatInt(time.Now().Unix(), 10)
	log.Printf("creating db with name %s", uniqueDBName)
	err = couchdb.CreateDatabase(ctx, uniqueDBName)
	if err != nil {
		return apiError(err)
	}
	log.Printf("successfully created db")

	log.Printf("pushing documents")
	for index, doc := range documents {
		err := couchdb.CreateDocument(ctx, strconv.Itoa(index+1), doc, uniqueDBName)
		if err != nil {
			return apiError(err)
		}
	}
	log.Printf("successfully pushed documents")

	log.Print("all done")
	return nil
}

func apiError(err error) error {
	switch v := err.(type) {
	case *couchdb.Error:
		return &APIError{
			Origin:  v.Origin,
			Message: v.Error(),
			Code:    v.StatusCode,
		}
	case *APIError:
		return err
	default:
		return &APIError{
			Origin:  nil,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
}
