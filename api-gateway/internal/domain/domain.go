package domain

import (
	pb "github.com/Ddarli/svc/gateway/pkg/proto"
)

type LoginRequest struct {
	Username string
	Email    string
	Password string
}

func (r *LoginRequest) ToProto() *pb.LoginRequest {
	return &pb.LoginRequest{
		Username: r.Username,
		Password: r.Password,
	}
}

type ValidateTokenRequest struct {
	Token string
}

func (r *ValidateTokenRequest) ToProto() *pb.ValidateTokenRequest {
	return &pb.ValidateTokenRequest{
		Token: r.Token,
	}
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
	Phone    string
}

func (r *RegisterRequest) ToProto() *pb.RegisterRequest {
	return &pb.RegisterRequest{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Phone:    r.Phone,
	}
}

type (
	LoginResponse struct {
		Token string
	}
	ValidateTokenResponse struct {
		Valid  bool
		UserID string
	}

	RegisterResponse struct {
		Token string
	}
)

func ProtoLoginResponseToModel(resp *pb.LoginResponse) *LoginResponse {
	return &LoginResponse{
		Token: resp.GetToken(),
	}
}

func ProtoValidateTokenResponseToModel(resp *pb.ValidateTokenResponse) *ValidateTokenResponse {
	return &ValidateTokenResponse{
		Valid:  resp.GetValid(),
		UserID: resp.GetUserId(),
	}
}

func ProtoRegisterRequestToModel(resp *pb.RegisterResponse) *RegisterResponse {
	return &RegisterResponse{
		Token: resp.GetToken(),
	}
}
