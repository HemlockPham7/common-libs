package redis

import "github.com/redis/go-redis/v9"

// NewClient creates a Redis client using configuration loaded from environment variables.
//
// Parameters:
//   - envPrefix: the prefix used to load Redis configuration values.
//
// Returns:
//   - A configured Redis client, or an error if the configuration cannot be processed.
func NewClient(envPrefix string) (*redis.Client, error) {
	cfg, err := newConfig(envPrefix)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb, nil
}
