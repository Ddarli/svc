package storage

import (
	"data-processor/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"path"
)

type StorageService struct {
	client *s3.S3
}

func New(config config.Config) *StorageService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("eu-central-003"),
		Endpoint:         aws.String("s3.eu-central-003.backblazeb2.com"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(
			"0030b235a057b230000000002",
			"K003HKX3VLdwXwW9TKG6M3x666AZcC8",
			"",
		),
	}))

	client := s3.New(sess)

	return &StorageService{client: client}
}

func (s *StorageService) PutData(bucket, key string, body io.ReadSeeker) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})

	return err
}

func (s *StorageService) GetData(bucketName, key string) error {
	output, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error get file from storage: %v", err)
	}
	defer output.Body.Close()

	fileName := path.Base(key)

	file, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, output.Body)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
