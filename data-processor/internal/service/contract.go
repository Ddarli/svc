package service

import (
	"context"
	"data-processor/internal/repository/entities"
	"github.com/google/uuid"
	"io"
)

type (
	repo interface {
		SelectMedicalDataByUserID(ctx context.Context, userID uuid.UUID) ([]entities.MedicalData, error)
		InsertMedicalData(ctx context.Context, file entities.MedicalData) error
	}
	storage interface {
		PutData(bucket, key string, body io.ReadSeeker) error
		GetData(bucketName, key string) error
	}
)
