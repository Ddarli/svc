package entities

import (
	"github.com/google/uuid"
	"time"
)

type MedicalData struct {
	ID            uuid.UUID `db:"id"`
	UserID        uuid.UUID `db:"user_id"`
	FileName      string    `db:"file_name"`
	FileSize      int       `db:"file_size"`
	FileExtension string    `db:"file_extension"`
	MimeType      string    `db:"mime_type"`
	S3Key         string    `db:"s3_key"`
	Hash          string    `db:"hash"`
	CreatedAt     time.Time `db:"created_at"`
}
