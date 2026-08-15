package ratelimitutils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate mockery --name Repository --filename common.go --outpkg mockRateLimit
type Repository interface {
	IncreaseRateLimit(ctx context.Context, key string, exp time.Duration) error
	GetCurrentRateLimit(ctx context.Context, key string) (int, error)
}

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) Repository {
	return &redisRepo{client: client}
}
