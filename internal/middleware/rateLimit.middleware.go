package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/global"
	limiter "github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
	"net/http"
	"time"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimit := &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"),
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("50-S"),
	}
	return rateLimit
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.Redis, limiter.StoreOptions{
		Prefix:          "rate-limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})
	if err != nil {
		return nil
	}

	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}

	instance := limiter.New(store, rate)
	return instance
}

// GLOBAL RATE LIMITER
func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"
		log.Println("global--->")
		limiterContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("Failed to check rate limit Global ", err)
			c.Next()
			return
		}
		if limiterContext.Reached {
			log.Printf("Rate limit breached GLOBAL %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached GLOBAL, try later"})
			return
		}
		c.Next()
	}
}

// PUBLIC API RATE LIMITER
func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)
		if rateLimitPath == rl.publicAPIRateLimiter {
			log.Println("Client IP ----> ", c.ClientIP())

			key := fmt.Sprintf("%s-%s", "111-222-333-44", urlPath)
			limiterContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit Public ", err)
				c.Next()
				return
			}
			if limiterContext.Reached {
				log.Printf("Rate limit breached Public %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached Public, try later"})
				return
			}
		}
		c.Next()
	}
}

// USER PRIVATE API RATE LIMITER
func (rl *RateLimiter) UserPrivateAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath == rl.userPrivateAPIRateLimiter {
			userId := 1001
			key := fmt.Sprintf("%s-%s", userId, urlPath)
			limiterContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit User Private ", err)
				c.Next()
				return
			}
			if limiterContext.Reached {
				log.Printf("Rate limit breached User Private %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached User Private, try later"})
				return
			}
		}
		c.Next()
	}
}

func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "/v1/2024/user/login" || urlPath == "/ping/80" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/v1/2024/info" || urlPath == "/ping/50" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
