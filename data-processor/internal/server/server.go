package server

import (
	"context"
	"data-processor/internal/domain"
	"data-processor/pkg/data-processor/proto/fileservice"
	"github.com/rotisserie/eris"
)

type Server struct {
	fileservice.UnimplementedFileServiceServer
	service service
}

func New(service service) *Server {
	return &Server{service: service}
}

func (s *Server) UploadUserFile(ctx context.Context, req *fileservice.UploadUserFileRequest) (*fileservice.UploadUserFileResponse, error) {
	resp := fileservice.UploadUserFileResponse{}

	res, err := s.service.UploadMedicalData(ctx, domain.UploadFileRequest{
		UserId:      req.UserId,
		FileName:    req.FileName,
		MimeType:    req.MimeType,
		FileData:    req.FileData,
		Description: req.Description,
	})
	if err != nil {
		return &resp, eris.Wrapf(err, "error uploading file")
	}

	resp.Message = res.Message

	return &resp, nil
}

func (s *Server) DownloadUserFile(ctx context.Context, req *fileservice.DownloadUserFileRequest) (*fileservice.DownloadUserFileResponse, error) {
	resp := fileservice.DownloadUserFileResponse{}
	res, err := s.service.DownloadMedicalData(ctx, domain.DownloadFileRequest{
		UserID: req.UserId,
		FileID: req.FileId,
	})
	if err != nil {
		return &resp, eris.Wrapf(err, "error downloading file")
	}

	resp.FileData = res.FileData
	resp.MimeType = res.MimeType
	resp.FileName = res.FileName

	return &resp, nil
}

func (s *Server) ListUserFiles(ctx context.Context, req *fileservice.ListUserFilesRequest) (*fileservice.ListUserFilesResponse, error) {
	resp := fileservice.ListUserFilesResponse{}
	res, err := s.service.ListUserFiles(ctx, domain.ListUserFilesRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &resp, eris.Wrapf(err, "error listing files")
	}

	var files []*fileservice.FileMetadata
	for _, file := range res.FilesMetadata {
		protoData := file.ToProto()
		files = append(files, protoData)
	}

	resp.Files = files

	return &resp, nil
}
