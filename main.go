package main

import (
	"fmt"
	"log"
	"net/http"

	"couchdb/configuration"
	"couchdb/couchdb"
	"couchdb/handler"
)

const cfgFilename = "configuration.yaml"

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	initCouchdb()
	initCfg()
	initAPI()
}

func initCouchdb() {
	log.Print("initializing couchdb")
	err := couchdb.Init()
	if err != nil {
		panic(err)
	}
	log.Print("initialized couchdb")
}

func initCfg() {
	log.Print("initializing config")
	err := configuration.Init(cfgFilename)
	if err != nil {
		panic(err)
	}
	log.Print("initialized config")
}

func initAPI() {
	log.Printf("listening on %d", configuration.HTTPPort())
	http.Handle("/create", logIncomingRequest(handler.CreateHandler()))
	http.ListenAndServe(
		fmt.Sprintf(":%d", configuration.HTTPPort()),
		http.DefaultServeMux,
	)
}

func logIncomingRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("incoming request: %s - %s - %+v", r.Method, r.URL.Path, r.Header)
		next.ServeHTTP(w, r)
	})
}
