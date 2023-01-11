package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"examples-server.postgres.database.azure.com"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"5432"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"postgres"`
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"adminTerraform@examples-server"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"QAZwsx123"`
	AppHost          string `envconfig:"HOST" default:""`
	AppPort          string `envconfig:"PORT" default:"8080"`
	Network          string `envconfig:"NETWORK" default:"dev"`
}

func New() (*Config, error) {
	godotenv.Load()

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
