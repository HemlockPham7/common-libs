package sqldb

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func MigrateSQLDB(db *gorm.DB, migrationPath string, mode string, steps int) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	pgDriver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, db.Name(), pgDriver)
	if err != nil {
		return err
	}

	return migrateSchema(m, mode, steps)
}

func migrateSchema(m *migrate.Migrate, mode string, steps int) error {
	var migrationErr error

	switch mode {
	case "up":
		migrationErr = m.Up()
	case "steps":
		if steps == 0 {
			return errors.New("[Database migration] steps must not be 0. Please use a positive number to migrate")
		}
		migrationErr = m.Steps(steps)
	case "down":
		migrationErr = m.Down()
	default:
		return errors.New("[Database migration] invalid mode. Please use 'up', 'steps', or 'down'")
	}

	if migrationErr != nil && !errors.Is(migrationErr, migrate.ErrNoChange) {
		return fmt.Errorf("[Database migration] Error: %s", migrationErr.Error())
	}

	return nil
}
