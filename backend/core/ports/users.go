package ports

import (
	"context"

	"sixels.io/manekani/core/domain/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req user.CreateUserRequest) (*user.User, error)
	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
}
