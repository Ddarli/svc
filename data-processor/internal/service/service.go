package service

import (
	"bytes"
	"context"
	"crypto/sha256"
	"data-processor/internal/domain"
	"data-processor/internal/repository/entities"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
)

type Service struct {
	storage storage
	repo    repo
}

func New(storage storage, repo repo) *Service {
	return &Service{storage: storage, repo: repo}
}

func (s *Service) UploadMedicalData(ctx context.Context, req domain.UploadFileRequest) (domain.UploadFileResponse, error) {
	var res domain.UploadFileResponse

	filePath := fmt.Sprintf("patients/%s/%s", req.UserId, req.FileName)

	reader := bytes.NewReader(req.FileData)

	err := s.storage.PutData(filePath, reader)
	if err != nil {
		return res, fmt.Errorf("error saving file to the storage: %v", err)
	}

	medicalData := entities.MedicalData{
		ID:          uuid.New(),
		UserID:      uuid.MustParse(req.UserId),
		FileName:    req.FileName,
		MimeType:    req.MimeType,
		Description: req.Description,
		S3Key:       filePath,
		Hash:        calculateFileHash(req.FileData),
	}

	err = s.repo.InsertMedicalData(ctx, medicalData)
	if err != nil {
		return res, eris.Wrapf(err, "error saving file meta data to the storage")
	}

	res.Message = "File uploaded successfully"

	return res, nil
}

func (s *Service) DownloadMedicalData(ctx context.Context, req domain.DownloadFileRequest) (domain.DownloadFileResponse, error) {
	resp := domain.DownloadFileResponse{}

	metaData, err := s.repo.SelectByID(ctx, req.FileID)
	if err != nil {
		return resp, eris.Wrapf(err, "error fetching data from db")
	}

	resp.FileName = metaData.FileName
	resp.MimeType = metaData.MimeType

	file, err := s.storage.GetData(metaData.S3Key)
	if err != nil {
		return resp, eris.Wrapf(err, "error fetching data from storage")
	}

	resp.FileData = file

	return resp, nil
}

func (s *Service) ListUserFiles(ctx context.Context, request domain.ListUserFilesRequest) (domain.ListUserFilesResponse, error) {
	var res domain.ListUserFilesResponse

	files, err := s.repo.SelectMedicalDataByUserID(ctx, uuid.MustParse(request.UserId))
	if err != nil {
		return res, eris.Wrapf(err, "error fetching data from db")
	}

	var fileList []domain.FileMetadata
	for _, file := range files {
		fileList = append(fileList, domain.FileMetadata{
			ID:          file.ID,
			UserID:      file.UserID,
			FileName:    file.FileName,
			Description: file.Description,
		})
	}

	res.FilesMetadata = fileList

	return res, nil
}

func calculateFileHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
