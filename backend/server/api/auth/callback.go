package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"sixels.io/manekani/core/domain/errors"
)

func (api *AuthApi) OAuthCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := c.Query("state")

		authSession := sessions.DefaultMany(c, "auth-session")

		cookieState, hasState := authSession.Get("OauthState").(string)
		cookieNonce, hasNonce := authSession.Get("OauthNonce").(string)

		println(cookieState, cookieNonce)

		if !hasState || !hasNonce || state != cookieState {
			e := errors.InvalidRequest("invalid auth state")
			c.Error(e)
			c.JSON(e.Status, e)
			return
		}

		ctx := c.Request.Context()
		authToken, err := api.Exchange(ctx, c.Query("code"))
		if err != nil {
			c.Error(fmt.Errorf("token exchange failed: %w", err))
			return
		}

		idToken, err := api.VerifyIDToken(ctx, authToken)
		if err != nil {
			c.Error(fmt.Errorf("validation falied: %w", err))
			return
		}

		if idToken.Nonce != cookieNonce {
			c.Error(errors.InvalidRequest("invalid state nonce"))
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			c.Error(errors.Unknown(fmt.Errorf("invalid token claims: %w", err)))
			return
		}

		profileSession := sessions.DefaultMany(c, "user-session")

		profileSession.Set("AuthToken", authToken)

		if err := profileSession.Save(); err != nil {
			c.Error(err)
			return
		}

		c.Redirect(http.StatusPermanentRedirect, "/user")
	}
}
