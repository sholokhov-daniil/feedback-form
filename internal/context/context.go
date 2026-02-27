package context

import (
	"context"

	ex "github.com/sholokhov-daniil/feedback-form/internal/exceptions"
	"github.com/sholokhov-daniil/feedback-form/internal/models/user"
)

const userContextKey string = "user"
const authTypeContextKey string = "user_auth"

func SetUser(ctx context.Context, u *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, u)
}

func GetUser(ctx context.Context) (*models.User, error) {
	v := ctx.Value(userContextKey)

	if v == nil {
		return nil, ex.ErrorUserNotFound
	}

	u, ok := v.(*models.User)

	if !ok {
		return nil, ex.ErrorUserInvalidType
	}

	return u, nil
}

func SetUserAuth(ctx context.Context, ua *models.UserAuth) context.Context {
	return context.WithValue(ctx, authTypeContextKey, ua)
}

func GetUserAuth(ctx context.Context) (*models.UserAuth, error) {
	v := ctx.Value(authTypeContextKey)

	if v == nil {
		return nil, ex.ErrorUserAuthNotFound
	}

	ua, ok := v.(*models.UserAuth)

	if !ok {
		return nil, ex.ErrorUserInvalidType
	}

	return ua, nil
}
