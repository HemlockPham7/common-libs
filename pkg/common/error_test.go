package common

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		inputError    error
		expectedPanic bool
	}{
		{
			name:          "error is nil",
			inputError:    nil,
			expectedPanic: false,
		},
		{
			name:          "error is not nil",
			inputError:    errors.New("test error"),
			expectedPanic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.expectedPanic {
				assert.Panics(t, func() {
					HandleError(tc.inputError)
				})
			} else {
				assert.NotPanics(t, func() {
					HandleError(tc.inputError)
				})
			}
		})
	}
}
