package users

import (
	"context"

	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/srs"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/deckprogress"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
	cards_svc "sixels.io/manekani/services/ent/cards"
	"sixels.io/manekani/services/ent/util"
)

func (repo *UsersRepository) GetUserCards(ctx context.Context, userID string, deckID uuid.UUID) ([]*srs.Card, error) {
	user, err := repo.client.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	queriedCards, err := user.QueryDecksProgress().
		Where(deckprogress.HasDeckWith(deck.IDEQ(deckID))).
		QueryCards().
		WithSubject(func(sq *ent.SubjectQuery) {
			sq.WithDependencies(func(sq *ent.SubjectQuery) {
				sq.Select(subject.FieldID)
			})
			sq.WithDependents(func(sq *ent.SubjectQuery) {
				sq.Select(subject.FieldID)
			})
			sq.WithSimilar(func(sq *ent.SubjectQuery) {
				sq.Select(subject.FieldID)
			})
			sq.WithDeck(func(dq *ent.DeckQuery) {
				sq.Select(deck.FieldID)
			})
		}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return util.MapArray(queriedCards, cardFromEnt), nil
}

func (repo *UsersRepository) UnsubscribeUserFromDeck(ctx context.Context, userID string, deckID uuid.UUID) error {
	_, err := util.WithTx(ctx, repo.client.Client, func(tx *ent.Tx) (_ *struct{}, err error) {
		// delete cards
		_, err = repo.client.Card.Delete().
			Where(card.And(
				card.HasDeckProgressWith(deckprogress.HasUserWith(
					user.IDEQ(userID),
				)),
				card.HasSubjectWith(
					subject.HasDeckWith(deck.IDEQ(deckID))),
			)).
			Exec(ctx)
		if err != nil {
			return
		}
		// delete progress
		_, err = repo.client.DeckProgress.Delete().
			Where(
				deckprogress.And(
					deckprogress.HasUserWith(user.IDEQ(userID)),
					deckprogress.HasDeckWith(deck.IDEQ(deckID)),
				),
			).
			Exec(ctx)
		return
	})
	return err
}

func cardFromEnt(e *ent.Card) *srs.Card {
	return &srs.Card{
		Id:          e.ID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		Progress:    e.Progress,
		TotalErrors: e.TotalErrors,
		UnlockedAt:  e.UnlockedAt,
		StartedAt:   e.StartedAt,
		PassedAt:    e.PassedAt,
		AvailableAt: e.AvailableAt,
		BurnedAt:    e.BurnedAt,
		Subject:     *cards_svc.SubjectFromEnt(e.Edges.Subject),
	}
}
