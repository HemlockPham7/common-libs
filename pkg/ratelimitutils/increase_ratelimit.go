package ratelimitutils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// IncreaseRateLimit increments the rate limit counter for the given key
// and sets its expiration duration.
//
// Parameters:
//   - ctx: the context used for the Redis operations.
//   - key: the Redis key used to track the rate limit.
//   - exp: the expiration duration for the rate limit counter.
//
// Returns:
//   - An error if the Redis operations fail.
func (r *redisRepo) IncreaseRateLimit(ctx context.Context, key string, exp time.Duration) error {
	_, err := r.client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, exp)
		return nil
	})

	return err
}
