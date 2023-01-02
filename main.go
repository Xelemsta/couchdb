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
	http.Handle("/create", handler.CreateHandler())
	http.ListenAndServe(
		fmt.Sprintf(":%d", configuration.HTTPPort()),
		http.DefaultServeMux,
	)
}
