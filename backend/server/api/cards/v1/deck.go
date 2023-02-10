package cards

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
)

func (api *CardsApi) QueryDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("deck not found"))
			return
		}

		ctx := c.Request.Context()
		queried, err := api.cards.QueryDeck(ctx, id)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, queried)
		}

	}
}

func (api *CardsApi) AllDecks() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := new(cards.QueryManyDecksRequest)
		if err := c.BindQuery(filters); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		decks, err := api.cards.AllDecks(ctx, *filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, decks)
		}
	}
}
