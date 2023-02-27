package cards

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain "sixels.io/manekani/core/domain/cards"
)

func (svc *CardsService) CreateSubject(ctx context.Context, ownerID string, req domain.CreateSubjectRequest) (*domain.Subject, error) {
	isOwner, err := checkResourceOwner(ctx, req.Deck, ownerID, svc.repo.DeckOwner)
	if err != nil {
		return nil, err
	}
	if !isOwner {
		return nil, fmt.Errorf("%q is not the owner of '%v'", ownerID, req.Deck)
	}

	return svc.repo.CreateSubject(ctx, ownerID, req)
}

func (svc *CardsService) QuerySubject(ctx context.Context, id uuid.UUID) (*domain.Subject, error) {
	return svc.repo.QuerySubject(ctx, id)
}

func (svc *CardsService) UpdateSubject(ctx context.Context, id uuid.UUID, userID string, req domain.UpdateSubjectRequest) (*domain.Subject, error) {
	isOwner, err := checkResourceOwner(ctx, id, userID, svc.repo.SubjectOwner)
	if err != nil {
		return nil, err
	}
	if !isOwner {
		return nil, fmt.Errorf("%q is not the owner of '%v'", userID, id)
	}

	return svc.repo.UpdateSubject(ctx, id, req)
}

func (svc *CardsService) DeleteSubject(ctx context.Context, id uuid.UUID, userID string) error {
	isOwner, err := checkResourceOwner(ctx, id, userID, svc.repo.DeckOwner)
	if err != nil {
		return err
	}
	if !isOwner {
		return fmt.Errorf("%q is not the owner of '%v'", userID, id)
	}

	return svc.repo.DeleteSubject(ctx, id)
}

func (svc *CardsService) AllSubjects(ctx context.Context, req domain.QueryManySubjectsRequest) ([]domain.PartialSubject, error) {
	return svc.repo.AllSubjects(ctx, req)
}
