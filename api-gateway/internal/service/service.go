package service

import (
	"context"
	"github.com/Ddarli/svc/gateway/internal/client"
	"github.com/Ddarli/svc/gateway/internal/domain"
	"log/slog"
)

type Service struct {
	authClient *client.AuthClient
}

func New(authClient *client.AuthClient) *Service {
	return &Service{authClient: authClient}
}

func (s *Service) Auth(ctx context.Context, req domain.LoginRequest) string {
	protoRequest := req.ToProto()

	resp, err := s.authClient.Login(ctx, protoRequest)
	if err != nil {
		slog.Error(err.Error())

		return ""
	}

	return resp.Token
}

func (s *Service) ValidateToken(ctx context.Context, req domain.ValidateTokenRequest) bool {
	protoRequest := req.ToProto()

	resp, err := s.authClient.Validate(ctx, protoRequest)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())

		return false
	}

	return resp.Valid
}

func (s *Service) Register(ctx context.Context, req domain.RegisterRequest) string {
	protoRequest := req.ToProto()

	resp, err := s.authClient.Register(ctx, protoRequest)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())

		return ""
	}

	return resp.Token
}
