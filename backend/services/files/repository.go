package files

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const BUCKET_NAME = "manekani"

type FilesRepository struct {
	minio_client *minio.Client
}

func NewRepository(ctx context.Context) (FilesRepository, error) {
	endpoint := "minio:9000"
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return FilesRepository{}, fmt.Errorf("failed to connect with the S3 repository: %w", err)
	}

	// create the default bucket if not exists
	bucketExists, err := client.BucketExists(ctx, BUCKET_NAME)
	if err != nil {
		return FilesRepository{}, fmt.Errorf("failed to check if bucket exists: %w", err)
	}
	if !bucketExists {
		if err := client.MakeBucket(
			ctx, BUCKET_NAME,
			minio.MakeBucketOptions{Region: "sa-east1", ObjectLocking: true},
		); err != nil {
			return FilesRepository{}, fmt.Errorf("failed to create the bucket: %w", err)
		}
	}

	return FilesRepository{
		minio_client: client,
	}, nil
}
