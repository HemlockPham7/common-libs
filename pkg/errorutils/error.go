package errorutils

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExtractToken = errors.New("failed to extract token")
)
