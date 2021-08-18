package config

import (
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/errors"
	"github.com/kelseyhightower/envconfig"
)

// MainConfig - global config variable
var MainConfig *Config

// Config constants
const DefaultChanCapacity = 2

// Main config struct
type Config struct {
	PostgreSQLConfig
}

// Main DB config struct
type PostgreSQLConfig struct {
	PostgresHost   string `envconfig:"postgres_host" required:"true"`
	PostgresPort   string `envconfig:"postgres_port" required:"true"`
	PostgresName   string `envconfig:"postgres_name" required:"true"`
	PostgresSchema string `envconfig:"postgres_schema" required:"true"`
	PostgresUser   string `envconfig:"postgres_user" required:"true"`
	PostgresPass   string `envconfig:"postgres_pass" required:"true"`
}

// Initialize service configuration from ENV
func InitConfigs() error {
	var cfg Config

	err := envconfig.Process("APP_CONFIG", &cfg)
	if err != nil {
		return errors.ENVReadError.SetDevMessage(err.Error())
	}

	MainConfig = &cfg
	return nil
}
