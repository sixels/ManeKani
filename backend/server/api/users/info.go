package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"sixels.io/manekani/core/domain/user"
)

func (api *UserApi) GetBasicUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionRequired := false
		currentSession, err := session.GetSession(c.Request, c.Writer, &sessmodels.VerifySessionOptions{SessionRequired: &sessionRequired})
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		if currentSession == nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		authUser, err := emailpassword.GetUserByID(currentSession.GetUserID())
		if err != nil {
			c.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}
		userInfo, err := api.users.QueryUser(c.Request.Context(), authUser.ID)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, user.UserBasic{
			Email:    userInfo.Email,
			Username: userInfo.Username,
		})
	}
}
