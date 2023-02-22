package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
)

func (api *UserApi) RequiresUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionRequired := true
		currentSession, err := session.GetSession(c.Request, c.Writer, &sessmodels.VerifySessionOptions{SessionRequired: &sessionRequired})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		authUser, err := emailpassword.GetUserByID(currentSession.GetUserID())
		if err != nil {
			c.AbortWithError(http.StatusNotAcceptable, err)
			return
		}

		userInfo, err := api.users.QueryUser(c.Request.Context(), authUser.ID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("user", userInfo)
	}
}
