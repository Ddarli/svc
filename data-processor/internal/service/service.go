package service

import (
	"context"
	"crypto/sha256"
	"data-processor/internal/repository/entities"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"time"
)

type Service struct {
	storage storage
	repo    repo
}

func New(storage storage, repo repo) *Service {
	return &Service{storage: storage, repo: repo}
}

func (s *Service) AddNewMedicalData(ctx context.Context, file multipart.File) {
	defer file.Close()

	hash, err := calculateFileHash(file)
	if err != nil {
	}

	fileData := entities.MedicalData{
		ID:            uuid.New(),
		UserID:        uuid.UUID{},
		FileName:      "",
		FileSize:      0,
		FileExtension: "",
		MimeType:      "",
		S3Key:         "",
		Hash:          hash,
		CreatedAt:     time.Time{},
	}

	err := s.repo.InsertMedicalData(ctx)

}

func calculateFileHash(file multipart.File) (string, error) {
	hash := sha256.New()

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("ошибка при seek: %w", err)
	}

	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("ошибка при подсчете хэша: %w", err)
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("ошибка при втором seek: %w", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
