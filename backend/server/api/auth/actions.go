package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (api *AuthApi) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := randomString()
		if err != nil {
			c.Error(err)
			return
		}
		nonce, err := randomString()
		if err != nil {
			c.Error(err)
			return
		}

		authSession := sessions.DefaultMany(c, "auth-session")
		authSession.Options(sessions.Options{
			Path:     "/auth/callback",
			MaxAge:   int((20 * time.Minute).Seconds()),
			Secure:   c.Request.TLS != nil,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

		authSession.Set("OauthState", state)
		authSession.Set("OauthNonce", nonce)

		if err := authSession.Save(); err != nil {
			c.Error(err)
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, api.AuthCodeURL(state, oidc.Nonce(nonce)))
	}
}

func randomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
