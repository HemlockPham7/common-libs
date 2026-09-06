package sqldb

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	User     string `envconfig:"DB_USER" default:"admin"`
	Password string `envconfig:"DB_PASSWORD" default:"admin"`
	DBName   string `envconfig:"DB_NAME" default:"bookmark"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
}

// newConfig loads SQL database configuration from environment variables
// using the specified environment variable prefix.
//
// Parameters:
//   - envPrefix: the prefix used to load database configuration values.
//
// Returns:
//   - The loaded database configuration, or an error if the configuration cannot be processed.
func newConfig(envPrefix string) (*config, error) {
	cfg := &config{}
	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// GetDSN builds a PostgreSQL connection string from the database configuration.
//
// Returns:
//   - A PostgreSQL DSN containing the configured host, user, password, database name, and port.
func (cfg *config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
}
