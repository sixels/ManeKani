package cards

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/server/api/cards/util"
)

func (api *CardsApiV1) AllUserReviews() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		filters := new(cards.QueryManyReviewsRequest)
		if err := c.BindQuery(filters); err != nil {
			c.Error(err)
			return
		}

		cards, err := api.Cards.AllUserReviews(
			c.Request.Context(), userID, filters,
		)
		if err != nil {
			c.Error(fmt.Errorf("query user reviews error: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, cards)

	}
}

func (api *CardsApiV1) CreateReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		var req cards.CreateReviewRequest
		if err := c.Bind(&req); err != nil {
			c.Error(fmt.Errorf("create review bind error: %w", err))
			c.Status(http.StatusBadRequest)
			return
		}

		review, err := api.Cards.CreateReview(c.Request.Context(), userID, &req)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}

		if err := api.Users.ResolveActions(c.Request.Context(), userID); err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
		}

		c.JSON(http.StatusCreated, review)
	}
}
