package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"go-microservice-api/internal/contextkeys"
)

func randomID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = randomID()
		}
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), contextkeys.ReqID, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
