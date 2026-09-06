package sqldb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewClient creates a GORM PostgreSQL client using configuration loaded from environment variables.
//
// Parameters:
//   - envPrefix: the prefix used to load database configuration values.
//
// Returns:
//   - A configured GORM PostgreSQL client, or an error if the configuration cannot be processed
//     or the database connection cannot be initialized.
func NewClient(envPrefix string) (*gorm.DB, error) {
	cfg, err := newConfig(envPrefix)
	if err != nil {
		return nil, err
	}

	dsn := cfg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
