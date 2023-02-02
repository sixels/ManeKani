package users

import (
	"context"
	"log"
	"time"

	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/services/ent/util"
)

func (repo *UsersRepository) createLevelCards(ctx context.Context, tx *ent.Tx, userID string, level int32) error {
	levelSubjects, err := tx.Subject.Query().
		Where(subject.LevelEQ(level)).
		Select(subject.FieldID, subject.FieldKind).
		All(ctx)

	if err != nil {
		log.Println(err)
		return err
	}

	createCards := util.MapArray(levelSubjects, func(levelSubject *ent.Subject) *ent.CardCreate {
		cardConstructor := tx.Card.Create().
			SetUserID(userID).
			SetSubjectID(levelSubject.ID)

		if levelSubject.Kind == subject.KindRadical {
			cardConstructor.SetUnlockedAt(time.Now())
		}

		return cardConstructor
	})

	if err = tx.Card.CreateBulk(createCards...).Exec(ctx); err != nil &&
		!(ent.IsConstraintError(err) ||
			ent.IsNotSingular(err)) {
		return err
	}

	return nil
}
