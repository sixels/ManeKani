package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"sixels.io/manekani/core/domain/errors"
)

func (api *AuthApi) Login(c echo.Context) error {
	state, err := randomString()
	if err != nil {
		return errors.Unknown(err)
	}
	nonce, err := randomString()
	if err != nil {
		return errors.Unknown(err)
	}

	sess, _ := session.Get("manekani-auth", c)
	sess.Options = &sessions.Options{
		Path:     "/auth/callback",
		MaxAge:   int((20 * time.Minute).Seconds()),
		Secure:   c.IsTLS(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	sess.Values["oauthstate"] = state
	sess.Values["oauthnonce"] = nonce

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return errors.Unknown(err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, api.AuthCodeURL(state, oidc.Nonce(nonce)))
}

func randomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
