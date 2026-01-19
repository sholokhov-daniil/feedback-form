package middleware

import (
	"context"
	"log/slog" // Современный стандарт логирования
	"net/http"
	"strings"

	"github.com/sholokhov-daniil/feedback-form/internal/models"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
)

type contextKey string
const userContextKey contextKey = "user_auth"

// Хелпер для получения пользователя из контекста (типобезопасно)
func GetUser(ctx context.Context) (*models.UserAuth, bool) {
	u, ok := ctx.Value(userContextKey).(*models.UserAuth)
	return u, ok
}

func AuthBearerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getToken(r)

		if token == "" {
			unauthorized(w)
			return
		}

        ua, err := repository.GetByToken(token)
        if err != nil {
            slog.Error("auth: repository error", "err", err)
            unauthorized(w)
            return
        }

        if ua == nil {
            slog.Error("auth: token not found " + token)
			unauthorized(w)
            return
        }

        ctx := context.WithValue(r.Context(), userContextKey, ua)
        next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unauthorized(w http.ResponseWriter) {
	http.Error(w, response.GetUnauthorizedResponse().ToJson(), http.StatusUnauthorized)
}

func getToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}

	return parts[1]
}