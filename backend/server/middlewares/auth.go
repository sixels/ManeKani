package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"sixels.io/manekani/services/auth"
)

func LoginRequired(authenticator auth.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.DefaultMany(c, "user-session")
		if token, ok := session.Get("AuthToken").(oauth2.Token); !ok || !token.Valid() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
