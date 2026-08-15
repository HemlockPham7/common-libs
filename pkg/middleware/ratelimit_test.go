package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/HemlockPham7/common-libs/internal/app/repository/ratelimit/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestRateLimit_RateLimit(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest             func(ctx *gin.Context)
		setupRateLimitRepository func(ctx context.Context) *mockRateLimit.Repository

		expectedCode     int
		expectedResponse string
	}{
		{
			name: "counter not exists",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Set("claims", jwt.MapClaims{
					"sub": "de305d54-75b4-431b-adb2-eb6b9e546099",
				})
			},

			setupRateLimitRepository: func(ctx context.Context) *mockRateLimit.Repository {
				mockRL := mockRateLimit.NewRepository(t)
				mockRL.On("GetCurrentRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099").Return(0, nil)
				mockRL.On("IncreaseRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099", 1*time.Minute).Return(nil)
				return mockRL
			},

			expectedCode:     http.StatusOK,
			expectedResponse: ``,
		},
		{
			name: "claim not exists",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
			},

			setupRateLimitRepository: func(ctx context.Context) *mockRateLimit.Repository {
				return mockRateLimit.NewRepository(t)
			},

			expectedCode:     http.StatusUnauthorized,
			expectedResponse: `{"message":"claim not exist"}`,
		},
		{
			name: "fail due to get current rate limit error",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Set("claims", jwt.MapClaims{
					"sub": "de305d54-75b4-431b-adb2-eb6b9e546099",
				})
			},

			setupRateLimitRepository: func(ctx context.Context) *mockRateLimit.Repository {
				mockRL := mockRateLimit.NewRepository(t)
				mockRL.On("GetCurrentRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099").Return(-1, assert.AnError)
				mockRL.On("IncreaseRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099", 1*time.Minute).Return(nil)
				return mockRL
			},

			expectedCode:     http.StatusOK,
			expectedResponse: ``,
		},
		{
			name: "too many request error",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Set("claims", jwt.MapClaims{
					"sub": "de305d54-75b4-431b-adb2-eb6b9e546099",
				})
			},

			setupRateLimitRepository: func(ctx context.Context) *mockRateLimit.Repository {
				mockRL := mockRateLimit.NewRepository(t)
				mockRL.On("GetCurrentRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099").Return(RateLimitCount, nil)
				return mockRL
			},

			expectedCode:     http.StatusTooManyRequests,
			expectedResponse: `{"error":"rate limit exceeded"}`,
		},
		{
			name: "failed to increase rate limit",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/self/info", nil)
				ctx.Set("claims", jwt.MapClaims{"sub": "de305d54-75b4-431b-adb2-eb6b9e546099"})
			},

			setupRateLimitRepository: func(ctx context.Context) *mockRateLimit.Repository {
				mockRL := mockRateLimit.NewRepository(t)
				mockRL.On("GetCurrentRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099").Return(5, nil)
				mockRL.On("IncreaseRateLimit", ctx, "rate_limit:de305d54-75b4-431b-adb2-eb6b9e546099", 1*time.Minute).Return(assert.AnError)
				return mockRL
			},

			expectedCode:     http.StatusOK,
			expectedResponse: ``,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)

			tc.setupRequest(ctx)
			mockRL := tc.setupRateLimitRepository(ctx)

			rateLimitMiddleware := NewRateLimit(mockRL)
			rateLimitMiddleware.RateLimit()(ctx)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedResponse, strings.TrimSpace(rec.Body.String()))
		})
	}
}
