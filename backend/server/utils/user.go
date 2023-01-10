package utils

import (
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"sixels.io/manekani/services/auth"
)

func GetCurrentUser(c *gin.Context, authenticator *auth.Authenticator) *oidc.UserInfo {
	userSession := sessions.DefaultMany(c, "user-session")
	authToken, ok := userSession.Get("AuthToken").(oauth2.Token)

	if !ok || !authToken.Valid() {
		return nil
	}

	ctx := c.Request.Context()

	user, err := authenticator.GetUserInfo(ctx, &authToken)
	if err != nil {
		log.Printf("user info error: %v\n", err)
		return nil
	}

	return user
}
