package domain

import (
	"data-processor/pkg/data-processor/proto/fileservice"
	"github.com/google/uuid"
)

const (
	statusReady = "ready"
)

type (
	UploadFileRequest struct {
		UserId      string
		FileName    string
		MimeType    string
		FileData    []byte
		Description string
	}

	UploadFileResponse struct {
		Message string
	}

	DownloadFileRequest struct {
		UserID string
		FileID string
	}
	DownloadFileResponse struct {
		FileName string
		MimeType string
		FileData []byte
	}

	ListUserFilesRequest struct {
		UserId string
	}

	ListUserFilesResponse struct {
		FilesMetadata []FileMetadata
	}

	FileMetadata struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		FileName    string
		Description string
	}
)

func (f *FileMetadata) ToProto() *fileservice.FileMetadata {
	return &fileservice.FileMetadata{
		FileId:      f.ID.String(),
		FileName:    f.FileName,
		Description: f.Description,
		Status:      statusReady,
	}
}
