package cards

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	domain "github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"github.com/sixels/manekani/core/ports/transactions"
	"github.com/sixels/manekani/server/api/cards/util"
)

// TODO: func (svc *CardsService) CreateDeck(ctx context.Context, userID string, req domain.CreateDeckRequest) (*domain.Deck, error)
// TODO: func (svc *CardsService) UpdateDeck(ctx context.Context, userID string, id uuid.UUID, req domain.UpdateDeckRequest) (*domain.Deck, error)
// TODO: func (svc *CardsService) DeleteDeck(ctx context.Context, userID string, id uuid.UUID) error

func (svc *CardsService) QueryDeck(ctx context.Context, id uuid.UUID) (*domain.Deck, error) {
	return svc.repo.QueryDeck(ctx, id)
}

func (svc *CardsService) AllDecks(ctx context.Context, req domain.QueryManyDecksRequest) ([]domain.DeckPartial, error) {
	return svc.repo.AllDecks(ctx, req)
}

func (svc *CardsService) AddDeckSubscriber(ctx context.Context, id uuid.UUID, userID string) error {
	_, err := svc.QueryDeck(ctx, id)
	if err != nil {
		return err
	}
	// TODO: check deck visibility (private/public/shared)

	tx := transactions.Begin(ctx)
	txCards, err := transactions.MakeTransactional(tx, svc.repo)
	if err != nil {
		return err
	}

	return tx.Run(func(ctx context.Context) error {
		// TODO: check if user already has cards in this deck (i.e: was already subscribed in the past or still is)
		deckProgressID, err := txCards.AddDeckSubscriber(ctx, id, userID)
		if err != nil {
			log.Printf("could not subscribe user %v to deck %v: %v", userID, id, err)
			return err
		}

		// get level 1 subjects
		subjects, err := txCards.AllSubjects(
			ctx,
			domain.QueryManySubjectsRequest{
				FilterDecks: filters.FilterDecks{
					Decks: util.Ptr((filters.CommaSeparatedUUID)(id.String())),
				},
				FilterLevels: filters.FilterLevels{
					Levels: util.Ptr((filters.CommaSeparatedInt32)("1")),
				},
			},
		)
		if err != nil {
			return err
		}
		log.Println("level 1 subjects:", len(subjects))

		// create cards from subjects
		cards := make([]domain.CreateCardRequest, len(subjects))
		for i, subject := range subjects {
			var (
				unlockedAt  *time.Time = nil
				availableAt *time.Time = nil
			)
			// unlock cards which have no dependencies
			if len(subject.Dependencies) == 0 {
				now := time.Now()
				unlockedAt = &now
				availableAt = &now
			}
			cards[i] = domain.CreateCardRequest{
				SubjectID:   subject.ID,
				UnlockedAt:  unlockedAt,
				AvailableAt: availableAt,
			}
		}

		if _, err := txCards.CreateManyCards(ctx, deckProgressID, userID, cards); err != nil {
			log.Printf("could not create user %v cards from deck %v: %v", userID, id, err)
			return err
		}
		return nil
	})
}

func (svc *CardsService) RemoveDeckSubscriber(ctx context.Context, id uuid.UUID, userID string) error {
	return svc.repo.RemoveDeckSubscriber(ctx, id, userID)
}

func (svc *CardsService) ResetDeckToLevel(ctx context.Context, id uuid.UUID, userID string, level int32) error {
	return fmt.Errorf("TODO: ResetDeckToLevel")
}
