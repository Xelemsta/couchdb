package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

const Filename = "configuration"

var config Configuration

type Configuration struct {
	CouchDB CouchDB
	HTTP    HTTP
}

type CouchDB struct {
	Host     string
	Port     int
	User     string
	Password string
}

type HTTP struct {
	Port int
}

func DBHost() string {
	return config.CouchDB.Host
}

func DBPort() int {
	return config.CouchDB.Port
}

func DBUser() string {
	return config.CouchDB.User
}

func DBPassword() string {
	return config.CouchDB.Password
}

func HTTPPort() int {
	return config.HTTP.Port
}

func Init(filename string) error {
	v := viper.GetViper()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	err := v.Unmarshal(&config)
	if err != nil {
		return err
	}

	_, err = config.IsValid()
	return err
}

func (c *Configuration) IsValid() (bool, error) {
	if c.CouchDB.Host == "" {
		return false, fmt.Errorf(`empty couchdb host`)
	}

	if c.CouchDB.User == "" {
		return false, fmt.Errorf(`empty couchdb user`)
	}

	if c.CouchDB.Password == "" {
		return false, fmt.Errorf(`empty couchdb password`)
	}

	if c.CouchDB.Port == 0 {
		return false, fmt.Errorf(`empty couchdb port`)
	}

	if c.HTTP.Port == 0 {
		return false, fmt.Errorf(`empty HTTP port`)
	}

	return true, nil
}
