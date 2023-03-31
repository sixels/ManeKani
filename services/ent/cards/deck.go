package cards

import (
	"context"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/predicate"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
	"github.com/sixels/manekani/services/ent/util"
)

var _ ports.DecksManager = (*CardsRepository)(nil)

func (repo *CardsRepository) QueryDeck(ctx context.Context, deckID uuid.UUID) (*cards.Deck, error) {
	deck, err := repo.client.DeckClient().Query().
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
	return util.Ptr(DeckFromEnt(deck)), nil
}

func (repo *CardsRepository) AllDecks(ctx context.Context, req cards.QueryManyDecksRequest) ([]cards.DeckPartial, error) {
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

	queried, err := repo.client.DeckClient().Query().
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
	return repo.client.DeckClient().Query().
		Where(deck.IDEQ(deckID)).
		QueryOwner().
		OnlyID(ctx)
}

func (repo *CardsRepository) AddDeckSubscriber(ctx context.Context, deckID uuid.UUID, userID string) (int, error) {
	deckProgressID, err := repo.client.DeckProgressClient().Create().
		SetLevel(1).
		SetDeckID(deckID).
		SetUserID(userID).
		OnConflict().
		DoNothing().
		ID(ctx)
	if err != nil {
		return 0, err
	}

	if err := repo.client.DeckClient().UpdateOneID(deckID).
		AddSubscriberIDs(userID).
		Exec(ctx); err != nil {
		return 0, err
	}

	return deckProgressID, nil
}

func (repo *CardsRepository) RemoveDeckSubscriber(ctx context.Context, deckID uuid.UUID, userID string) error {
	return repo.client.DeckClient().UpdateOneID(deckID).
		RemoveSubscriberIDs(userID).
		Exec(ctx)
}

func (repo *CardsRepository) DeckSubscriberExists(ctx context.Context, deckID uuid.UUID, userID string) (int, bool, error) {
	dpID, err := repo.client.DeckProgressClient().Query().
		Where(deckprogress.HasDeckWith(deck.IDEQ(deckID))).
		Where(deckprogress.HasUserWith(user.ID(userID))).
		OnlyID(ctx)

	switch err.(type) {
	case nil:
		return dpID, true, nil
	case *ent.NotFoundError:
		return 0, false, nil
	default:
		return 0, false, err
	}
}

func DeckFromEnt(e *ent.Deck) cards.Deck {
	return cards.Deck{
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

func PartialDeckFromEnt(e *ent.Deck) cards.DeckPartial {
	return cards.DeckPartial{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Owner:       e.Edges.Owner.ID,
	}
}
