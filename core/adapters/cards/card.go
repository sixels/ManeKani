package cards

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/sixels/manekani/core/domain/cards"
)

func (svc *CardsAdapter) QueryCard(ctx context.Context, id uuid.UUID) (*domain.Card, error) {
	return svc.repo.QueryCard(ctx, id)
}
func (svc *CardsAdapter) UpdateCard(ctx context.Context, id uuid.UUID, req domain.UpdateCardRequest) (*domain.Card, error) {
	return svc.repo.UpdateCard(ctx, id, req)
}
func (svc *CardsAdapter) AllCards(ctx context.Context, userID string, req domain.QueryManyCardsRequest) ([]domain.Card, error) {
	return svc.repo.AllCards(ctx, userID, req)
}

// func (svc *CardsService) CreateManyCards(ctx context.Context, deckProgressID int, userID string, reqs []domain.CreateCardRequest) ([]domain.Card, error) {
// 	return svc.repo.CreateManyCards(ctx, deckProgressID, userID, reqs)
// }
