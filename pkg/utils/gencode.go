package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//go:generate mockery --name GenCode --filename gencode.go
type GenCode interface {
	GenerateCode(length int) (string, error)
}

type genCodeService struct {
}

// NewGenCode creates a service for generating secure random alphanumeric codes.
//
// Returns:
//   - A GenCode service that generates cryptographically secure random codes.
func NewGenCode() GenCode {
	return &genCodeService{}
}

// GenerateCode generates a cryptographically secure random alphanumeric code.
//
// Parameters:
//   - length: the number of characters to generate.
//
// Returns:
//   - A random alphanumeric code of the requested length, or an error if random
//     data cannot be generated.
func (s *genCodeService) GenerateCode(length int) (string, error) {
	code := make([]byte, length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[n.Int64()]
	}
	return string(code), nil
}
