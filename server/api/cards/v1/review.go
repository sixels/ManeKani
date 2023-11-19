package cards

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sixels/manekani/server/api/apicommon"
	"net/http"

	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/server/api/cards/util"
)

// AllUserReviews gets all reviews for a user
func (a *CardsApiV1) AllUserReviews(c echo.Context) error {
	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	var filters cards.QueryManyReviewsRequest
	if err := c.Bind(&filters); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	reviews, err := a.Cards.AllReviews(
		c.Request().Context(), userID, filters,
	)
	if err != nil {
		log.Errorf("query user reviews error: %w", err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusOK, reviews))
}

// CreateReview creates a new review
func (a *CardsApiV1) CreateReview(c echo.Context) error {
	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	var req cards.CreateReviewAPIRequest
	if err := c.Bind(&req); err != nil {
		log.Errorf("create review bind error: %w", err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	review, err := a.Cards.CreateReview(c.Request().Context(), userID, req)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusCreated, review))
}
