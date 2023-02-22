package ports

import (
	"context"

	"github.com/google/uuid"
	domain "sixels.io/manekani/core/domain/cards"
)

// TODO: update repository methods
type (
	SubjectRepository interface {
		CreateSubject(ctx context.Context, req domain.CreateSubjectRequest) (*domain.Subject, error)
		QuerySubject(ctx context.Context, id uuid.UUID) (*domain.Subject, error)
		UpdateSubject(ctx context.Context, id uuid.UUID, req domain.UpdateSubjectRequest) (*domain.Subject, error)
		DeleteSubject(ctx context.Context, id uuid.UUID) error
		AllSubjects(ctx context.Context, req domain.QueryManySubjectsRequest) ([]*domain.PartialSubject, error)
	}

	DeckRepository interface {
		QueryDeck(ctx context.Context, id uuid.UUID) (*domain.Deck, error)
		AllDecks(ctx context.Context, req domain.QueryManyDecksRequest) ([]*domain.DeckPartial, error)
	}
)
