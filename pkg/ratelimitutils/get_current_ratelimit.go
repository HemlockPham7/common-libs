package ratelimitutils

import "context"

// GetCurrentRateLimit returns the current request count for the given rate limit key.
//
// Parameters:
//   - ctx: the context used for the Redis operation.
//   - key: the Redis key used to track the rate limit.
//
// Returns:
//   - The current request count, or an error if the key cannot be retrieved or parsed.
func (r *redisRepo) GetCurrentRateLimit(ctx context.Context, key string) (int, error) {
	return r.client.Get(ctx, key).Int()
}
