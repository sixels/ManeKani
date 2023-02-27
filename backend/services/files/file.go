package files

import (
	"context"
	"fmt"
	"log"

	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/core/domain/files"

	"github.com/minio/minio-go/v7"
)

func (repo FilesRepository) CreateFile(ctx context.Context, req files.CreateFileRequest) (string, error) {
	objectName := objectNameFromFile(req.FileInfo)
	log.Printf("object beign created with key: (%d bytes) '%s'\n", req.Size, objectName)

	uploadInfo, err := repo.minio_client.PutObject(ctx,
		BUCKET_NAME,
		objectName,
		req.Handle,
		req.Size,
		minio.PutObjectOptions{
			ContentType: req.ContentType,
		})

	if err != nil {
		log.Printf("create file '%s' failed with error: %v\n", objectName, err)
		return "", errors.Unknown(err)
	}

	return uploadInfo.Key, nil
}

func (repo FilesRepository) QueryFile(ctx context.Context, key string) (*files.ObjectWrapperResponse, error) {
	log.Printf("querying file '%s'\n", key)
	object, err := repo.minio_client.GetObject(ctx, BUCKET_NAME, key, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("query file '%s' failed with error: %v\n", key, err)
		return nil, errors.Unknown(err)
	}

	info, err := object.Stat()
	if err != nil {
		// NOT FOUND
		log.Println(err)
		return nil, errors.NotFound(fmt.Sprintf("no such file: '%s': %v", key, err))
	}

	return &files.ObjectWrapperResponse{
		ReadCloser:  object,
		Size:        info.Size,
		ContentType: info.ContentType,
	}, nil
}

func (repo FilesRepository) DeleteFile(ctx context.Context, name string) error {
	log.Printf("deleting file '%s'\n", name)
	if err := repo.minio_client.RemoveObject(ctx, BUCKET_NAME, name, minio.RemoveObjectOptions{}); err != nil {
		log.Printf("delete file '%s' failed with error: %v\n", name, err)
		return errors.Unknown(err)
	}
	return nil
}

func objectNameFromFile(f files.FileInfo) string {
	return f.Namespace + "/" + f.Kind + "/" + f.Name
}
