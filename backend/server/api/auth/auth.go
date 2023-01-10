package auth

import (
	"github.com/gin-gonic/gin"
	"sixels.io/manekani/services/auth"
)

type AuthApi struct {
	*auth.Authenticator
}

func New(auth *auth.Authenticator) *AuthApi {
	return &AuthApi{
		Authenticator: auth,
	}
}

func (api *AuthApi) SetupRoutes(router *gin.Engine) {
	router.GET("/auth/callback", api.OAuthCallback())
	router.GET("/login", api.Login())
}
