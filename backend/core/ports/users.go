package ports

import (
	"context"

	"sixels.io/manekani/core/domain/user"
)

// TODO: update user repository methods
type UserRepository interface {
	CreateUser(ctx context.Context, req user.CreateUserRequest) (*user.User, error)
	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
}
