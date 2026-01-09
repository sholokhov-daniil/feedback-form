package middleware

import (
	"context"
	"log/slog" // Современный стандарт логирования
	"net/http"
	"strings"
	"github.com/jmoiron/sqlx"
	"github.com/sholokhov-daniil/feedback-form/internal/models"
)

type contextKey string
const userContextKey contextKey = "user_auth"

// Хелпер для получения пользователя из контекста (типобезопасно)
func GetUser(ctx context.Context) (*models.UserAuth, bool) {
	u, ok := ctx.Value(userContextKey).(*models.UserAuth)
	return u, ok
}

func Auth(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				slog.Warn("auth: missing or invalid header", "remote_addr", r.RemoteAddr)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			
			// Проверка токена с учетом времени жизни прямо в SQL
			var ua models.UserAuth
			query := `
				SELECT id, user_id, active, expires_at 
				FROM user_auth 
				WHERE secret_hash = $1 AND active = true AND (expires_at IS NULL OR expires_at > NOW())
				LIMIT 1`
			
			if err := db.Get(&ua, query, token); err != nil {
				slog.Error("auth: failed to get user", "err", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, &ua)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}