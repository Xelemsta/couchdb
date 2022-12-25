package configuration_test

import (
	"couchdb/configuration"
	"couchdb/testutils"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestReadConfiguration(t *testing.T) {
	err := testutils.Init()
	td.CmpNoError(t, err)

	td.Cmp(t, configuration.DBHost(), "localhost")
	td.Cmp(t, configuration.DBPort(), 5984)
	td.Cmp(t, configuration.DBUser(), "admin")
	td.Cmp(t, configuration.DBPassword(), "password")
	td.Cmp(t, configuration.HTTPPort(), 10000)
}

func TestValidConfiguration(t *testing.T) {
	dataset := map[string]configuration.Configuration{
		"invalid_host": {
			CouchDB: configuration.CouchDB{
				Host:     "",
				Port:     2000,
				User:     "test",
				Password: "psswd",
			},
			HTTP: configuration.HTTP{
				Port: 10000,
			},
		},
		"invalid_port": {
			CouchDB: configuration.CouchDB{
				Host:     "localhost",
				Port:     0,
				User:     "test",
				Password: "psswd",
			},
			HTTP: configuration.HTTP{
				Port: 10000,
			},
		},
		"invalid_user": {
			CouchDB: configuration.CouchDB{
				Host:     "localhost",
				Port:     2000,
				User:     "",
				Password: "psswd",
			},
			HTTP: configuration.HTTP{
				Port: 10000,
			},
		},
		"invalid_password": {
			CouchDB: configuration.CouchDB{
				Host:     "localhost",
				Port:     2000,
				User:     "test",
				Password: "",
			},
			HTTP: configuration.HTTP{
				Port: 10000,
			},
		},
		"invalid_http_port": {
			CouchDB: configuration.CouchDB{
				Host:     "localhost",
				Port:     0,
				User:     "test",
				Password: "psswd",
			},
			HTTP: configuration.HTTP{
				Port: 0,
			},
		},
	}

	for msg, cfg := range dataset {
		valid, err := cfg.IsValid()
		td.CmpFalse(t, valid, msg)
		td.CmpError(t, err, msg)
	}

	validCfg := configuration.Configuration{
		CouchDB: configuration.CouchDB{
			Host:     "localhost",
			Port:     2000,
			User:     "test",
			Password: "psswd",
		},
		HTTP: configuration.HTTP{
			Port: 10000,
		},
	}

	valid, err := validCfg.IsValid()
	td.CmpTrue(t, valid)
	td.CmpNoError(t, err)
}
