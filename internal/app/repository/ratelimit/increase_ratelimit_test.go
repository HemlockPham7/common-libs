package ratelimit

import (
	"context"
	"testing"
	"time"

	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepo_IncreaseRateLimit(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context, t *testing.T) *redis.Client

		expectedError error
		inputKey      string
		inputExp      time.Duration

		verifyFunc func(ctx context.Context, redisClient *redis.Client, inputKey string)
	}{
		{
			name: "success",
			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClientMock := redisPkg.InitMockRedis(t)
				return redisClientMock
			},
			expectedError: nil,
			inputKey:      "test_key",
			inputExp:      time.Hour,
			verifyFunc: func(ctx context.Context, redisClient *redis.Client, inputKey string) {
				currentRateLimit, err := redisClient.Get(ctx, inputKey).Result()
				assert.NoError(t, err)
				assert.Equal(t, "1", currentRateLimit)
			},
		},
		{
			name: "failed due to redis client closed",
			setupMock: func(ctx context.Context, t *testing.T) *redis.Client {
				redisClientMock := redisPkg.InitMockRedis(t)
				_ = redisClientMock.Close()
				return redisClientMock
			},
			expectedError: redis.ErrClosed,
			inputKey:      "test_key",
			inputExp:      time.Hour,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisClientMock := tc.setupMock(ctx, t)
			storage := NewRedisRepo(redisClientMock)

			err := storage.IncreaseRateLimit(ctx, tc.inputKey, tc.inputExp)

			assert.ErrorIs(t, err, tc.expectedError)
			if err == nil {
				tc.verifyFunc(ctx, redisClientMock, tc.inputKey)
			}
		})
	}
}
