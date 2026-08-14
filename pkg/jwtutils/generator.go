package jwtutils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWTGenerator interface {
	GenerateJWT(jwtContent jwt.MapClaims) (string, error)
}

type jwtGenerator struct {
	privateKey *rsa.PrivateKey
}

func NewJWTGenerator(privateKeyPath string) (JWTGenerator, error) {
	privateKeyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, err
	}

	return &jwtGenerator{
		privateKey: privateKey,
	}, nil
}

func (g *jwtGenerator) GenerateJWT(jwtContent jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtContent)

	tokenString, err := token.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
