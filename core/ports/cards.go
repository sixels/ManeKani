package ports

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/ports/transactions"
)

type (
	SubjectsManager interface {
		CreateSubject(ctx context.Context, ownerID string, req domain.CreateSubjectRequest) (*domain.Subject, error)
		QuerySubject(ctx context.Context, id uuid.UUID) (*domain.Subject, error)
		UpdateSubject(ctx context.Context, id uuid.UUID, req domain.UpdateSubjectRequest) (*domain.Subject, error)
		DeleteSubject(ctx context.Context, id uuid.UUID) error
		AllSubjects(ctx context.Context, req domain.QueryManySubjectsRequest) ([]domain.PartialSubject, error)
		SubjectOwner(ctx context.Context, id uuid.UUID) (string, error)
	}

	// TODO: create, update, delete deck
	DecksManager interface {
		QueryDeck(ctx context.Context, id uuid.UUID) (*domain.Deck, error)
		AllDecks(ctx context.Context, req domain.QueryManyDecksRequest) ([]domain.DeckPartial, error)
		DeckOwner(ctx context.Context, id uuid.UUID) (string, error)

		AddDeckSubscriber(ctx context.Context, id uuid.UUID, userID string) (deckProgressID int, err error)
		RemoveDeckSubscriber(ctx context.Context, id uuid.UUID, userID string) error
	}

	// TODO: query review
	ReviewsManager interface {
		CreateReview(ctx context.Context, userID string, req domain.CreateReviewRequest) (*domain.Review, error)
		AllReviews(ctx context.Context, userID string, req domain.QueryManyReviewsRequest) ([]domain.Review, error)
	}

	CardsManager interface {
		QueryCard(ctx context.Context, id uuid.UUID) (*domain.Card, error)
		UpdateCard(ctx context.Context, id uuid.UUID, req domain.UpdateCardRequest) (*domain.Card, error)
		AllCards(ctx context.Context, userID string, req domain.QueryManyCardsRequest) ([]domain.Card, error)
		CreateManyCards(ctx context.Context, deckProgressID int, userID string, reqs []domain.CreateCardRequest) ([]domain.Card, error)
	}

	CardsRepository interface {
		transactions.TransactionalRepository

		SubjectsManager
		DecksManager
		ReviewsManager
		CardsManager
	}
)

// TODO define service interface
// type CardsService interface {
// }
