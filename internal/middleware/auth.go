package middleware

import (
	"log/slog" // Современный стандарт логирования
	"net/http"
	"strings"

	"github.com/sholokhov-daniil/feedback-form/internal/context"
	"github.com/sholokhov-daniil/feedback-form/internal/repository"
	"github.com/sholokhov-daniil/feedback-form/internal/response"
)

const authTypeContextKey string = "user_auth"
const userContextKey string = "user"

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

		ctx := r.Context()

		u, err := repository.NewUserRepository().GetByID(ctx, ua.UserID)

		if err != nil {
			slog.Error("user: not found", "err", err)
			unauthorized(w)
			return
		}

		ctx = context.SetUserAuth(ctx, ua)
		ctx = context.SetUser(ctx, u)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unauthorized(w http.ResponseWriter) {
	http.Error(w, response.CreateUnauthorizedResponse().ToJson(), http.StatusUnauthorized)
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
