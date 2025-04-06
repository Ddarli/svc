package service

import (
	"auth/internal/model"
	"context"
)

type (
	Repository interface {
		RegisterUser(ctx context.Context, user model.User) error
		GetUserByUsername(ctx context.Context, username string) (model.User, error)
	}
)
