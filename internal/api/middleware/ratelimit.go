package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HemlockPham7/common-libs/internal/app/repository/ratelimit"
	"github.com/HemlockPham7/common-libs/pkg/requestutils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RateLimit interface {
	RateLimit() gin.HandlerFunc
}

type rateLimit struct {
	repository ratelimit.Repository
}

func NewRateLimit(repository ratelimit.Repository) RateLimit {
	return &rateLimit{repository: repository}
}

const (
	RateLimitInterval  = 1 * time.Minute // sliding window ton tai trong bao lau, moi mot phut check bao nhieu request
	RateLimitCount     = 10              // so max request thuc hien trong 1 phut
	RateLimitKeyFormat = "rate_limit:%s"
)

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
			return
		}

		// increase rate limit
		if err := r.repository.IncreaseRateLimit(c, rateLimitKey, RateLimitInterval); err != nil {
			log.Error().Err(err).Msg("failed to increase rate limit")
		}
		c.Next()
	}
}
