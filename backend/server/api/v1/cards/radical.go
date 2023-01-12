package v1

import (
	"net/http"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"

	"github.com/gin-gonic/gin"
)

// CreateRadical godoc
// @Id post-radical-create
// @Summary Create a new radical
// @Description Creates a radical with the given values
// @Tags cards, radical
// @Accept mpfd
// @Produce json
// @Param radical body cards.CreateRadicalRequest true "The radical to be created"
// @Success 201 {object} cards.Radical
// @Router /api/v1/radical [post]
func (api *CardsApi) CreateRadical() gin.HandlerFunc {
	return func(c *gin.Context) {
		radical := new(cards.CreateRadicalRequest)
		if err := c.Bind(radical); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		created, err := api.cards.CreateRadical(ctx, *radical)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusCreated, created)
		}
	}
}

// QueryRadical godoc
// @Id get-radical-query
// @Summary Query a radical
// @Description Search a radical by its name
// @Tags cards, radical
// @Accept */*
// @Produce json
// @Param name path string true "Radical name"
// @Success 200 {object} cards.Radical
// @Router /api/v1/radical/{name} [get]
func (api *CardsApi) QueryRadical() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		ctx := c.Request.Context()
		queried, err := api.cards.QueryRadical(ctx, name)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, queried)
		}
	}
}

// UpdateRadical godoc
// @Id patch-radical-update
// @Summary Update a radical
// @Description Update a radical with the given values
// @Tags cards, radical
// @Accept json
// @Produce json
// @Param name path string true "Radical name"
// @Param radical body cards.UpdateRadicalRequest true "Radical fields to update"
// @Success 200 {object} cards.Radical
// @Router /api/v1/radical/{name} [patch]
func (api *CardsApi) UpdateRadical() gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.Param("name")

		radical := new(cards.UpdateRadicalRequest)
		if err := c.Bind(radical); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		updated, err := api.cards.UpdateRadical(ctx, name, *radical)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, updated)
		}
	}

}

// DeleteRadical godoc
// @Id delete-radical-delete
// @Summary Delete a radical
// @Description Delete a radical by its name
// @Tags cards, radical
// @Accept */*
// @Produce json
// @Param name path string true "Radical name"
// @Success 200
// @Router /api/v1/radical/{name} [delete]
func (api *CardsApi) DeleteRadical() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		ctx := c.Request.Context()
		if err := api.cards.DeleteRadical(ctx, name); err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}

// AllRadicals godoc
// @Id get-radical-all
// @Summary Query all radicals
// @Description Return a list of all radicals
// @Tags cards, radical
// @Accept */*
// @Produce json
// @Success 200 {array} cards.PartialRadicalResponse
// @Router /api/v1/radical [get]
func (api *CardsApi) AllRadicals() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := new(cards.QueryAllRadicalRequest)
		if err := c.BindQuery(filters); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		radicals, err := api.cards.AllRadicals(ctx, *filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, radicals)
		}
	}
}

// QueryRadicalKanjis godoc
// @Id get-radical-kanji
// @Summary Query kanjis from radical
// @Description Return a list of all kanji that are composed by the given radical
// @Tags cards, radical
// @Accept */*
// @Produce json
// @Param name path string true "Radical name"
// @Success 200 {array} cards.PartialKanjiResponse
// @Router /api/v1/radical/{:name}/kanji [get]
func (api *CardsApi) QueryRadicalKanjis() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		ctx := c.Request.Context()
		kanjis, err := api.cards.QueryRadicalKanjis(ctx, name)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, kanjis)
		}
	}
}
