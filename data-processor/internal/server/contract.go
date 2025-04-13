package server

import (
	"context"
	"data-processor/internal/domain"
)

type (
	service interface {
		UploadMedicalData(ctx context.Context, req domain.UploadFileRequest) (domain.UploadFileResponse, error)
		DownloadMedicalData(ctx context.Context, req domain.DownloadFileRequest) (domain.DownloadFileResponse, error)
		ListUserFiles(ctx context.Context, request domain.ListUserFilesRequest) (domain.ListUserFilesResponse, error)
	}
)
