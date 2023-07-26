package cards

import (
	"fmt"
	"github.com/sixels/manekani/server/api/apicommon"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/server/api/cards/util"
)

// AllUserReviews gets all reviews for a user
func (api *CardsApiV1) AllUserReviews() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
			return
		}

		var filters cards.QueryManyReviewsRequest
		if err := c.BindQuery(&filters); err != nil {
			c.Error(err)
			apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
			return
		}

		cards, err := api.Cards.AllReviews(
			c.Request.Context(), userID, filters,
		)
		if err != nil {
			c.Error(fmt.Errorf("query user reviews error: %w", err))
			apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
			return
		}

		apicommon.Respond(c, apicommon.Response(http.StatusOK, cards))
	}
}

// CreateReview creates a new review
func (api *CardsApiV1) CreateReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
			return
		}

		var req cards.CreateReviewAPIRequest
		if err := c.Bind(&req); err != nil {
			c.Error(fmt.Errorf("create review bind error: %w", err))
			apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
			return
		}

		review, err := api.Cards.CreateReview(c.Request.Context(), userID, req)
		if err != nil {
			c.Error(err)
			apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
			return
		}

		apicommon.Respond(c, apicommon.Response(http.StatusCreated, review))
	}
}
