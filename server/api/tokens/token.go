package tokens

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/tokens"
)

func (api *TokenApi) QueryTokens() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
			c.Status(http.StatusForbidden)
			return
		}

		tokens, err := api.tokens.QueryTokens(c.Request.Context(), userID)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		c.JSON(http.StatusOK, tokens)
	}
}

func (api *TokenApi) CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
			c.Status(http.StatusForbidden)
			return
		}

		var caps tokens.APITokenCapabilities
		if err := c.ShouldBindJSON(&caps); err != nil {
			c.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}

		token, err := api.tokens.CreateToken(c.Request.Context(), userID, caps)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func (api *TokenApi) DeleteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
			c.Status(http.StatusForbidden)
			return
		}

		paramID := c.Param("id")
		if paramID == "" {
			c.Status(http.StatusBadRequest)
			return
		}
		tokenID, err := uuid.Parse(paramID)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusForbidden)
			return
		}

		if err := api.tokens.DeleteToken(c.Request.Context(), userID, tokenID); err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	}
}
