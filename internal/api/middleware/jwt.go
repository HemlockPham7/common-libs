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

func NewJWTAuth(jwtVal jwtutils.JWTValidator) JWTAuth {
	return &jwtAuth{
		jwtVal: jwtVal,
	}
}

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
