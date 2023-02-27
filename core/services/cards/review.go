package cards

import (
	"context"
	"fmt"
	"time"

	domain "github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/core/ports/transactions"
	"github.com/sixels/manekani/server/api/cards/util"
)

func (svc *CardsService) CreateReview(ctx context.Context, userID string, req domain.CreateReviewAPIRequest) (*domain.Review, error) {
	tx := transactions.Begin(ctx)
	txCards, err := transactions.MakeTransactional(tx, svc.repo)
	if err != nil {
		return nil, err
	}
	return transactions.RunWithResult(tx, func(ctx context.Context) (*domain.Review, error) {
		card, err := txCards.QueryCard(ctx, *req.CardID.Only())
		if err != nil {
			return nil, err
		}

		if err := validateSessionType(req.SessionType, card); err != nil {
			return nil, err
		}

		startProgress := card.Progress
		totalErrors := uint(0)
		for _, errQty := range req.Errors {
			totalErrors += uint(errQty)
		}

		endProgress := calculateNextProgress(req.SessionType, startProgress, totalErrors)
		txCards.UpdateCard(ctx, card.ID,
			cardUpdates(card, endProgress, int32(totalErrors)))

		review, err := txCards.CreateReview(ctx, userID, domain.CreateReviewRequest{
			CardID:        *req.CardID.Only(),
			Errors:        req.Errors,
			StartProgress: startProgress,
			EndProgress:   endProgress,
		})
		if err != nil {
			return nil, err
		}

		if err := unlockDependents(ctx, txCards, card, userID); err != nil {
			return nil, err
		}

		return review, nil
	})
}

func (svc *CardsService) AllReviews(ctx context.Context, userID string, req domain.QueryManyReviewsRequest) ([]domain.Review, error) {
	return svc.repo.AllReviews(ctx, userID, req)
}

func unlockDependents(ctx context.Context, txCards ports.CardsRepository, card *domain.Card, userID string) error {
	// search for locked dependents of card's subject
	dependents, err := txCards.AllCards(ctx, userID, domain.QueryManyCardsRequest{
		WithDependencies: util.Ptr((filters.CommaSeparatedUUID)(card.Subject.ID.String())),
		IsUnlocked:       util.Ptr(false),
	})
	if err != nil {
		return err
	}

	// for each locked dependent, check if all its depencies have been passed
	for _, dependent := range dependents {
		totalDeps := len(dependent.Subject.Dependencies)
		depsPassed, err := txCards.AllCards(ctx, userID, domain.QueryManyCardsRequest{
			IsPassed:       util.Ptr(true),
			WithDependents: util.Ptr((filters.CommaSeparatedUUID)(dependent.Subject.ID.String())),
		})
		if err != nil {
			return err
		}

		// if all dependencies have passed, we are good to unlock it
		if len(depsPassed) == totalDeps {
			now := time.Now()
			txCards.UpdateCard(ctx, dependent.ID, domain.UpdateCardRequest{
				UnlockedAt:  util.Ptr(&now),
				AvailableAt: util.Ptr(&now),
			})
		}
	}

	return nil
}

// Return a card update request based on the review result
func cardUpdates(card *domain.Card, endProgress uint8, totalErrors int32) domain.UpdateCardRequest {
	cardUpdates := domain.UpdateCardRequest{
		Progress:    &endProgress,
		TotalErrors: util.Ptr(card.TotalErrors + totalErrors),
	}
	now := time.Now()

	availableAt := calculateAvailableAt(endProgress, card.Subject.Level)
	if availableAt != nil {
		cardUpdates.AvailableAt = &availableAt
	} else {
		cardUpdates.AvailableAt = util.Ptr[*time.Time](nil)
		cardUpdates.BurnedAt = util.Ptr(&now)
	}
	if card.StartedAt == nil {
		cardUpdates.StartedAt = util.Ptr(&now)
	}
	if endProgress >= 5 && card.PassedAt == nil {
		cardUpdates.PassedAt = util.Ptr(&now)
	}
	return cardUpdates
}

// Checks if a given card is valid for this session type
func validateSessionType(sessionType domain.SessionType, card *domain.Card) error {
	// a lesson is only available if the card was not started already
	if sessionType == domain.SessionLesson && card.StartedAt != nil {
		return fmt.Errorf("card is not available for lesson")
	}
	// a review is only available for cards that already passed in a lesson
	if sessionType == domain.SessionReview && card.StartedAt == nil {
		return fmt.Errorf("card is not available for review")
	}
	return nil
}

/*
Calculate the next progress based on the amount of errors
if there are some error, the next progress will be:

	endProgress := startProgress - (math.Ceil(errors / penaltyScale) * penaltyScale)

Otherwise, it will just add one to the starting progress
*/
func calculateNextProgress(sessionType domain.SessionType, startProgress uint8, errors uint) (endProgress uint8) {
	var penaltyScale uint8 = 1
	if startProgress >= 5 {
		penaltyScale = 2
	}

	if errors == 0 || sessionType == domain.SessionLesson {
		endProgress = startProgress + 1
	} else {
		errorAdjustment := (errors + uint(penaltyScale) - 1) / uint(penaltyScale)
		penalty := errorAdjustment * uint(penaltyScale)

		if penalty > uint(startProgress) {
			penalty = uint(startProgress)
		}
		endProgress = startProgress - uint8(penalty)
	}

	return endProgress
}

/*
Calculate the time when the next review will be ready for a card,
according to its current progress

	 `
		| progress | wait time |
		| :------: | :-------: |
		|     1    |    4h     | (half the time for level 1 and 2)
		|     2    |    8h     | (half the time for level 1 and 2)
		|     3    |    1d     |
		|     4    |    2d     |
		|     5    |    1w     |
		|     6    |    2w     |
		|     7    |    1mo    |
		|     8    |    4mo    |
		|     9    |    burn   |
	 `
*/
func calculateAvailableAt(progress uint8, level int32) *time.Time {
	// card is burned
	if progress >= 9 {
		return nil
	}

	var (
		availableAt time.Time
		now         = time.Now()
	)
	switch progress {
	case 1, 2:
		timeScale := 4
		if level <= 2 {
			timeScale = 2
		}
		availableAt = now.Add(time.Hour * time.Duration(timeScale*int(progress)))
	case 3, 4:
		availableAt = now.AddDate(0, 0, int(progress)-2)
	case 5, 6:
		availableAt = now.AddDate(0, 0, 7*(int(progress)-4))
	case 7, 8:
		availableAt = now.AddDate(0, 1, 0)
	}

	availableAt = availableAt.Truncate(time.Hour)
	return &availableAt
}
