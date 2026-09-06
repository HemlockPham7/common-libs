package redis

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// InitMockRedis creates an in-memory Redis client for testing.
//
// The returned client connects to a miniredis instance that is automatically
// started and cleaned up with the test.
//
// Parameters:
//   - t: the testing instance used to manage the mock Redis lifecycle.
//
// Returns:
//   - A Redis client connected to the in-memory Redis server.
func InitMockRedis(t *testing.T) *redis.Client {
	mock := miniredis.RunT(t)
	return redis.NewClient(&redis.Options{
		Addr: mock.Addr(),
	})
}
