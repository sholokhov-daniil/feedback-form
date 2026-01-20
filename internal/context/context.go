package context

import (
	"context"
	"errors"

	"github.com/sholokhov-daniil/feedback-form/internal/models"
)

const userContextKey string = "user"
const authTypeContextKey string = "user_auth"

func SetUser(ctx context.Context, u *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, u)
}

func GetUser(ctx context.Context) (*models.User, error) {
	v := ctx.Value(userContextKey)

	if v == nil {
		return nil, errors.New("user not found in context")
	}

	u, ok := v.(*models.User)

	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return u, nil
}

func SetUserAuth(ctx context.Context, ua *models.UserAuth) context.Context {
	return context.WithValue(ctx, authTypeContextKey, ua)
}

func GetUserAuth(ctx context.Context) (*models.UserAuth, error) {
	v := ctx.Value(authTypeContextKey)

	if v == nil {
		return nil, errors.New("user auth not found in context")
	}

	ua, ok := v.(*models.UserAuth)

	if !ok {
		return nil, errors.New("invalid user auth type in context")
	}

	return ua, nil
}