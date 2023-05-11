package users

import (
	"context"
	"log"

	domain "github.com/sixels/manekani/core/domain/user"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/user"
	"github.com/sixels/manekani/services/ent/util"
)

var _ ports.UserRepository = (*UsersRepository)(nil)

func (repo *UsersRepository) Exists(ctx context.Context, id string) (bool, error) {
	exists, err := repo.client.UserClient().Query().Where(user.IDEQ(id)).Exist(ctx)
	if err != nil {
		log.Printf("error checking username: %v\n", err)
		return false, util.ParseEntError(err)
	}
	return exists, nil
}

func (repo *UsersRepository) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	created, err := repo.client.UserClient().Create().
		SetID(req.ID).
		SetEmail(req.Email).
		SetUsername(req.Username).
		Save(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return userFromEnt(created), nil
}

func (repo *UsersRepository) QueryUser(ctx context.Context, id string) (*domain.User, error) {
	queried, err := repo.client.UserClient().Get(ctx, id)
	if err != nil {
		return nil, util.ParseEntError(err)
	}
	return userFromEnt(queried), nil
}

func (repo *UsersRepository) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	doExists, err := repo.client.UserClient().Query().Where(user.UsernameEQ(username)).Exist(ctx)
	if err != nil {
		log.Printf("error checking username %v\n", err)
		return false, util.ParseEntError(err)
	}
	return !doExists, nil
}

func userFromEnt(e *ent.User) *domain.User {
	return &domain.User{
		ID:             e.ID,
		Email:          e.Email,
		Username:       e.Username,
		PendingActions: e.PendingActions,
	}
}
