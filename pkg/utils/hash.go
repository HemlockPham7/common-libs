package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name Hasher --filename hash.go
type Hasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}

type hasher struct {
}

// NewHasher creates a password hashing service using bcrypt.
//
// Returns:
//   - A Hasher service for hashing and comparing passwords.
func NewHasher() Hasher {
	return &hasher{}
}

var ErrHashFailed = errors.New("failed to hash password")

// Hash securely hashes a password using bcrypt.
//
// Parameters:
//   - password: the plain-text password to hash.
//
// Returns:
//   - The bcrypt password hash, or an error if hashing fails.
func (h *hasher) Hash(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHashFailed
	}
	return string(hashBytes), nil
}

// Compare compares a plain-text password against a bcrypt hash.
//
// Parameters:
//   - hash: the bcrypt hash to compare against.
//   - password: the plain-text password to verify.
//
// Returns:
//   - true if the password matches the hash; otherwise, false.
func (h *hasher) Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
