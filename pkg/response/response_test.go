package response

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type testRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestResponse_InputFieldError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		err             error
		expectedMessage Message
	}{
		{
			name:            "error is not validation error",
			err:             errors.New("some error"),
			expectedMessage: InputErrResponse,
		},
		{
			name: "validation error",
			err: func() error {
				req := testRequest{
					Name:  "",
					Email: "invalid-email",
				}

				validate := validator.New()
				return validate.Struct(req)
			}(),
			expectedMessage: Message{
				Message: "Input error",
				Details: []string{
					"Name is invalid (required)",
					"Email is invalid (email)",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := InputFieldError(tc.err)

			assert.Equal(t, tc.expectedMessage, result)
		})
	}
}
