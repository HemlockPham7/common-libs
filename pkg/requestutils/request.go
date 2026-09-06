package requestutils

import (
	"errors"
	"io"
	"net/http"

	"github.com/HemlockPham7/common-libs/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	InputValidator = validator.New(validator.WithRequiredStructEnabled())
)

// BindInputFromRequest binds and validates request input from the JSON body,
// URI parameters, query parameters, and request headers.
//
// JSON body binding is skipped for GET requests. The request is aborted with
// HTTP 400 if binding or validation fails.
//
// Parameters:
//   - c: the Gin context containing the incoming HTTP request.
//
// Returns:
//   - A pointer to the populated and validated input, or an error if binding or validation fails.
func BindInputFromRequest[T any](c *gin.Context) (*T, error) {
	reqInput := new(T)

	if c.Request.Method != http.MethodGet {
		if err := c.ShouldBindJSON(reqInput); err != nil && !errors.Is(err, io.EOF) {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.InputFieldError(err))
			return nil, err
		}
	}

	if err := c.ShouldBindUri(reqInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.InputFieldError(err))
		return nil, err
	}
	if err := c.ShouldBindQuery(reqInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.InputFieldError(err))
		return nil, err
	}
	if err := c.ShouldBindHeader(reqInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.InputFieldError(err))
		return nil, err
	}
	if err := InputValidator.Struct(reqInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.InputFieldError(err))
		return nil, err
	}
	return reqInput, nil
}

// BindInputFromRequestWithAuth binds and validates request input and extracts
// the authenticated user ID from the JWT claims.
//
// Parameters:
//   - c: the Gin context containing the incoming HTTP request and JWT claims.
//
// Returns:
//   - The populated and validated input, the authenticated user ID,
//     or an error if request binding, validation, or authentication fails.
func BindInputFromRequestWithAuth[T any](c *gin.Context) (*T, string, error) {
	input, err := BindInputFromRequest[T](c)
	if err != nil {
		return nil, "", err
	}

	uid, err := GetUserIDFromRequest(c)
	if err != nil {
		return nil, "", err
	}
	return input, uid, nil
}
