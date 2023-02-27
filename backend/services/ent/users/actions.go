package users

import (
	"context"
	"encoding/json"
	"time"

	user_domain "github.com/sixels/manekani/core/domain/user"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/card"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
)

func (repo *UsersRepository) ResolveActions(ctx context.Context, userID string) error {
	panic("legacy")
	// usr, err := repo.client.UserClient().Query().
	// 	Where(user.IDEQ(userID)).
	// 	Select(user.FieldID, user.FieldPendingActions).
	// 	Only(ctx)
	// if err != nil {
	// 	return err
	// }

	// actions := usr.PendingActions
	// failedActions := []schema.PendingAction{}
	// for _, action := range actions {
	// 	if _, err := util.WithTx(ctx, repo.client.Client.(*ent.Client), func(tx *ent.Tx) (*struct{}, error) {
	// 		switch action.Action {
	// 		case user_domain.ActionCheckCardUnlocks:
	// 			meta, err := redoEncoding[user_domain.CheckCardUnlocksMeta](action.Metadata)
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			return nil, actionCheckCardUnlocks(
	// 				ctx, tx, userID, *meta)
	// 		}

	// 		return nil, nil
	// 	}); err != nil {
	// 		log.Printf("user action failed: %v (%+v): %v \n", action.Action, err, action.Metadata)
	// 		failedActions = append(failedActions, action)
	// 	}
	// }

	// if err := repo.client.UserClient().UpdateOneID(userID).
	// 	SetPendingActions(failedActions).
	// 	Exec(ctx); err != nil {
	// 	return err
	// }

	// return nil
}

func actionCheckCardUnlocks(
	ctx context.Context,
	tx *ent.Tx,
	userID string,
	req user_domain.CheckCardUnlocksMeta,
) error {
	reqCard, err := tx.Card.Query().
		Where(card.And(
			card.IDEQ(req.Card),
			card.HasDeckProgressWith(deckprogress.HasUserWith(
				user.IDEQ(userID),
			)),
		)).
		WithSubject(func(sq *ent.SubjectQuery) { sq.Select(subject.FieldID) }).
		Only(ctx)
	if err != nil {
		return err
	}

	// search for locked dependents of reqCard's subject
	dependents, err := tx.Card.Query().
		Where(card.And(
			card.HasDeckProgressWith(
				deckprogress.HasUserWith(user.IDEQ(userID))),
			card.UnlockedAtIsNil(),
			card.HasSubjectWith(subject.HasDependenciesWith(
				subject.IDEQ(reqCard.Edges.Subject.ID),
			)),
		)).
		WithSubject(func(sq *ent.SubjectQuery) {
			sq.Select(subject.FieldID).
				WithDependencies(func(sq *ent.SubjectQuery) { sq.Select(subject.FieldID) })
		}).
		All(ctx)
	if err != nil {
		return err
	}

	// for each locked dependent, check if all its depencies have been passed
	for _, dependent := range dependents {
		totalDeps := len(dependent.Edges.Subject.Edges.Dependencies)
		depsPassed, err := tx.Card.Query().
			Where(card.And(
				card.HasDeckProgressWith(
					deckprogress.HasUserWith(user.IDEQ(userID))),
				card.PassedAtNotNil(),
				card.HasSubjectWith(subject.HasDependentsWith(
					subject.IDEQ(dependent.Edges.Subject.ID),
				)),
			)).Count(ctx)
		if err != nil {
			return err
		}

		// if all dependencies have passed, we are good to unlock it
		if depsPassed == totalDeps {
			now := time.Now()
			dependent.Update().
				SetUnlockedAt(now).
				SetAvailableAt(now)
		}
	}

	return nil
}

func redoEncoding[T any](t any) (*T, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	data := new(T)
	err = json.Unmarshal(bytes, data)
	return data, err
}

// func (repo *UsersRepository) createLevelCards(ctx context.Context, tx *ent.Tx, userID string, level int32) error {
// 	levelSubjects, err := tx.Subject.Query().
// 		Where(subject.LevelEQ(level)).
// 		Select(subject.FieldID, subject.FieldKind).
// 		All(ctx)

// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	createCards := util.MapArray(levelSubjects, func(levelSubject *ent.Subject) *ent.CardCreate {
// 		cardConstructor := tx.Card.Create().
// 			SetUserID(userID).
// 			SetSubjectID(levelSubject.ID)
// 		if levelSubject.Kind == "radical" {
// 			cardConstructor.SetUnlockedAt(time.Now())
// 		}
// 		return cardConstructor
// 	})

// 	if err = tx.Card.CreateBulk(createCards...).Exec(ctx); err != nil &&
// 		!(ent.IsConstraintError(err) ||
// 			ent.IsNotSingular(err)) {
// 		return err
// 	}

// 	return nil
// }
