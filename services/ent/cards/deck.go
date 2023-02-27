package cards

import (
	"context"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/predicate"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
	"github.com/sixels/manekani/services/ent/util"
)

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

func (repo *CardsRepository) AddDeckSubscriber(ctx context.Context, deckID uuid.UUID, userID string) error {
	return repo.client.DeckClient().UpdateOneID(deckID).
		AddSubscriberIDs(userID).
		Exec(ctx)
}

func (repo *CardsRepository) RemoveDeckSubscriber(ctx context.Context, deckID uuid.UUID, userID string) error {
	return repo.client.DeckClient().UpdateOneID(deckID).
		AddSubscriberIDs(userID).
		Exec(ctx)
}

// 	_, err := util.WithTx(ctx, repo.client.Client.(*ent.Client), func(_ *ent.Tx) (_ *struct{}, err error) {
// 		// delete cards
// 		_, err = repo.client.CardClient().Delete().
// 			Where(card.And(
// 				card.HasDeckProgressWith(deckprogress.HasUserWith(
// 					user.IDEQ(userID),
// 				)),
// 				card.HasSubjectWith(
// 					subject.HasDeckWith(deck.IDEQ(deckID))),
// 			)).
// 			Exec(ctx)
// 		if err != nil {
// 			return
// 		}
// 		// delete progress
// 		_, err = repo.client.DeckProgressClient().Delete().
// 			Where(
// 				deckprogress.And(
// 					deckprogress.HasUserWith(user.IDEQ(userID)),
// 					deckprogress.HasDeckWith(deck.IDEQ(deckID)),
// 				),
// 			).
// 			Exec(ctx)
// 		return
// 	})
// 	return err
// }

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
