package cards

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/server/api/apicommon"
	"github.com/sixels/manekani/server/api/cards/util"
)

func (api *CardsApiV1) GetDecks(c *gin.Context, params GetDecksParams) {
	filters := new(cards.QueryManyDecksRequest)
	if err := c.BindQuery(filters); err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
		return
	}

	ctx := c.Request.Context()
	decks, err := api.Cards.AllDecks(ctx, *filters)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}
	apicommon.Respond(c, apicommon.Response(http.StatusOK, decks))
}

func (api *CardsApiV1) CreateDeck(c *gin.Context) {
	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
		return
	}

	var req DeckCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.Error(fmt.Errorf("create-subject bind error: %w", err))
		apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
		return
	}

	deck, err := api.Cards.CreateDeck(c.Request.Context(), userID, cards.CreateDeckRequest{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	apicommon.Respond(c, apicommon.Response(http.StatusCreated, deck))
}

func (api *CardsApiV1) GetDeck(c *gin.Context, id string) {
	deckID, err := uuid.Parse(id)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, errors.NotFound("deck not found")))
		return
	}
	ctx := c.Request.Context()
	queried, err := api.Cards.QueryDeck(ctx, deckID)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}
	apicommon.Respond(c, apicommon.Response(http.StatusOK, queried))
}

func (api *CardsApiV1) SubscribeUserToDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		deckID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(fmt.Errorf("subscribe: invalid deck ID"))
			apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
			return
		}

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
			return
		}

		if err := api.Cards.AddDeckSubscriber(
			c.Request.Context(), deckID, userID,
		); err != nil {
			c.Error(fmt.Errorf("subscribe error: %w", err))
			apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
			return
		}

		apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
	}
}

func (api *CardsApiV1) UnsubscribeUserFromDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		deckID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(fmt.Errorf("subscribe: invalid deck ID"))
			apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
			return
		}

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
			return
		}

		if err := api.Cards.RemoveDeckSubscriber(
			c.Request.Context(), deckID, userID,
		); err != nil {
			c.Error(fmt.Errorf("unsubscribe error: %w", err))
			apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
			return
		}

		apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
	}
}
