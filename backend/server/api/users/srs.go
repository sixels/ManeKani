package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/srs"
	"sixels.io/manekani/core/domain/user"
)

func (api *UserApi) GetSRSInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxUser, ok := c.Get("user")
		if !ok {
			c.Error(fmt.Errorf("user is not set in the context"))
			c.Status(http.StatusUnauthorized)
			return
		}
		userData := ctxUser.(*user.User)

		deckID, err := uuid.Parse(c.Param("deck-id"))
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		userCards, err := api.users.GetUserCards(c.Request.Context(), userData.ID, deckID)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, srs.SRSUserData{
			Cards: userCards,
		})
	}
}

func (api *UserApi) UnsubscribeUserFromDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		deckID := uuid.MustParse(c.Param("deck-id"))

		ctxUser, ok := c.Get("user")
		if !ok {
			c.Error(fmt.Errorf("user is not set in the context"))
			c.Status(http.StatusUnauthorized)
			return
		}
		userData := ctxUser.(*user.User)

		err := api.users.UnsubscribeUserFromDeck(c.Request.Context(), userData.ID, deckID)
		if err != nil {
			c.Error(fmt.Errorf("reset error: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
