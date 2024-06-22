package middleware

import (
	"net/http"

	"github.com/daiki-kim/chat-app/pkg/auth"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"go.uber.org/zap"
)

func JwtAuthForWS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.URL.Query().Get("token")
		if tokenStr == "" {
			logger.Error("no token found")
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

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
