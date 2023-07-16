package tokens

import (
	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/adapters/tokens"
)

type TokenApi struct {
	tokens tokens.TokensAdapter
}

func (api *TokenApi) ServiceName() string {
	return "tokens"
}

func New(repo tokens.TokensAdapter) *TokenApi {
	return &TokenApi{tokens: repo}
}

func (api *TokenApi) SetupRoutes(router *gin.Engine) {
	RegisterHandlers(router, api)

	// handler := router.Group("/api/tokens")
	// router.GET("/auth/validate-token", api.ValidateToken())

	// handler.GET("/", api.QueryTokens())
	// handler.POST("/", api.CreateToken())
	// handler.DELETE("/:id", api.DeleteToken())

}
