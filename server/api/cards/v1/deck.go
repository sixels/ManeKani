package cards

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"net/http"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/server/api/apicommon"
	"github.com/sixels/manekani/server/api/cards/util"
)

func (a *CardsApiV1) GetDecks(c echo.Context, params GetDecksParams) error {
	log.Debugf("getting decks. params: %v", params)
	fmt.Println("abbdsad")
	reqFilters := cards.QueryManyDecksRequest{
		FilterPagination: filters.FilterPagination{
			Page: params.Page,
		},
		FilterIDs: filters.FilterIDs{
			IDs: (*filters.CommaSeparatedUUID)(params.Ids),
		},
		FilterSubjects: filters.FilterSubjects{
			Subjects: (*filters.CommaSeparatedUUID)(params.Subjects),
		},
		FilterOwners: filters.FilterOwners{
			Owners: (*filters.CommaSeparatedString)(params.Owners),
		},
		FilterNames: filters.FilterNames{
			Names: (*filters.CommaSeparatedString)(params.Names),
		},
	}

	ctx := c.Request().Context()
	decks, err := a.Cards.AllDecks(ctx, reqFilters)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusOK, decks))
}

func (a *CardsApiV1) CreateDeck(c echo.Context) error {
	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	var req DeckCreateRequest
	if err := c.Bind(&req); err != nil {
		log.Error("create-subject bind error: %w", err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	deck, err := a.Cards.CreateDeck(c.Request().Context(), userID, cards.CreateDeckRequest{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusCreated, deck))
}

func (a *CardsApiV1) GetDeck(c echo.Context, id string) error {
	deckID, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, errors.NotFound("deck not found"))
	}
	ctx := c.Request().Context()
	queried, err := a.Cards.QueryDeck(ctx, deckID)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}
	return apicommon.Respond(c, apicommon.Response(http.StatusOK, queried))
}

func (a *CardsApiV1) SubscribeUserToDeck(c echo.Context) error {
	deckID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error("subscribe: invalid deck ID")
		return apicommon.Error(http.StatusBadRequest, err)
	}

	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	if err := a.Cards.AddDeckSubscriber(
		c.Request().Context(), deckID, userID,
	); err != nil {
		log.Error("subscribe error: %w", err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
}

func (a *CardsApiV1) UnsubscribeUserFromDeck(c echo.Context) error {
	deckID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error("subscribe: invalid deck ID")
		return apicommon.Error(http.StatusBadRequest, err)
	}

	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	if err := a.Cards.RemoveDeckSubscriber(
		c.Request().Context(), deckID, userID,
	); err != nil {
		log.Error("unsubscribe error: %w", err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
}
