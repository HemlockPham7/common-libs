package ratelimit

import "context"

func (r *redisRepo) GetCurrentRateLimit(ctx context.Context, key string) (int, error) {
	return r.client.Get(ctx, key).Int()
}
