package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

//Configuration for app
//include port db cred etc
type Configuration struct {
	PORT       string `envconfig:"PORT" default:":9999"`
	MongoURI   string `envconfig:"MONGO_URI" required:"true"`
	PGHOST     string `envconfig:"PG_HOST" required:"true"`
	PGPASS     string `envconfig:"PG_PASS" default:"postgres"`
	PGUSER     string `envconfig:"PG_USER" default:"admin"`
	PGPORT     string `envconfig:"PG_PORT" default:"5432"`
	PGURI      string `envconfig:"-"`
	PGDATABASE string `envconfig:"PG_DATABASE" default:":datingbot"`
}

//ConfigureAPI config of app
func ConfigureAPI() Configuration {
	var configuration Configuration
	err := envconfig.Process("", &configuration)
	if err != nil {
		panic(err)
	}
	return configurePGURI(configuration)
}

func configurePGURI(c Configuration) Configuration {
	pgURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.PGHOST, c.PGPORT, c.PGUSER, c.PGPASS, c.PGDATABASE)
	c.PGURI = pgURI
	return c
}
