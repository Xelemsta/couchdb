package couchdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"couchdb/configuration"
)

var client *http.Client

// Init inits basic http client for couchdb.
func Init() error {
	client = &http.Client{}
	return nil
}

// PerformRequest wraps all couchdb request
func PerformRequest(ctx context.Context, method string, path string, requestBody interface{}) (interface{}, error) {
	var requestJSON []byte
	var err error
	if requestBody != nil {
		requestJSON, err = json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("http://%s:%d/%s", configuration.DBHost(), configuration.DBPort(), path),
		bytes.NewReader(requestJSON),
	)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(configuration.DBUser(), configuration.DBPassword())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var couchdbError *Error
		err = json.NewDecoder(resp.Body).Decode(&couchdbError)
		if err != nil {
			return nil, err
		}
		couchdbError.StatusCode = resp.StatusCode
		return nil, couchdbError
	}

	var responseBody interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	return responseBody, err
}
