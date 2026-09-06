package requestutils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrEmptyID      = errors.New("empty id in token claims")
	ErrNotExist     = errors.New("claim not exist")
)

var PrivateKeyCompromiseWarning = newrelic.Error{
	Message: "JWT has invalid format - Private key might be compromised!",
	Class:   "JWTError",
}

// GetUserIDFromRequest extracts the authenticated user ID from JWT claims stored in the Gin context.
//
// The user ID is read from the "sub" claim. The request is aborted with HTTP 401
// if the claims are missing, invalid, or the user ID is empty.
//
// Parameters:
//   - c: the Gin context containing the authenticated user's JWT claims.
//
// Returns:
//   - The authenticated user ID, or an error if the claims or user ID are invalid.
func GetUserIDFromRequest(c *gin.Context) (string, error) {
	claims, exist := c.Get("claims")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "claim not exist"})
		c.Abort()
		return "", ErrNotExist
	}

	tokenInfo, ok := claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		c.Abort()
		return "", ErrInvalidToken
	}

	uid, ok := tokenInfo["sub"].(string)
	if !ok {
		newrelic.FromContext(c).NoticeError(PrivateKeyCompromiseWarning)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		c.Abort()
		return "", ErrInvalidToken
	} else if uid == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "empty id in token claims"})
		c.Abort()
		return "", ErrEmptyID
	}

	return uid, nil
}
