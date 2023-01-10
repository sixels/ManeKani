package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"sixels.io/manekani/services/auth"
)

func LoginRequired(authenticator auth.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.DefaultMany(c, "user-session")
		if token, ok := session.Get("AuthToken").(auth.StaticToken); !ok || !auth.ReviveToken(token).Valid() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
