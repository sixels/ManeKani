package users

//
//import (
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/gommon/log"
//	"net/http"
//
//	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
//	"github.com/supertokens/supertokens-golang/recipe/session"
//	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
//)
//
//func (api *UserApi) RequiresUser(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		sessionRequired := true
//		currentSession, err := session.GetSession(c.Request(), c.Response().Writer, &sessmodels.VerifySessionOptions{SessionRequired: &sessionRequired})
//		if err != nil {
//			log.Error(err)
//			return c.NoContent(http.StatusUnauthorized)
//		}
//
//		authUser, err := emailpassword.GetUserByID(currentSession.GetUserID())
//		if err != nil {
//			log.Error(err)
//			return c.NoContent(http.StatusNotAcceptable)
//		}
//
//		userInfo, err := api.users.QueryUser(c.Request().Context(), authUser.ID)
//		if err != nil {
//			log.Error(err)
//			return c.NoContent(http.StatusInternalServerError)
//		}
//
//		c.Set("user", userInfo)
//		return next(c)
//	}
//}
