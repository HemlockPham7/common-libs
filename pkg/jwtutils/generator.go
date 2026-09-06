package jwtutils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockery --name JWTGenerator --filename jwt_generator.go --outpkg mocks
type JWTGenerator interface {
	GenerateJWT(jwtContent jwt.MapClaims) (string, error)
}

type jwtGenerator struct {
	privateKey *rsa.PrivateKey
}

// NewJWTGenerator creates a JWT generator using an RSA private key loaded from the given file path.
//
// Parameters:
//   - privateKeyPath: the path to the RSA private key file.
//
// Returns:
//   - A JWTGenerator initialized with the private key, or an error if the key cannot be read or parsed.
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

// GenerateJWT generates a signed JWT using the configured RSA private key.
//
// Parameters:
//   - jwtContent: the claims to include in the JWT.
//
// Returns:
//   - The signed JWT string, or an error if the token cannot be signed.
func (g *jwtGenerator) GenerateJWT(jwtContent jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtContent)

	tokenString, err := token.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
