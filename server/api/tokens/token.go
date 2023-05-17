package tokens

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/server/api/apicommon"
)

// QueryTokens godoc
//
//	@Id				get-token-query
//	@Summary		Get all user's API tokens
//	@Description	Get all user's API tokens
//	@Tags			user, tokens
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	apicommon.APIResponse[[]tokens.UserTokenPartial]
//	@Failure		401	{object}	apicommon.APIResponse[any]
//	@Failure		500	{object}	apicommon.APIResponse[any]
//	@Security		Login
//	@Router			/api/tokens [get]
func (api *TokenApi) QueryTokens() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// CreateToken godoc
//
//	@Id				post-token-create
//	@Summary		Create a new API token
//	@Description	Creates a new API token with the given permissions
//	@Tags			user, tokens
//	@Accept			json
//	@Produce		json
//	@Param			request	body		tokens.GenerateTokenRequest	true	"The API token options"
//	@Success		201		{object}	apicommon.APIResponse[string]
//	@Failure		400		{object}	apicommon.APIResponse[any]
//	@Failure		401		{object}	apicommon.APIResponse[any]
//	@Failure		500		{object}	apicommon.APIResponse[any]
//	@Security		Login
//	@Router			/api/tokens [post]
func (api *TokenApi) CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// DeleteToken godoc
//
//	@Id				delete-token-delete
//	@Summary		Delete an API token
//	@Description	Deletes an API token by the given ID
//	@Tags			user, tokens
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"The API token ID"
//	@Success		204	{object}	apicommon.APIResponse[any]
//	@Failure		400	{object}	apicommon.APIResponse[any]
//	@Failure		401	{object}	apicommon.APIResponse[any]
//	@Failure		500	{object}	apicommon.APIResponse[any]
//	@Security		Login
//	@Router			/api/tokens/{id} [delete]
func (api *TokenApi) DeleteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
			c.Status(http.StatusForbidden)
			return
		}

		tokenID, err := uuid.Parse(c.Param("id"))
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
}
