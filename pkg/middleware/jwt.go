package middleware

import (
	"net/http"
	"strings"

	"github.com/HemlockPham7/common-libs/pkg/jwtutils"
	"github.com/gin-gonic/gin"
)

type JWTAuth interface {
	JWTAuth() gin.HandlerFunc
}

type jwtAuth struct {
	jwtVal jwtutils.JWTValidator
}

// NewJWTAuth creates a JWT authentication middleware using the provided JWT validator.
//
// Parameters:
//   - jwtVal: the JWT validator used to validate incoming Bearer tokens.
//
// Returns:
//   - A JWTAuth middleware configured with the provided validator.
func NewJWTAuth(jwtVal jwtutils.JWTValidator) JWTAuth {
	return &jwtAuth{
		jwtVal: jwtVal,
	}
}

// JWTAuth returns a Gin middleware that authenticates requests using a Bearer token.
//
// The middleware extracts the JWT from the Authorization header, validates the token,
// and stores the JWT claims in the Gin context under the "claims" key.
// Requests with a missing, malformed, or invalid token are rejected with HTTP 401.
//
// Returns:
//   - A Gin handler function for JWT authentication.
func (j *jwtAuth) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenString := parts[1]

		// validate token
		tokenClaims, err := j.jwtVal.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// set claims to context
		c.Set("claims", tokenClaims)

		c.Next()
	}
}
