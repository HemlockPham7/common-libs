package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HemlockPham7/common-libs/pkg/ratelimitutils"
	"github.com/HemlockPham7/common-libs/pkg/requestutils"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

type RateLimit interface {
	RateLimit() gin.HandlerFunc
}

type rateLimit struct {
	repository ratelimitutils.Repository
}

// NewRateLimit creates a rate limiting middleware using the provided repository.
//
// Parameters:
//   - repository: the rate limit repository used to track request counts.
//
// Returns:
//   - A RateLimit middleware configured with the provided repository.
func NewRateLimit(repository ratelimitutils.Repository) RateLimit {
	return &rateLimit{repository: repository}
}

const (
	// RateLimitInterval defines the duration of the rate limiting window.
	RateLimitInterval = 1 * time.Minute // sliding window ton tai trong bao lau, moi mot phut check bao nhieu request
	// RateLimitCount defines the maximum number of requests allowed within the rate limiting window.
	RateLimitCount = 1000 // so max request thuc hien trong 1 phut
	// RateLimitKeyFormat defines the Redis key format used to track request rates.
	RateLimitKeyFormat = "rate_limit:%s"
)

// RateLimit returns a Gin middleware that limits requests per user.
//
// The middleware uses the user ID from the request to track the number of requests
// within the configured rate limiting window. Requests exceeding the limit are
// rejected with HTTP 429 Too Many Requests.
func (r *rateLimit) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user-id from request
		uid, err := requestutils.GetUserIDFromRequest(c)
		if err != nil {
			return
		}

		// create rate limit key
		rateLimitKey := fmt.Sprintf(RateLimitKeyFormat, uid)

		// get current rate limit
		currentRate, err := r.repository.GetCurrentRateLimit(c, rateLimitKey)
		if err != nil {
			log.Error().Err(err).Msg("failed to get current rate limit")
		}

		// check if rate limit exceeded
		if currentRate >= RateLimitCount {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			nrTransaction := newrelic.FromContext(c)
			nrTransaction.Application().RecordCustomEvent("RateLimit", map[string]interface{}{
				"client":   uid,
				"cur_rate": currentRate,
				"endpoint": c.Request.URL.Path,
			})
			return
		}

		// increase rate limit
		if err := r.repository.IncreaseRateLimit(c, rateLimitKey, RateLimitInterval); err != nil {
			log.Error().Err(err).Msg("failed to increase rate limit")
		}
		c.Next()
	}
}
