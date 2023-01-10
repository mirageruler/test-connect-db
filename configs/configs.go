package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"5435"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"postgres"`
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"postgres"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
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
