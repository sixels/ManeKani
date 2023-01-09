package auth

import (
	"fmt"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func (api *AuthApi) SetupRoutes(router *echo.Echo) {
	fmt.Printf("TODO")

	authRoute := router.Group("/auth")
	keycloakRoute := authRoute.Group("/kc")

	keycloakUrl, err := url.Parse(os.Getenv("KEYCLOAK_URL"))
	if err != nil {
		panic("KEYCLOAK_URL environment variable is not valid")
	}

	keycloakRoute.Use(middleware.Proxy(middleware.NewRandomBalancer(
		[]*middleware.ProxyTarget{{URL: keycloakUrl}})))

	authRoute.GET("/callback", api.OAuthCallback)
	router.GET("/login", api.Login)
}
