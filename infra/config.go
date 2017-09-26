package infra

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// DefaultConfig is a configuration for this app set by default.
var DefaultConfig Config

func init() {
	err := envconfig.Process("", &DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// Config is a configuration for this app.
type Config struct {
	DBUser     string `default:"root"`
	DBPassword string `envconfig:"MYSQL_ROOT_PASSWORD"`
}
