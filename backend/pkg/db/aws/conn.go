package aws

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/seregaa020292/capitalhub/config"
)

// Minio AWS S3 Client constructor
func NewAWSClient(cfg config.AWSConfig) (*minio.Client, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
