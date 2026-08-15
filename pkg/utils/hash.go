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

func NewHasher() Hasher {
	return &hasher{}
}

var ErrHashFailed = errors.New("failed to hash password")

func (h *hasher) Hash(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHashFailed
	}
	return string(hashBytes), nil
}

func (h *hasher) Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
