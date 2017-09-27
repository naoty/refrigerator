package config

import (
	"github.com/kelseyhightower/envconfig"
)

// New creates a new Config.
func New() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}

// Config is a configuration for this app.
type Config struct {
	DBUser     string `default:"root"`
	DBPassword string `envconfig:"MYSQL_ROOT_PASSWORD"`
}
