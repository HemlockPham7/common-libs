package dbutils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCatchDBError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		inputErr    error
		expectedErr error
	}{
		{
			name:        "nil error",
			inputErr:    nil,
			expectedErr: nil,
		},
		{
			name:        "duplication error",
			inputErr:    errors.New("unique constraint violation"),
			expectedErr: ErrDuplicationType,
		},
		{
			name:        "record not found error",
			inputErr:    gorm.ErrRecordNotFound,
			expectedErr: ErrRecordNotFoundType,
		},
		{
			name:        "other error",
			inputErr:    errors.New("some other database error"),
			expectedErr: errors.New("some other database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotErr := CatchDBError(tc.inputErr)
			assert.Equal(t, tc.expectedErr, gotErr)
		})
	}
}
