package middleware

import (
	"net/http"
	"strings"

	"github.com/daiki-kim/chat-app/pkg/auth"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"go.uber.org/zap"
)

func JwtAuthForHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			logger.Error("no token found")
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			logger.Error("invalid token format")
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenStr := tokenParts[1]
		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			logger.Error("failed to parse token", zap.Error(err))
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		r = r.WithContext(auth.SetUserIDToContext(r.Context(), claims.UserID))
		next.ServeHTTP(w, r)
	})
}
