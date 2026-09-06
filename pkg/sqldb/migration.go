package sqldb

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// MigrateSQLDB runs database schema migrations using the specified migration mode.
//
// Parameters:
//   - db: the GORM database connection used to run the migrations.
//   - migrationPath: the path to the migration files.
//   - mode: the migration mode, which can be "up", "steps", or "down".
//   - steps: the number of migrations to apply when mode is "steps".
//
// Returns:
//   - An error if the database connection, migration initialization, or migration fails.
func MigrateSQLDB(db *gorm.DB, migrationPath string, mode string, steps int) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// golang-migrate requires a driver to know: "Which PostgreSQL instance am I working with?" — the pgDriver serves as the adapter between:
	pgDriver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	// Initialize the migration instance
	m, err := migrate.NewWithDatabaseInstance(migrationPath, db.Name(), pgDriver)
	if err != nil {
		return err
	}

	return migrateSchema(m, mode, steps)
}

// migrateSchema executes the requested migration operation.
//
// Parameters:
//   - m: the initialized migration instance.
//   - mode: the migration mode, which can be "up", "steps", or "down".
//   - steps: the number of migrations to apply when mode is "steps".
//
// Returns:
//   - An error if the migration mode is invalid or the migration fails.
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
