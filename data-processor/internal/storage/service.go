package storage

import (
	"data-processor/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log/slog"
)

type StorageService struct {
	client *s3.S3
	cfg    config.Config
}

func New(config config.Config) *StorageService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("eu-central-003"),
		Endpoint:         aws.String("s3.eu-central-003.backblazeb2.com"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(
			config.Storage.SecretID,
			config.Storage.Secret,
			"",
		),
	}))

	client := s3.New(sess)

	return &StorageService{client: client, cfg: config}
}

func (s *StorageService) PutData(key string, body io.ReadSeeker) error {
	slog.Info("putting data to storage", "key", key)

	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.cfg.Storage.Bucket),
		Key:    aws.String(key),
		Body:   body,
	})

	return err
}

func (s *StorageService) GetData(key string) ([]byte, error) {
	slog.Info("getting data to storage", "key", key)

	output, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.cfg.Storage.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("error get file from storage: %v", err)
	}
	defer output.Body.Close()

	data, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading object body: %v", err)
	}

	return data, nil
}
