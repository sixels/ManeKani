package cards

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/cards/filters"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/deckprogress"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/review"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *CardsRepository) CreateReview(ctx context.Context, userID string, req cards.CreateReviewRequest) (*cards.Review, error) {
	reqCard, err := repo.client.CardClient().Query().
		Where(card.IDEQ(req.CardID)).
		WithSubject(func(sq *ent.SubjectQuery) { sq.Select(subject.FieldLevel) }).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	createdReview, err := repo.client.ReviewClient().Create().
		SetCardID(reqCard.ID).
		SetErrors(req.Errors).
		SetStartProgress(req.StartProgress).
		SetEndProgress(req.EndProgress).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	createdReview.Edges.Card = reqCard

	return util.Ptr(ReviewFromEnt(createdReview)), nil
}

func (repo *CardsRepository) AllReviews(ctx context.Context, userID string, req cards.QueryManyReviewsRequest) ([]cards.Review, error) {
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

	queried, err := repo.client.ReviewClient().Query().
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

func ReviewFromEnt(e *ent.Review) cards.Review {
	return cards.Review{
		ID:            e.ID,
		CreatedAt:     e.CreatedAt,
		Errors:        e.Errors,
		StartProgress: e.StartProgress,
		EndProgress:   e.EndProgress,
		Card:          e.Edges.Card.ID,
	}
}
