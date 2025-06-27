package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func RateLimitMiddleware(limit int, window time.Duration) echo.MiddlewareFunc {
	limiter := NewRateLimiter(limit, window)
	
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP := c.RealIP()
			
			if !limiter.Allow(clientIP) {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"error": "Rate limit exceeded. Please try again later.",
				})
			}
			
			return next(c)
		}
	}
}

func (rl *RateLimiter) Allow(clientIP string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	now := time.Now()
	windowStart := now.Add(-rl.window)
	
	requests, exists := rl.requests[clientIP]
	if !exists {
		requests = []time.Time{}
	}
	
	var validRequests []time.Time
	for _, reqTime := range requests {
		if reqTime.After(windowStart) {
			validRequests = append(validRequests, reqTime)
		}
	}
	
	if len(validRequests) >= rl.limit {
		return false
	}
	
	validRequests = append(validRequests, now)
	rl.requests[clientIP] = validRequests
	
	return true
}

func (rl *RateLimiter) Cleanup() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	now := time.Now()
	windowStart := now.Add(-rl.window)
	
	for clientIP, requests := range rl.requests {
		var validRequests []time.Time
		for _, reqTime := range requests {
			if reqTime.After(windowStart) {
				validRequests = append(validRequests, reqTime)
			}
		}
		
		if len(validRequests) == 0 {
			delete(rl.requests, clientIP)
		} else {
			rl.requests[clientIP] = validRequests
		}
	}
}
