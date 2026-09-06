package sqldb

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMockDB creates an in-memory SQLite database for testing.
//
// The database uses silent logging to keep test output clean.
//
// Parameters:
//   - t: the testing instance used to report database initialization failures.
//
// Returns:
//   - An in-memory SQLite database configured for testing.
func InitMockDB(t *testing.T) *gorm.DB {
	dsn := ":memory:"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to create test db: %v", err)
	}
	return db
}
