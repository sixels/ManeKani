package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sixels.io/manekani/core/domain/user"
)

func (api *UserApi) GetBasicUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxUser, ok := c.Get("user")
		if !ok {
			c.Status(http.StatusUnauthorized)
		}

		userInfo, err := api.users.QueryUser(
			c.Request.Context(), ctxUser.(*user.User).ID,
		)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, user.UserBasic{
			Email:    userInfo.Email,
			Username: userInfo.Username,
			Decks:    userInfo.Decks,
		})
	}
}
