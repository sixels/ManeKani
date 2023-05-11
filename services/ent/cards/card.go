package cards

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/card"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/predicate"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
	"github.com/sixels/manekani/services/ent/util"
)

var _ ports.CardsManager = (*CardsRepository)(nil)

func (repo *CardsRepository) QueryCard(ctx context.Context, id uuid.UUID) (*cards.Card, error) {
	queried, err := repo.client.CardClient().
		Query().
		Where(card.IDEQ(id)).
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
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return util.Ptr(CardFromEnt(queried)), nil
}

func (repo *CardsRepository) UpdateCard(ctx context.Context, id uuid.UUID, req cards.UpdateCardRequest) (*cards.Card, error) {
	updater := repo.client.CardClient().UpdateOneID(id).
		SetNillableProgress(req.Progress).
		SetNillableTotalErrors(req.TotalErrors)

	setOrClear(req.UnlockedAt, updater.SetUnlockedAt, updater.ClearUnlockedAt)
	setOrClear(req.StartedAt, updater.SetStartedAt, updater.ClearStartedAt)
	setOrClear(req.PassedAt, updater.SetPassedAt, updater.ClearPassedAt)
	setOrClear(req.AvailableAt, updater.SetAvailableAt, updater.ClearAvailableAt)
	setOrClear(req.BurnedAt, updater.SetBurnedAt, updater.ClearBurnedAt)

	if err := updater.Exec(ctx); err != nil {
		return nil, err
	}

	// TODO: better way to get edges from updated instead of querying again
	queried, err := repo.client.CardClient().
		Query().
		Where(card.IDEQ(id)).
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
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return util.Ptr(CardFromEnt(queried)), nil
}

func (repo *CardsRepository) AllCards(ctx context.Context, userID string, req cards.QueryManyCardsRequest) ([]cards.Card, error) {
	fs := filters.NewFilter([]predicate.Card{
		card.HasDeckProgressWith(deckprogress.HasUserWith(user.IDEQ(userID))),
	})

	filters.In(fs, req.IDs.Separate(), card.IDIn)
	filters.In(fs, req.Decks.Separate(), func(ids ...uuid.UUID) predicate.Card {
		return card.HasSubjectWith(subject.HasDeckWith(deck.IDIn(ids...)))
	})
	filters.In(fs, req.Levels.Separate(), func(levels ...int32) predicate.Card {
		return card.HasSubjectWith(subject.LevelIn(levels...))
	})
	filters.With(fs, req.AvailableAfter, card.AvailableAtGTE)
	filters.With(fs, req.AvailableBefore, card.AvailableAtLTE)
	filters.With(fs, req.IsUnlocked, func(c bool) predicate.Card {
		p := card.UnlockedAtNotNil()
		if !c {
			return card.Not(p)
		}
		return p
	})
	filters.With(fs, req.IsBurned, func(c bool) predicate.Card {
		p := card.BurnedAtNotNil()
		if !c {
			return card.Not(p)
		}
		return p
	})
	filters.With(fs, req.ProgressBefore, card.ProgressLTE)
	filters.With(fs, req.ProgressAfter, card.ProgressGTE)
	filters.With(fs, req.IsStarted, func(c bool) predicate.Card {
		p := card.StartedAtNotNil()
		if !c {
			return card.Not(p)
		}
		return p
	})
	filters.In(fs, req.WithDependencies.Separate(), func(ids ...uuid.UUID) predicate.Card {
		return card.HasSubjectWith(subject.HasDependenciesWith(subject.IDIn(ids...)))
	})
	filters.In(fs, req.WithDependents.Separate(), func(ids ...uuid.UUID) predicate.Card {
		return card.HasSubjectWith(subject.HasDependentsWith(subject.IDIn(ids...)))
	})

	log.Println(repo.client.CardClient().Query().Where(
		card.And(
			card.HasDeckProgressWith(deckprogress.HasUserWith(user.IDEQ(userID))),
			card.HasSubjectWith(subject.HasDependenciesWith(subject.IDEQ(
				uuid.MustParse("1175bc9c-ec71-492a-81e9-64400317e8c3"),
			))),
			card.UnlockedAtNotNil(),
		),
	).CountX(ctx))

	page := 0
	if req.Page != nil {
		page = int(*req.Page)
	}

	queried, err := repo.client.CardClient().
		Query().
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

func (repo *CardsRepository) CreateManyCards(ctx context.Context, deckProgressID int, userID string, reqs []cards.CreateCardRequest) ([]cards.Card, error) {
	toCreate := make([]*ent.CardCreate, len(reqs))
	for i, req := range reqs {
		toCreate[i] = repo.client.CardClient().Create().
			SetDeckProgressID(deckProgressID).
			SetSubjectID(req.SubjectID).
			SetNillableProgress(req.Progress).
			SetNillableTotalErrors(req.TotalErrors).
			SetNillableUnlockedAt(req.UnlockedAt).
			SetNillableAvailableAt(req.AvailableAt)
	}
	log.Printf("creating %d cards", len(toCreate))
	created, err := repo.client.CardClient().CreateBulk(toCreate...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: better way to get edges from updated instead of querying again
	queried, err := repo.client.CardClient().
		Query().
		Where(card.IDIn(util.MapArray(created, func(a *ent.Card) uuid.UUID {
			return a.ID
		})...)).
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

func CardFromEnt(e *ent.Card) cards.Card {
	return cards.Card{
		ID:          e.ID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		Progress:    e.Progress,
		TotalErrors: e.TotalErrors,
		UnlockedAt:  e.UnlockedAt,
		StartedAt:   e.StartedAt,
		PassedAt:    e.PassedAt,
		AvailableAt: e.AvailableAt,
		BurnedAt:    e.BurnedAt,
		Subject:     PartialSubjectFromEnt(e.Edges.Subject),
	}
}

func setOrClear[T any, U any](v **T, set func(T) U, clear func() U) {
	if v == nil {
		return
	}

	if *v != nil {
		set(**v)
	} else {
		clear()
	}
}
