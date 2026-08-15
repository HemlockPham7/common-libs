package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/HemlockPham7/common-libs/pkg/jwtutils/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuth_JWTAuth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest       func(ctx *gin.Context)
		setupMockValidator func() *mocks.JWTValidator

		expectedCode     int
		expectedResponse string
		expectedAbort    bool
		checkClaims      bool
	}{
		{
			name: "valid token - cliams set in context",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Request.Header.Set("Authorization", "Bearer valid-token")
			},

			setupMockValidator: func() *mocks.JWTValidator {
				mockValidator := mocks.NewJWTValidator(t)
				mockValidator.On("ValidateJWT", "valid-token").
					Return(jwt.MapClaims{"sub": "de305d54-75b4-431b-adb2-eb6b9e546099"}, nil)
				return mockValidator
			},

			expectedCode:  http.StatusOK,
			expectedAbort: false,
			checkClaims:   true,
		},
		{
			name: "missing authorization header",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
			},

			setupMockValidator: func() *mocks.JWTValidator {
				return mocks.NewJWTValidator(t)
			},

			expectedCode:     http.StatusUnauthorized,
			expectedAbort:    true,
			expectedResponse: `{"error":"Authorization header missing"}`,
		},
		{
			name: "invalid format - no space separator",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Request.Header.Set("Authorization", "BearerTokenWIthNoSpace")
			},

			setupMockValidator: func() *mocks.JWTValidator {
				return mocks.NewJWTValidator(t)
			},

			expectedCode:     http.StatusUnauthorized,
			expectedAbort:    true,
			expectedResponse: `{"error":"Invalid Authorization header format"}`,
		},
		{
			name: "invalid token - validator returns error",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Request.Header.Set("Authorization", "Bearer invalid-token")
			},

			setupMockValidator: func() *mocks.JWTValidator {
				mockValidator := mocks.NewJWTValidator(t)
				mockValidator.On("ValidateJWT", "invalid-token").Return(nil, assert.AnError)
				return mockValidator
			},

			expectedCode:     http.StatusUnauthorized,
			expectedAbort:    true,
			expectedResponse: `{"error":"Invalid token"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)

			tc.setupRequest(ctx)
			jwtAuth := NewJWTAuth(tc.setupMockValidator())
			handler := jwtAuth.JWTAuth()
			handler(ctx)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedResponse, strings.TrimSpace(rec.Body.String()))
			assert.Equal(t, tc.expectedAbort, ctx.IsAborted())

			if tc.checkClaims {
				claims, exists := ctx.Get("claims")
				assert.True(t, exists)
				assert.IsType(t, jwt.MapClaims{}, claims)
			}
		})
	}
}
