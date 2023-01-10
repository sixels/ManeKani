package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"sixels.io/manekani/services/auth"
)

type OpenIdAuthConfig struct {
	Authenticator *auth.Authenticator
}

func LoginRequired(authenticator auth.Authenticator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			sess, _ := session.Get("manekani-profile", c)
			if token, ok := sess.Values["AuthToken"].(auth.StaticToken); !ok || !auth.ReviveToken(token).Valid() {
				return c.NoContent(http.StatusUnauthorized)
			}
			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}
