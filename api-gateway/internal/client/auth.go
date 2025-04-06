package client

import (
	"context"
	"github.com/Ddarli/svc/gateway/internal/domain"
	pb "github.com/Ddarli/svc/gateway/pkg/proto"
	"google.golang.org/grpc"
	"log/slog"
)

type AuthClient struct {
	client pb.AuthServiceClient
}

func New(ctx context.Context, cfg *Config) *AuthClient {
	client, err := createClient(ctx, cfg.AuthClient.Port)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return &AuthClient{client: client}
}

func (a *AuthClient) Login(ctx context.Context, req *pb.LoginRequest) (*domain.LoginResponse, error) {
	resp, err := a.client.Login(ctx, req)
	if err != nil {
		slog.Error(err.Error())

		return nil, err
	}

	res := domain.ProtoLoginResponseToModel(resp)

	return res, nil
}

func (a *AuthClient) Validate(ctx context.Context, req *pb.ValidateTokenRequest) (*domain.ValidateTokenResponse, error) {
	protoResponse, err := a.client.ValidateToken(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())

		return nil, err
	}

	response := domain.ProtoValidateTokenResponseToModel(protoResponse)

	return response, nil
}

func (a *AuthClient) Register(ctx context.Context, req *pb.RegisterRequest) (*domain.RegisterResponse, error) {
	protoResponse, err := a.client.Register(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())

		return nil, err
	}

	response := domain.ProtoRegisterRequestToModel(protoResponse)

	return response, nil
}

func createClient(ctx context.Context, port string) (pb.AuthServiceClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewAuthServiceClient(conn)

	slog.InfoContext(ctx, "run auth client on", "port", port)

	return client, nil
}
