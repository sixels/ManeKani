package cards

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/server/api/cards/util"
)

func (api *CardsApiV1) AllUserCards() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		var filters cards.QueryManyCardsRequest
		if err := c.BindQuery(&filters); err != nil {
			c.Error(err)
			return
		}

		cards, err := api.Cards.AllCards(
			c.Request.Context(), userID, filters,
		)
		if err != nil {
			c.Error(fmt.Errorf("query user cards error: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, cards)
	}
}
