package service

import (
	"auth/internal/model"
	"auth/pkg/transport/medicalpb"
	"context"
)

type (
	Repository interface {
		RegisterUser(ctx context.Context, user model.User) error
		GetUserByUsername(ctx context.Context, username string) (model.User, error)
	}

	BlockchainProcessor interface {
		GenerateNewAccount(ctx context.Context, req *medicalpb.Empty) (response *medicalpb.AccountResponse, err error)
	}
)
