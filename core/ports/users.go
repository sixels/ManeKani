package ports

import (
	"context"

	"github.com/sixels/manekani/core/domain/user"
)

// TODO: update user repository methods and add users service
type UserRepository interface {
	CreateUser(ctx context.Context, req user.CreateUserRequest) (*user.User, error)
	IsUsernameAvailable(ctx context.Context, username string) (bool, error)
}