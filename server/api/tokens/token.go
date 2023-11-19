package tokens

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"

	"github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/server/api/apicommon"
)

func (api *TokenApi) GetTokens(c echo.Context) error {
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
		return apicommon.Error(
			http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)),
		)
	}

	queryTokens, err := api.tokens.QueryTokens(c.Request().Context(), userID)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusOK, queryTokens))
}

func (api *TokenApi) CreateToken(c echo.Context) error {
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
		return apicommon.Error(
			http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)),
		)
	}

	var req tokens.GenerateTokenRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	token, err := api.tokens.CreateToken(c.Request().Context(), userID, req)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusCreated, token))
}

func (api *TokenApi) DeleteToken(c echo.Context, id string) error {
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
		return apicommon.Error(http.StatusForbidden, errors.New(http.StatusText(http.StatusUnauthorized)))
	}

	tokenID, err := ulid.Parse(id)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, errors.New("invalid token id"))
	}

	if err := api.tokens.DeleteToken(c.Request().Context(), userID, tokenID); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}
	return apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
}

func (api *TokenApi) ValidateToken(c echo.Context, params ValidateTokenParams) error {
	token := strings.TrimPrefix(params.Authorization, "Bearer ")

	tk, err := ValidateToken(c.Request().Context(), api, token)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	print("user token validated")
	return c.JSON(http.StatusOK, ValidateTokenResponse{
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
// 			log.Error(err)
// 			return apicommon.Error(http.StatusInternalServerError, err)
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
// 			log.Error(err)
// 			return apicommon.Error(http.StatusUnauthorized, err)
// 			return
// 		}
// 		token, err := api.tokens.CreateToken(c.Request.Context(), userID, req)
// 		if err != nil {
// 			log.Error(err)
// 			return apicommon.Error(http.StatusInternalServerError, err)
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
// 			log.Error(err)
// 			c.Status(http.StatusBadRequest)
// 			return
// 		}
// 		if err := api.tokens.DeleteToken(c.Request.Context(), userID, tokenID); err != nil {
// 			log.Error(err)
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
// 			log.Error(err)
// 			return apicommon.Error(http.StatusBadRequest, err)
// 			return
// 		}

// 		tk, err := ValidateToken(c.Request.Context(), api, token)
// 		if err != nil {
// 			log.Error(err)
// 			return apicommon.Error(http.StatusUnauthorized, err)
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
