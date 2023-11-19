package users

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"

	"github.com/sixels/manekani/core/domain/user"
)

func (api *UserApi) GetBasicUserInfo(c echo.Context) error {
	ctxUser, ok := c.Get("user").(*user.User)
	if !ok || ctxUser == nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	userInfo, err := api.users.QueryUser(
		c.Request().Context(), ctxUser.ID,
	)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user.UserBasic{
		Email:    userInfo.Email,
		Username: userInfo.Username,
		Decks:    userInfo.Decks,
	})
}
