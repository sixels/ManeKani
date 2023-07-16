package tokens

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/server/api/apicommon"
)

func (api *TokenApi) GetTokens(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		apicommon.Respond(
			c,
			apicommon.Error(
				http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)),
			),
		)
		return
	}

	tokens, err := api.tokens.QueryTokens(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	apicommon.Respond(c, apicommon.Response(http.StatusOK, tokens))
}

func (api *TokenApi) CreateToken(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		apicommon.Respond(
			c,
			apicommon.Error(
				http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)),
			),
		)
		return
	}

	var req tokens.GenerateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
		return
	}

	token, err := api.tokens.CreateToken(c.Request.Context(), userID, req)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	apicommon.Respond(c, apicommon.Response(http.StatusCreated, token))
}

func (api *TokenApi) DeleteToken(c *gin.Context, id string) {
	userID := c.GetString("userID")
	if userID == "" {
		c.Status(http.StatusForbidden)
		return
	}

	tokenID, err := ulid.Parse(id)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	if err := api.tokens.DeleteToken(c.Request.Context(), userID, tokenID); err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func (api *TokenApi) ValidateToken(c *gin.Context, params ValidateTokenParams) {
	log.Println(params.Authorization, "<-")
	token := strings.TrimPrefix(params.Authorization, "Bearer ")

	tk, err := ValidateToken(c.Request.Context(), api, token)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
		return
	}

	print("user token validated")
	c.JSON(http.StatusOK, ValidateTokenResponse{
		Subject: tk.UserID,
	})
}

// func (api *TokenApi) QueryTokens() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userID := c.GetString("userID")
// 		if userID == "" {
// 			apicommon.Respond(
// 				c,
// 				apicommon.Error(
// 					http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)),
// 				),
// 			)
// 			return
// 		}
// 		tokens, err := api.tokens.QueryTokens(c.Request.Context(), userID)
// 		if err != nil {
// 			c.Error(err)
// 			apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
// 			return
// 		}
// 		apicommon.Respond(c, apicommon.Response(http.StatusOK, tokens))
// 	}
// }

// func (api *TokenApi) CreateToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userID := c.GetString("userID")
// 		if userID == "" {
// 			apicommon.Respond(
// 				c,
// 				apicommon.Error(
// 					http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)),
// 				),
// 			)
// 			return
// 		}
// 		var req tokens.GenerateTokenRequest
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.Error(err)
// 			apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
// 			return
// 		}
// 		token, err := api.tokens.CreateToken(c.Request.Context(), userID, req)
// 		if err != nil {
// 			c.Error(err)
// 			apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
// 			return
// 		}
// 		apicommon.Respond(c, apicommon.Response(http.StatusCreated, token))
// 	}
// }

// func (api *TokenApi) DeleteToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userID := c.GetString("userID")
// 		if userID == "" {
// 			c.Status(http.StatusForbidden)
// 			return
// 		}
// 		tokenID, err := ulid.Parse(c.Param("id"))
// 		if err != nil {
// 			c.Error(err)
// 			c.Status(http.StatusBadRequest)
// 			return
// 		}
// 		if err := api.tokens.DeleteToken(c.Request.Context(), userID, tokenID); err != nil {
// 			c.Error(err)
// 			c.Status(http.StatusInternalServerError)
// 			return
// 		}
// 		c.Status(http.StatusNoContent)
// 	}
// }

type ValidateTokenResponse struct {
	Subject string   `json:"subject"`
	Extra   struct{} `json:"extra"`
}

// func (api *TokenApi) ValidateToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token, err := auth.GetAuthTokenHeader(c.Request)
// 		if err != nil {
// 			c.Error(err)
// 			apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
// 			return
// 		}

// 		tk, err := ValidateToken(c.Request.Context(), api, token)
// 		if err != nil {
// 			c.Error(err)
// 			apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
// 			return
// 		}

// 		print("user token validated")
// 		c.JSON(http.StatusOK, ValidateTokenResponse{
// 			Subject: tk.UserID,
// 		})
// 	}
// }

func ValidateToken(ctx context.Context, api *TokenApi, token string) (tokens.UserToken, error) {
	tk, err := api.tokens.GetToken(ctx, token)
	if err != nil {
		return tokens.UserToken{}, err
	}

	if tk.Status != tokens.TokenStatusActive {
		return tokens.UserToken{}, errors.New("token is not active")
	}

	return tk, nil
}
