package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

		userCards, err := api.users.GetUserCards(c.Request.Context(), userData.Id)
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

func (api *UserApi) ResetSRSData() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxUser, ok := c.Get("user")
		if !ok {
			c.Error(fmt.Errorf("user is not set in the context"))
			c.Status(http.StatusUnauthorized)
			return
		}
		userData := ctxUser.(*user.User)

		err := api.users.ResetSRSData(c.Request.Context(), userData.Id)
		if err != nil {
			c.Error(fmt.Errorf("reset error: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
