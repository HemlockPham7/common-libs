package response

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Message struct {
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

var (
	InternalErrResponse = Message{
		Message: "Processing Error",
		Details: nil,
	}
	InputErrResponse = Message{
		Message: "Input Error",
		Details: nil,
	}
	InstanceErrResponse = Message{
		Message: "Instance is not ready!",
		Details: nil,
	}
	UnauthorizedResponse = Message{
		Message: "Unauthorized",
		Details: nil,
	}
)

// InputFieldError converts a validation error into a standardized input error response.
//
// If err is not a validator.ValidationErrors, it returns the default InputErrResponse.
// Otherwise, it returns a response containing the validation errors for each invalid field.
//
// Parameters:
//   - err: the validation error to convert.
//
// Returns:
//   - A standardized Message containing the validation error details.
func InputFieldError(err error) Message {
	if ok := errors.As(err, &validator.ValidationErrors{}); !ok { // khi fail input binding validation
		return InputErrResponse
	}

	var errs []string
	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, err.Field()+" is invalid ("+err.Tag()+")")
	}

	return Message{
		Message: "Input error",
		Details: errs,
	}
}
