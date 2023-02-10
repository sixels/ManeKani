package cards

import (
	"context"

	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/cards/filters"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/deck"
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
