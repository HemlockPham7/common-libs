package dbutils

import (
	"errors"
	"strings"
)

var errorFilters = []func(err error) (bool, error){
	filterDuplicationUsername,
	filterDuplicationEmail,
	filterDuplicationType,
	filterRecordNotFound,
	filterForeignKeyError,
	filterRecordNotFoundType,
}

var (
	ErrDuplicationUsername = errors.New("username already exists")
	ErrDuplicationEmail    = errors.New("email already exists")
	ErrDuplicationType     = errors.New("duplicated type error")
	ErrRecordNotFound      = errors.New("record not found")
	ErrForeignKeyType      = errors.New("foreign key constraint")
	ErrRecordNotFoundType  = errors.New("record not found")
)

// CatchDBError normalizes database errors into application-level errors.
//
// It returns nil when err is nil. If the error matches a known database error,
// it returns the corresponding application-level error. Otherwise, it returns
// the original error unchanged.
//
// Parameters:
//   - err: the database error to normalize.
//
// Returns:
//   - A normalized application-level error, the original error if no filter matches,
//     or nil when err is nil.
func CatchDBError(err error) error {
	if err == nil {
		return nil
	}

	for _, filter := range errorFilters {
		match, filteredErr := filter(err)
		if match {
			return filteredErr
		}
	}

	return err
}

func filterDuplicationUsername(err error) (bool, error) {
	return strings.Contains(strings.ToLower(err.Error()), `duplicate key value violates unique constraint "uni_users_username"`), ErrDuplicationUsername
}

func filterDuplicationEmail(err error) (bool, error) {
	return strings.Contains(strings.ToLower(err.Error()), `duplicate key value violates unique constraint "uni_users_email"`), ErrDuplicationEmail
}

func filterDuplicationType(err error) (bool, error) {
	return strings.Contains(strings.ToLower(err.Error()), "unique constraint"), ErrDuplicationType
}

func filterRecordNotFound(err error) (bool, error) {
	return strings.Contains(strings.ToLower(err.Error()), "record not found"), ErrRecordNotFound
}

func filterForeignKeyError(err error) (bool, error) {
	return strings.Contains(strings.ToLower(err.Error()), "foreign key constraint"), ErrForeignKeyType
}

func filterRecordNotFoundType(err error) (bool, error) {
	return strings.Contains(strings.ToLower(err.Error()), "record not found"), ErrRecordNotFoundType
}
