package utils

import (
	"fmt"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"sixels.io/manekani/services/auth"
)

func GetCurrentUser(c echo.Context, authenticator *auth.Authenticator) *oidc.UserInfo {
	profileSession, _ := session.Get("manekani-profile", c)
	staticToken, ok := profileSession.Values["AuthToken"].(auth.StaticToken)
	authToken := auth.ReviveToken(staticToken)

	if !ok || !authToken.Valid() {
		return nil
	}

	ctx := c.Request().Context()

	user, err := authenticator.GetUserInfo(ctx, authToken)
	if err != nil {
		log.Printf("user info error: %v\n", err)
		return nil
	}
	fmt.Printf("%v\n", user)

	return user
}
