package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/client-go"
	"github.com/sixels/manekani/core/adapters/tokens"
)

type AuthService struct {
	tokens *tokens.TokensAdapter
	ory    *ory.APIClient
}

func NewAuthService(tokenProvider *tokens.TokensAdapter) *AuthService {
	oryConfig := ory.NewConfiguration()
	oryConfig.Servers = ory.ServerConfigurations{{URL: KRATOS_ADMIN_API_URL}}
	oryConfig.Host = strings.TrimPrefix(KRATOS_API_URL, "http://")
	ory := ory.NewAPIClient(oryConfig)

	return &AuthService{
		tokens: tokenProvider,
		ory:    ory,
	}
}

func (auth *AuthService) SetupRoutes(router *gin.Engine) {
	router.GET("/auth/tk/validate", auth.ValidateToken())

	// hookHandler := router.Group("/hook/user")
	// hookHandler.POST("/register", hooks.RegisterUser(api.users))

	// router.GET("/user", api.RequiresUser(), api.GetBasicUserInfo())

	// router.GET("/user/decks", api.RequiresUser(), api.GetSRSInfo())
}
