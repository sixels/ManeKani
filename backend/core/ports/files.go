package ports

import (
	"context"

	domain "github.com/sixels/manekani/core/domain/files"
)

type FilesRepository interface {
	CreateFile(ctx context.Context, req domain.CreateFileRequest) (string, error)
	QueryFile(ctx context.Context, name string) (*domain.ObjectWrapperResponse, error)
	DeleteFile(ctx context.Context, name string) error
}
