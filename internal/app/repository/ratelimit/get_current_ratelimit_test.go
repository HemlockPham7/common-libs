package ratelimit

import (
	"context"
	"testing"
	"time"

	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepo_GetCurrentRateLimit(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context, t *testing.T) *redis.Client

		inputKey string

		expectedError error
	}{
		{
			name: "success",

			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClientMock := redisPkg.InitMockRedis(t)
				redisClientMock.Set(ctx, "rate_limit:123", 1, time.Hour)
				return redisClientMock
			},

			inputKey:      "rate_limit:123",
			expectedError: nil,
		},
		{
			name: "failed due to redis client closed",

			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClientMock := redisPkg.InitMockRedis(t)
				_ = redisClientMock.Close()
				return redisClientMock
			},

			inputKey:      "rate_limit:123",
			expectedError: redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisClientMock := tc.setupMock(ctx, t)
			storage := NewRedisRepo(redisClientMock)

			currentRateLimit, err := storage.GetCurrentRateLimit(ctx, tc.inputKey)
			assert.ErrorIs(t, tc.expectedError, err)
			if err == nil {
				assert.NotNil(t, currentRateLimit)
			}
		})
	}
}
