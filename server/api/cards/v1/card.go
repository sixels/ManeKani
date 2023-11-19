package cards

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sixels/manekani/server/api/apicommon"
	"net/http"

	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/server/api/cards/util"
)

func (a *CardsApiV1) AllUserCards(c echo.Context) error {
	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	var filters cards.QueryManyCardsRequest
	if err := c.Bind(&filters); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	allCards, err := a.Cards.AllCards(
		c.Request().Context(), userID, filters,
	)
	if err != nil {
		log.Errorf("query user allCards error: %w", err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusOK, allCards))
}
