package users

import (
	"context"
	"log"

	domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *UsersRepository) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	// username is set later
	created, err := repo.client.User.Create().
		SetID(req.Id).
		SetEmail(req.Email).
		SetUsername(req.Username).
		Save(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return userFromEnt(created), nil
}

func (repo *UsersRepository) QueryUser(ctx context.Context, id string) (*domain.User, error) {
	queried, err := repo.client.User.Get(ctx, id)
	if err != nil {
		return nil, util.ParseEntError(err)
	}
	return userFromEnt(queried), nil
}

func (repo *UsersRepository) UsernameAvailable(ctx context.Context, username string) (bool, error) {
	doExists, err := repo.client.User.Query().Where(user.UsernameEQ(username)).Exist(ctx)
	if err != nil {
		log.Println(err)
		return false, util.ParseEntError(err)
	}
	return !doExists, nil
}

func userFromEnt(e *ent.User) *domain.User {
	return &domain.User{
		Id:       e.ID,
		Email:    e.Email,
		Username: e.Username,
	}
}
