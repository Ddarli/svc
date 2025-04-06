package grpc

import (
	"auth/internal/model"
	"auth/internal/service"
	pb "auth/pkg/proto"
	"context"
	"github.com/google/uuid"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	service *service.Service
}

func NewAuthServer(service *service.Service) *AuthServer {
	return &AuthServer{
		service: service,
	}
}

func (s *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.service.Authorize(ctx, model.AuthRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (s *AuthServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	token, err := s.service.Register(ctx, model.RegisterRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{Token: token}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (resp *pb.ValidateTokenResponse, err error) {
	resp = &pb.ValidateTokenResponse{}

	id, err := s.service.ValidateToken(ctx, request.Token)
	if err != nil {
		return resp, err
	}

	if id != uuid.Nil {
		resp.Valid = true
		resp.UserId = id.String()
	} else {
		resp.Valid = false
		resp.UserId = uuid.Nil.String()
	}

	return resp, err
}
