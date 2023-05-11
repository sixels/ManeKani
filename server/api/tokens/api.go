package tokens

import (
	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/adapters/tokens"
	"github.com/sixels/manekani/server/auth"
)

type TokenApi struct {
	tokens *tokens.TokensAdapter
	auth   *auth.AuthService
}

func New(repo *tokens.TokensAdapter, auth *auth.AuthService) *TokenApi {
	return &TokenApi{tokens: repo, auth: auth}
}

func (api *TokenApi) SetupRoutes(router *gin.Engine) {
	handler := router.Group("/api/token")

	handler.GET("/", api.auth.EnsureLogin(), api.QueryTokens())
	handler.POST("/", api.auth.EnsureLogin(), api.CreateToken())
	handler.DELETE("/:id", api.auth.EnsureLogin(), api.DeleteToken())
}
