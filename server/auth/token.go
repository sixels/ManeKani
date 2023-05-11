package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ValidateTokenResponse struct {
	Subject string   `json:"subject"`
	Extra   struct{} `json:"extra"`
}

func (auth *AuthService) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		headerPrefix := "Bearer "

		if !strings.HasPrefix(token, headerPrefix) {
			c.Status(http.StatusUnauthorized)
			return
		}
		token = token[len(headerPrefix):]

		tk, err := auth.tokens.GetToken(c.Request.Context(), token)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.JSON(http.StatusOK, ValidateTokenResponse{
			Subject: tk.UserID,
		})
	}
}
