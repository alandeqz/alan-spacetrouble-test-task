package config

import (
	"log/slog"

	"github.com/kelseyhightower/envconfig"

	"alan-tabeo-test-task/src/logging"
)

type Configuration struct {
	DSN           string `envconfig:"DB_DSN" default:"host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable"` // database DSN
	ListenAddress string `envconfig:"LISTEN_ADDR" default:"8080"`                                                                                  // HTTP listen address
	SchemaName    string `envconfig:"DB_SCHEMA_NAME" default:"alan_tabeo_test_task"`                                                               // database schema name
}

// NewConfiguration creates and returns a new configuration object.
func NewConfiguration() (*Configuration, error) {
	cfg := new(Configuration)

	if err := envconfig.Process("", cfg); err != nil {
		slog.Error("failed to get the environment variables", logging.Error, err.Error())

		return nil, err
	}

	return cfg, nil
}
