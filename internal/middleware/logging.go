package middleware

import (
	"net/http"
	"time"

	"go-microservice-api/internal/logger"
	"go-microservice-api/internal/metrics"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		metrics.IncRequest()
		next.ServeHTTP(w, r)
		logger.Infof("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
