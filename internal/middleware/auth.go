package middleware

import (
	"context"
	"net/http"
	"strings"

	"go-microservice-api/internal/auth"
	"go-microservice-api/internal/contextkeys"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseClaimsFromRequest(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		uid, _ := claims["user_id"].(string)
		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)

		ctx := context.WithValue(r.Context(), contextkeys.UserID, uid)
		ctx = context.WithValue(ctx, contextkeys.Email, email)
		ctx = context.WithValue(ctx, contextkeys.Role, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
