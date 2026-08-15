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
