package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type db struct {
	Driver   string `envconfig:"APP_DB_DRIVER" default:"memory"`
	Host     string `envconfig:"APP_DB_HOST" default:"localhost"`
	Port     string `envconfig:"APP_DB_PORT" default:"5432"`
	User     string `envconfig:"APP_DB_USER" default:"db"`
	Password string `envconfig:"APP_DB_PASSWORD" default:"db"`
	Database string `envconfig:"APP_DB_DATABASE" default:"banking"`
	Debug    bool   `envconfig:"APP_DB_DEBUG" default:"true"`
}

// Config app config
type Config struct {
	DB db
}

// Environ get APP config
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}
