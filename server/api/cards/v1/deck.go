package cards

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/server/api/cards/util"
)

func (api *CardsApiV1) QueryDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("deck not found"))
			return
		}

		ctx := c.Request.Context()
		queried, err := api.Cards.QueryDeck(ctx, id)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, queried)
		}
	}
}

func (api *CardsApiV1) AllDecks() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := new(cards.QueryManyDecksRequest)
		if err := c.BindQuery(filters); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		decks, err := api.Cards.AllDecks(ctx, *filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, decks)
		}
	}
}

func (api *CardsApiV1) SubscribeUserToDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		deckID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(fmt.Errorf("subscribe: invalid deck ID"))
			c.Status(http.StatusBadRequest)
		}

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		if err := api.Cards.AddDeckSubscriber(
			c.Request.Context(), deckID, userID,
		); err != nil {
			c.Error(fmt.Errorf("subscribe error: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}

func (api *CardsApiV1) UnsubscribeUserFromDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		deckID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(fmt.Errorf("subscribe: invalid deck ID"))
			c.Status(http.StatusBadRequest)
		}

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		if err := api.Cards.RemoveDeckSubscriber(
			c.Request.Context(), deckID, userID,
		); err != nil {
			c.Error(fmt.Errorf("unsubscribe error: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
