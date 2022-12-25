package testutils

import (
	"couchdb/configuration"
	"couchdb/couchdb"
	"os"
)

func Init() error {
	cfgFileName := "configuration_test.yaml"
	file_test, err := os.OpenFile(cfgFileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer file_test.Close()
	defer os.Remove(file_test.Name())

	_, err = file_test.Write([]byte(`
couchdb:
  host: localhost
  port: 5984
  user: admin
  password: password
http:
  port: 10000
`))
	if err != nil {
		return err
	}

	err = configuration.Init(cfgFileName)
	if err != nil {
		return err
	}

	return couchdb.Init()
}
