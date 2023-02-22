package cards

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/cards/filters"
	user_domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/deckprogress"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/review"
	"sixels.io/manekani/ent/schema"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *CardsRepository) CreateReview(ctx context.Context, userID string, req *cards.CreateReviewRequest) (*cards.Review, error) {
	return util.WithTx(ctx, repo.client.Client, func(tx *ent.Tx) (*cards.Review, error) {
		reqCard, err := tx.Card.Query().
			Where(card.IDEQ(*req.Card.Only())).
			WithSubject(func(sq *ent.SubjectQuery) { sq.Select(subject.FieldLevel) }).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		// if it is a lesson session, check if the card was already started
		if req.SessionType == cards.SessionLesson && reqCard.StartedAt != nil {
			return nil, fmt.Errorf("card is not available for lesson")
		}
		// check if AvailableAt is valid
		if reqCard.AvailableAt.Compare(now) > 0 {
			return nil, fmt.Errorf("card is not available yet")
		}

		var (
			startProgress      = reqCard.Progress
			totalErrors   uint = 0
		)
		for _, errQty := range req.Errors {
			totalErrors += uint(errQty)
		}

		endProgress := calculateNextProgress(req.SessionType, startProgress, totalErrors)

		createdReview, err := tx.Review.Create().
			SetCardID(reqCard.ID).
			SetErrors(req.Errors).
			SetStartProgress(startProgress).
			SetEndProgress(endProgress).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		createdReview.Edges.Card = reqCard

		updateCard := reqCard.Update().
			SetProgress(endProgress).
			SetTotalErrors(reqCard.TotalErrors + int32(totalErrors))

		if reqCard.StartedAt == nil {
			updateCard.SetStartedAt(now)
		}
		if endProgress == 5 && reqCard.PassedAt == nil {
			updateCard.SetPassedAt(now)
		}

		availableAt := calculateAvailableAt(endProgress, reqCard.Edges.Subject.Level)
		if availableAt != nil {
			updateCard.SetAvailableAt(*availableAt)
		} else {
			updateCard.ClearAvailableAt()
			updateCard.SetBurnedAt(now)
		}

		if err := updateCard.Exec(ctx); err != nil {
			return nil, err
		}

		if err := tx.User.UpdateOneID(userID).
			AppendPendingActions([]schema.PendingAction{
				{
					Action:   user_domain.ActionCheckCardUnlocks,
					Required: true,
					Metadata: user_domain.CheckCardUnlocksMeta{
						Card: reqCard.ID,
					},
				},
			}).
			Exec(ctx); err != nil {
			return nil, err
		}

		return ReviewFromEnt(createdReview), nil
	})
}

func (repo *CardsRepository) AllUserReviews(ctx context.Context, userID string, req *cards.QueryManyReviewsRequest) ([]*cards.Review, error) {
	fs := filters.NewFilter([]predicate.Review{
		review.HasCardWith(
			card.HasDeckProgressWith(
				deckprogress.HasUserWith(user.IDEQ(userID)))),
	})

	filters.In(fs, req.IDs.Separate(), review.IDIn)
	filters.With(fs, req.CreatedAfter, review.CreatedAtGTE)
	filters.With(fs, req.CreatedBefore, review.CreatedAtLTE)
	filters.In(fs, req.Cards.Separate(), func(ids ...uuid.UUID) predicate.Review {
		return review.HasCardWith(card.IDIn(ids...))
	})
	filters.With(fs, req.Passed, func(should_pass bool) predicate.Review {
		passed := func(s *sql.Selector) {
			s.Where(sql.ColumnsGT(review.FieldStartProgress, review.FieldEndProgress))
		}
		if !should_pass {
			return review.Not(passed)
		}
		return passed
	})

	page := 0
	if req.Page != nil {
		page = int(*req.Page)
	}

	queried, err := repo.client.Review.Query().
		Limit(500).
		Offset(page).
		Where(review.And(fs.Filters()...)).
		WithCard(func(cq *ent.CardQuery) {
			cq.Select(card.FieldID)
		}).
		Order(ent.Asc(review.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(queried, ReviewFromEnt), nil
}

func ReviewFromEnt(e *ent.Review) *cards.Review {
	return &cards.Review{
		ID:            e.ID,
		CreatedAt:     e.CreatedAt,
		Errors:        e.Errors,
		StartProgress: e.StartProgress,
		EndProgress:   e.EndProgress,
		Card:          e.Edges.Card.ID,
	}
}

/*
Calculate the next progress based on the amount of errors
if there are some error, the next progress will be:

	endProgress := startProgress - (math.Ceil(errors / penaltyScale) * penaltyScale)

Otherwise, it will just add one to the starting progress
*/
func calculateNextProgress(sessionType cards.SessionType, startProgress uint8, errors uint) (endProgress uint8) {
	var penaltyScale uint8 = 1
	if startProgress >= 5 {
		penaltyScale = 2
	}

	if errors == 0 || sessionType == cards.SessionLesson {
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
	return &availableAt
}
