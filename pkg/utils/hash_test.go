package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		inputPassword string

		expectedError error
	}{
		{
			name: "success",

			inputPassword: "password",

			expectedError: nil,
		},
		{
			name: "error",

			inputPassword: "password_is_too_much_long_so_that_it_can_trigger_the_error_of_exceeding_72_characters",

			expectedError: ErrHashFailed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testService := NewHasher()
			hashedPassword, err := testService.Hash(tc.inputPassword)

			assert.ErrorIs(t, err, tc.expectedError)
			if err == nil {
				assert.NotEmpty(t, hashedPassword)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		inputPassword       string
		inputHashedPassword string

		expectedMatch bool
	}{
		{
			name: "success",

			inputPassword:       "my_secure_password",
			inputHashedPassword: "$2a$10$yIIizEHMEKSm.OARDrSjHe4otTolPuCjjEy6IQ3RtRny3ZB7ToN.e",

			expectedMatch: true,
		},
		{
			name: "success",

			inputPassword:       "my_secure_password",
			inputHashedPassword: "",

			expectedMatch: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testService := NewHasher()
			match := testService.Compare(tc.inputHashedPassword, tc.inputPassword)

			assert.Equal(t, tc.expectedMatch, match)
		})
	}
}
