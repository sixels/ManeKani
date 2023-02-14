package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/schema"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/services/ent/util"
)

func (repo *UsersRepository) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	created, err := repo.client.User.Create().
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

func (repo *UsersRepository) SubscribeUserToDeck(ctx context.Context, userID string, deckID uuid.UUID) error {
	_, err := util.WithTx(ctx, repo.client.Client, func(tx *ent.Tx) (*struct{}, error) {
		// create the first level cards for now
		deckSubjects, err := repo.client.Deck.Query().
			Where(deck.IDEQ(deckID)).
			QuerySubjects().
			Where(subject.LevelEQ(1)).
			Select(subject.FieldID).
			WithDependencies(func(sq *ent.SubjectQuery) {
				sq.Select(subject.FieldID)
			}).
			All(ctx)
		if err != nil {
			return nil, err
		}

		createdCards, err := tx.Card.
			CreateBulk(
				util.MapArray(deckSubjects, func(subj *ent.Subject) *ent.CardCreate {
					create := tx.Card.Create().
						SetSubjectID(subj.ID).
						SetUnlockedAt(time.Now())
					if len(subj.Edges.Dependencies) == 0 {
						create = create.SetAvailableAt(time.Now())
					}
					return create
				})...,
			).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		// add the progress and subscribe the user
		if err := tx.DeckProgress.Create().
			AddCards(createdCards...).
			SetUserID(userID).
			SetDeckID(deckID).
			Exec(ctx); err != nil {
			return nil, err
		}
		if err := tx.Deck.UpdateOneID(deckID).
			AddSubscriberIDs(userID).
			Exec(ctx); err != nil {
			return nil, err
		}

		return nil, nil

	})

	if err != nil {
		return util.ParseEntError(err)
	}
	return nil
}

func (repo *UsersRepository) QueryUser(ctx context.Context, id string) (*domain.User, error) {
	queried, err := repo.client.User.Get(ctx, id)
	if err != nil {
		return nil, util.ParseEntError(err)
	}
	return userFromEnt(queried), nil
}

func (repo *UsersRepository) QueryUserResolved(ctx context.Context, id string) (*domain.User, error) {
	queried, err := repo.client.User.Get(ctx, id)
	if err != nil {
		return nil, util.ParseEntError(err)
	}

	var failedRequired error
	pendingActions := make([]schema.PendingAction, 0, len(queried.PendingActions))
	for _, pendingAction := range queried.PendingActions {
		err := func() error {
			switch pendingAction.Action {
			case domain.SUBSCRIBE_TO_DECK:
				meta := pendingAction.Metadata.(map[string]interface{})
				deckID := uuid.MustParse(meta["Deck"].(string))
				return repo.SubscribeUserToDeck(ctx, id, deckID)
			}
			return nil
		}()
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
