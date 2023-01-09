package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"sixels.io/manekani/core/domain/errors"
)

func (api *AuthApi) OAuthCallback(c echo.Context) error {
	state := c.QueryParam("state")

	authSession, _ := session.Get("manekani-auth", c)

	cookieState, hasState := authSession.Values["oauthstate"].(string)
	cookieNonce, hasNonce := authSession.Values["oauthnonce"].(string)

	println(cookieState, cookieNonce)

	if !hasState || !hasNonce || state != cookieState {
		return errors.InvalidRequest("invalid auth state")
	}

	ctx := c.Request().Context()

	code := c.QueryParam("code")
	log.Println("code:", code)

	oauthToken, err := api.Exchange(ctx, code)
	if err != nil {
		log.Println("A")
		return errors.Unknown(fmt.Errorf("token exchange failed: %w", err))
	}

	idToken, err := api.VerifyIDToken(ctx, oauthToken)
	if err != nil {
		log.Println("B")
		return errors.Unknown(fmt.Errorf("validation falied: %w", err))
	}

	if idToken.Nonce != cookieNonce {
		log.Println("C")
		return errors.InvalidRequest("invalid state nonce")
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		log.Println("D")
		return errors.Unknown(err)
	}

	profileSession, _ := session.Get("manekani-profile", c)
	profileSession.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
		Secure:   c.IsTLS(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	log.Printf("%#v", profile)

	profileSession.Values["manekani-acctoken"] = oauthToken.AccessToken

	if err := profileSession.Save(c.Request(), c.Response().Writer); err != nil {
		log.Printf("E: %v\n", err)
		return errors.Unknown(err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/user")
}
