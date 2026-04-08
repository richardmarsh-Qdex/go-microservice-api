package middleware

import (
	"net/http"
	"runtime/debug"

	"go-microservice-api/internal/logger"
	"go-microservice-api/internal/metrics"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				metrics.IncError()
				logger.Errorf("panic: %v\n%s", rec, debug.Stack())
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
