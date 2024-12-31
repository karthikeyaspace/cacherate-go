package middleware

// ratelimiting methods - for all users, per client ratelimit

import (
	"net/http"
	"sync"
	"time"

	"github.com/karthikeyaspace/proxy-go/internal/helpers"
	"golang.org/x/time/rate"
)

type Ratelimiter struct {
	limiter       *rate.Limiter
	clientLimiter map[string]*rate.Limiter
	mutex         sync.Mutex
	rate          rate.Limit
	burst         int
}

func NewRateLimiter(r rate.Limit, burst int) *Ratelimiter {
	return &Ratelimiter{
		limiter:       rate.NewLimiter(r, burst), // burst reqs per r time
		clientLimiter: make(map[string]*rate.Limiter),
		rate:          r,
		burst:         burst,
	}
}

// rate limiter for specific client
func (r *Ratelimiter) getClientRatelimiter(ip string) *rate.Limiter {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	limiter, exists := r.clientLimiter[ip]
	if !exists {
		limiter = rate.NewLimiter(r.rate, r.burst)
		r.clientLimiter[ip] = limiter
	}

	return limiter
}

func (r *Ratelimiter) Ratelimit(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// rate limiter for all users
		if !r.limiter.Allow() {
			retryAfter := r.limiter.Reserve().Delay().Seconds()
			w.Header().Set("Retry-After", time.Duration(retryAfter).String())
			helpers.HandleResponse(w, http.StatusTooManyRequests, map[string]interface{}{
				"success":     false,
				"message":     "Too many requests(overall ratelimit)",
				"Retry After": retryAfter,
			})
			return
		}

		// rate limiter for specific client
		ip := req.RemoteAddr
		clientLimiter := r.getClientRatelimiter(ip)

		if !clientLimiter.Allow() {
			retryAfter := clientLimiter.Reserve().Delay().Seconds()
			w.Header().Set("Retry-After", time.Duration(retryAfter).String())
			helpers.HandleResponse(w, http.StatusTooManyRequests, map[string]interface{}{
				"success":    false,
				"message":    "Too many requests(client ratelimit)",
				"retryAfter": retryAfter,
			})
			return
		}

		next.ServeHTTP(w, req)
	}
}
