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

func (repo *CardsRepository) QueryDeck(ctx context.Context, deckID uuid.UUID) (*cards.Deck, error) {
	deck, err := repo.client.Deck.Query().
		WithOwner(func(uq *ent.UserQuery) {
			uq.Select(user.FieldID)
		}).
		WithSubjects(func(sq *ent.SubjectQuery) {
			sq.Select(subject.FieldID)
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return DeckFromEnt(deck), nil
}

func (repo *CardsRepository) AllDecks(ctx context.Context, req cards.QueryManyDecksRequest) ([]*cards.DeckPartial, error) {
	var reqFilters []predicate.Deck
	reqFilters = filters.ApplyFilter(reqFilters, req.IDs.Separate(), deck.IDIn)
	reqFilters = filters.ApplyFilter(reqFilters, req.Owners.Separate(), func(ids ...string) predicate.Deck {
		return deck.HasOwnerWith(user.IDIn(ids...))
	})
	reqFilters = filters.ApplyFilter(reqFilters, req.Subjects.Separate(), func(ids ...uuid.UUID) predicate.Deck {
		return deck.HasSubjectsWith(subject.IDIn(ids...))
	})
	reqFilters = filters.ApplyFilter(reqFilters, req.Names.Separate(), deck.NameIn)

	page := 0
	if req.Page != nil {
		page = int(*req.Page)
	}

	queried, err := repo.client.Deck.Query().
		Limit(500).
		Offset(page).
		WithOwner(func(uq *ent.UserQuery) {
			uq.Select(user.FieldID)
		}).
		Select(
			deck.FieldID,
			deck.FieldName,
			deck.FieldDescription,
		).
		Where(deck.And(reqFilters...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(queried, PartialDeckFromEnt), nil
}

func (repo *CardsRepository) DeckOwner(ctx context.Context, deckID uuid.UUID) (string, error) {
	return repo.client.Deck.Query().
		Where(deck.IDEQ(deckID)).
		QueryOwner().
		OnlyID(ctx)
}

func (repo *CardsRepository) SubscribeUserToDeck(ctx context.Context, userID string, deckID uuid.UUID) error {
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
						SetSubjectID(subj.ID)
					if len(subj.Edges.Dependencies) == 0 {
						now := time.Now()
						create = create.
							SetAvailableAt(now).
							SetUnlockedAt(now)
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

func (repo *CardsRepository) UnsubscribeUserFromDeck(ctx context.Context, userID string, deckID uuid.UUID) error {
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

func DeckFromEnt(e *ent.Deck) *cards.Deck {
	return &cards.Deck{
		ID:          e.ID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		Name:        e.Name,
		Description: e.Description,
		Subjects: util.MapArray(e.Edges.Subjects, func(s *ent.Subject) uuid.UUID {
			return s.ID
		}),
		Owner: e.Edges.Owner.ID,
	}
}

func PartialDeckFromEnt(e *ent.Deck) *cards.DeckPartial {
	return &cards.DeckPartial{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Owner:       e.Edges.Owner.ID,
	}
}
