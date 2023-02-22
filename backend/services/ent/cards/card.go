package cards

import (
	"context"
	"time"

	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/cards/filters"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/deckprogress"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *CardsRepository) AllUserCards(ctx context.Context, userID string, req *cards.QueryManyCardsRequest) ([]*cards.Card, error) {
	fs := filters.NewFilter([]predicate.Card{
		card.HasDeckProgressWith(deckprogress.HasUserWith(user.IDEQ(userID))),
	})

	if req != nil {
		filters.In(fs, req.IDs.Separate(), card.IDIn)
		filters.In(fs, req.Decks.Separate(), func(ids ...uuid.UUID) predicate.Card {
			return card.HasSubjectWith(subject.HasDeckWith(deck.IDIn(ids...)))
		})
		filters.In(fs, req.Levels.Separate(), func(levels ...int32) predicate.Card {
			return card.HasSubjectWith(subject.LevelIn(levels...))
		})
		filters.With(fs, req.AvailableAfter, card.AvailableAtGTE)
		filters.With(fs, req.AvailableBefore, card.AvailableAtLTE)
		filters.With(fs, req.IsUnlocked, func(bool) predicate.Card { return card.UnlockedAtLTE(time.Now()) })
		filters.With(fs, req.IsBurned, func(bool) predicate.Card { return card.BurnedAtLTE(time.Now()) })
		filters.With(fs, req.ProgressBefore, card.ProgressLTE)
		filters.With(fs, req.ProgressAfter, card.ProgressGTE)
	}

	page := 0
	if req.Page != nil {
		page = int(*req.Page)
	}

	queried, err := repo.client.Card.Query().
		Limit(500).
		Offset(page).
		Where(card.And(fs.Filters()...)).
		WithSubject(func(sq *ent.SubjectQuery) {
			sq.
				WithOwner(func(uq *ent.UserQuery) {
					uq.Select(user.FieldID)
				}).
				WithDependencies(func(sq *ent.SubjectQuery) {
					sq.Select(subject.FieldID)
				}).
				WithDependents(func(sq *ent.SubjectQuery) {
					sq.Select(subject.FieldID)
				}).
				WithSimilar(func(sq *ent.SubjectQuery) {
					sq.Select(subject.FieldID)
				}).
				WithDeck(func(dq *ent.DeckQuery) {
					dq.Select(deck.FieldID)
				}).
				Select(
					subject.FieldID,
					subject.FieldKind,
					subject.FieldLevel,
					subject.FieldPriority,
					subject.FieldName,
					subject.FieldValue,
					subject.FieldValueImage,
					subject.FieldSlug,
					subject.FieldStudyData,
				)
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(queried, CardFromEnt), nil
}

func CardFromEnt(e *ent.Card) *cards.Card {
	return &cards.Card{
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
		Subject:     *PartialSubjectFromEnt(e.Edges.Subject),
	}
}
