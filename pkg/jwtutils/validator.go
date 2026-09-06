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

// NewJWTValidator creates a JWT validator using an RSA public key loaded from the given file path.
//
// Parameters:
//   - publicKeyPath: the path to the RSA public key file.
//
// Returns:
//   - A JWTValidator initialized with the public key, or an error if the key cannot be read or parsed.
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

// ValidateJWT validates a JWT using the configured RSA public key and extracts its claims.
//
// Parameters:
//   - tokenStr: the JWT string to validate.
//
// Returns:
//   - The JWT claims if the token is valid, or an error if the token is invalid
//     or its claims cannot be extracted.
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
