package service

import (
	"context"
	"fmt"
	"github.com/Ddarli/svc/gateway/internal/client"
	"github.com/Ddarli/svc/gateway/internal/domain"
	pb "github.com/Ddarli/svc/gateway/pkg/proto"
	"log/slog"
)

type Service struct {
	authClient *client.AuthClient
	dataClient *client.DataClient
}

func New(authClient *client.AuthClient, dataClient *client.DataClient) *Service {
	return &Service{authClient: authClient, dataClient: dataClient}
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

func (s *Service) UploadFile(ctx context.Context, req domain.UploadFileRequest) (*domain.UploadFileResponse, error) {
	var resp domain.UploadFileResponse

	res, err := s.dataClient.UploadFile(ctx, &pb.UploadUserFileRequest{
		UserId:      req.UserID,
		FileName:    req.FileName,
		MimeType:    req.MimeType,
		FileData:    req.FileData,
		Description: req.Description,
	})
	if err != nil {
		slog.Error(err.Error())

		return nil, fmt.Errorf("error upload file grpc client: %v", err)
	}

	resp.Message = res.GetMessage()

	return &resp, nil
}

func (s *Service) DownloadFile(ctx context.Context, req domain.DownloadFileRequest) (*domain.DownloadFileResponse, error) {
	var resp domain.DownloadFileResponse

	res, err := s.dataClient.DownloadFile(ctx, &pb.DownloadUserFileRequest{
		UserId: req.UserID,
		FileId: req.FileID,
	})
	if err != nil {
		slog.Error(err.Error())

		return nil, fmt.Errorf("error download file grpc client: %v", err)
	}

	resp.FileName = res.GetFileName()
	resp.MimeType = res.GetMimeType()
	resp.FileData = res.GetFileData()

	return &resp, nil
}

func (s *Service) ListFile(ctx context.Context, req domain.ListUserFileRequest) (*domain.ListUserFilesResponse, error) {
	var resp domain.ListUserFilesResponse

	res, err := s.dataClient.ListUserFiles(ctx, &pb.ListUserFilesRequest{
		UserId: req.UserID,
	})
	if err != nil {
		slog.Error(err.Error())

		return nil, fmt.Errorf("error list user files grpc client: %v", err)
	}
	for _, file := range res.GetFiles() {
		resp.FilesMetadata = append(resp.FilesMetadata, domain.FileMetadata{
			ID:          file.GetFileId(),
			UserID:      req.UserID,
			FileName:    file.GetFileName(),
			Description: file.GetDescription(),
		})
	}

	return &resp, nil
}
