package repository

import (
	"context"
	"data-processor/internal/repository/entities"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rotisserie/eris"
	"log"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

func (r *repo) InsertMedicalData(ctx context.Context, file entities.MedicalData) error {
	_, err := r.db.Exec(ctx, insertRecord, file.ID, file.UserID, file.FileName, file.Description, file.FileSize, file.MimeType, file.S3Key, file.Hash)
	if err != nil {
		log.Printf("Error inserting new user: %v", err)
	}

	return err
}

func (r *repo) SelectMedicalDataByUserID(ctx context.Context, userID uuid.UUID) ([]entities.MedicalData, error) {
	rows, err := r.db.Query(ctx, selectRecord, userID.String())
	if err != nil {
		return nil, fmt.Errorf("error fetch data from db: %v", err)
	}

	var records []entities.MedicalData

	for rows.Next() {
		var record entities.MedicalData

		err := rows.Scan(&record.ID, &record.UserID, &record.FileName, &record.Description, &record.FileSize, &record.MimeType, &record.S3Key, &record.Hash)
		if err != nil {
			return []entities.MedicalData{}, fmt.Errorf("error fetching data from database: %v", err)
		}

		records = append(records, record)
	}

	return records, nil
}

func (r *repo) SelectByID(ctx context.Context, id string) (entities.MedicalData, error) {
	if id == "" {
		return entities.MedicalData{}, eris.New("id is empty")
	}
	row := r.db.QueryRow(ctx, selectByID, id)

	var record entities.MedicalData

	err := row.Scan(&record.ID, &record.UserID, &record.FileName, &record.Description, &record.FileSize, &record.MimeType, &record.S3Key, &record.Hash)
	if err != nil {
		return entities.MedicalData{}, eris.Wrapf(err, "error in repoository")
	}

	return record, nil
}
