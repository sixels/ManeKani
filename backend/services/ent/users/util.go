package users

import (
	"context"
	"log"

	domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

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
