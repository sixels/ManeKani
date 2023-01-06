package files

import (
	"context"
	"fmt"
	"log"

	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/core/domain/files"

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
			ContentType: "image/png",
		})

	if err != nil {
		log.Printf("create file '%s' failed with error: %v\n", objectName, err)
		return "", errors.Unknown(err)
	}

	return uploadInfo.Key, nil
}

func (repo FilesRepository) QueryFile(ctx context.Context, name string) (*files.ObjectWrapperResponse, error) {
	log.Printf("querying file '%s'\n", name)
	object, err := repo.minio_client.GetObject(ctx, BUCKET_NAME, name, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("query file '%s' failed with error: %v\n", name, err)
		return nil, errors.Unknown(err)
	}

	info, err := object.Stat()
	if err != nil {
		// NOT FOUND
		return nil, errors.NotFound(fmt.Sprintf("no such file: '%s'", name))
	}

	return &files.ObjectWrapperResponse{
		ReadCloser:  object,
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
	return f.Kind + "/" + f.Namespace + "/" + f.Name
}
