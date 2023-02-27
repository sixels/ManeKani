package users

import (
	"context"

	domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *UsersRepository) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	created, err := repo.client.UserClient().Create().
		SetID(req.ID).
		SetEmail(req.Email).
		SetUsername(req.Username).
		// SetPendingActions([]schema.PendingAction{}).
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
