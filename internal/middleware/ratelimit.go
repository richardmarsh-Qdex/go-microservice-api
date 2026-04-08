package middleware

import (
	"net/http"
	"sync"
	"time"
)

type tokenBucket struct {
	tokens float64
	last   time.Time
	rate   float64
	cap    float64
	mu     sync.Mutex
}

func (b *tokenBucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(b.last).Seconds()
	b.last = now
	b.tokens += elapsed * b.rate
	if b.tokens > b.cap {
		b.tokens = b.cap
	}
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

var defaultBucket = &tokenBucket{tokens: 10, last: time.Now(), rate: 5, cap: 10}

func SimpleRateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !defaultBucket.allow() {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
