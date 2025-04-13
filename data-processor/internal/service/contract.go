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
		SelectByID(ctx context.Context, id string) (entities.MedicalData, error)
	}
	storage interface {
		PutData(key string, body io.ReadSeeker) error
		GetData(key string) ([]byte, error)
	}
)
