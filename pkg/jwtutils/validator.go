package jwtutils

import (
	"crypto/rsa"
	"os"

	"github.com/HemlockPham7/common-libs/pkg/errorutils"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockery --name JWTValidator --filename jwt_validator.go --outpkg mocks
type JWTValidator interface {
	ValidateJWT(tokenStr string) (jwt.MapClaims, error)
}

type jwtValidator struct {
	publicKey *rsa.PublicKey
}

func NewJWTValidator(publicKeyPath string) (JWTValidator, error) {
	publicKeyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return &jwtValidator{
		publicKey: publicKey,
	}, nil
}

func (v *jwtValidator) ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return v.publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errorutils.ErrInvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errorutils.ErrExtractToken
}
