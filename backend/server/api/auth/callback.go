package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/services/auth"
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

	authToken, err := api.Exchange(ctx, c.QueryParam("code"))
	if err != nil {
		return errors.Unknown(fmt.Errorf("token exchange failed: %w", err))
	}

	idToken, err := api.VerifyIDToken(ctx, authToken)
	if err != nil {
		return errors.Unknown(fmt.Errorf("validation falied: %w", err))
	}

	if idToken.Nonce != cookieNonce {
		return errors.InvalidRequest("invalid state nonce")
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return errors.Unknown(fmt.Errorf("invalid token claims: %w", err))
	}

	profileSession, _ := session.Get("manekani-profile", c)
	profileSession.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
		Secure:   c.IsTLS(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	profileSession.Values["AuthToken"] = auth.ToStaticToken(authToken)

	if err := profileSession.Save(c.Request(), c.Response().Writer); err != nil {
		return errors.Unknown(err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/user")
}
