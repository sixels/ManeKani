package users

import (
	"context"
	"fmt"
	"log"

	domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/schema"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *UsersRepository) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	created, err := repo.client.User.Create().
		SetID(req.Id).
		SetEmail(req.Email).
		SetUsername(req.Username).
		SetPendingActions([]schema.PendingAction{
			{
				Action:   domain.CREATE_LEVEL_CARDS,
				Required: true,
				Metadata: domain.CreateLevelCardsMeta{
					Level: 1,
				},
			},
		}).
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

func (repo *UsersRepository) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	doExists, err := repo.client.User.Query().Where(user.UsernameEQ(username)).Exist(ctx)
	if err != nil {
		log.Printf("error checking username %v\n", err)
		return false, util.ParseEntError(err)
	}
	return !doExists, nil
}

func (repo *UsersRepository) QueryUserResolved(ctx context.Context, id string) (*domain.User, error) {
	queried, err := repo.client.User.Get(ctx, id)
	if err != nil {
		return nil, util.ParseEntError(err)
	}

	var failedRequired error
	pendingActions := make([]schema.PendingAction, 0, len(queried.PendingActions))
	for _, pendingAction := range queried.PendingActions {
		_, err := util.WithTx(ctx, repo.client.Client, func(tx *ent.Tx) (*struct{}, error) {
			switch pendingAction.Action {
			case domain.CREATE_LEVEL_CARDS:
				meta := pendingAction.Metadata.(map[string]interface{})
				level := int32(meta["Level"].(float64))

				if err != repo.createLevelCards(ctx, tx, id, level) {
					return nil, err
				}
			}

			return nil, nil
		})

		if err != nil {
			pendingActions = append(pendingActions, pendingAction)
			log.Printf("could not resolve pending action for user: %v: %v", pendingAction, err)

			if pendingAction.Required && failedRequired == nil {
				failedRequired = err
			}
		}

	}

	usr := repo.client.User.UpdateOneID(id).
		SetPendingActions(pendingActions).
		SaveX(ctx)

	if failedRequired != nil {
		return nil, fmt.Errorf("failed to apply a necessary action to the user account")
	}

	return userFromEnt(usr), nil
}

func userFromEnt(e *ent.User) *domain.User {
	return &domain.User{
		Id:             e.ID,
		Email:          e.Email,
		Username:       e.Username,
		Level:          e.Level,
		PendingActions: e.PendingActions,
	}
}
