package ratelimitutils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *redisRepo) IncreaseRateLimit(ctx context.Context, key string, exp time.Duration) error {
	_, err := r.client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, exp)
		return nil
	})

	return err
}
