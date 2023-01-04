package v1

import (
	"log"
	"net/http"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"

	"github.com/labstack/echo/v4"
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
func (api *CardsApi) CreateRadical(c echo.Context) error {
	radical := new(cards.CreateRadicalRequest)
	if err := c.Bind(radical); err != nil {
		return err
	}
	log.Println("route: ", radical)

	ctx := c.Request().Context()
	created, err := api.cards.CreateRadical(ctx, *radical)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusCreated, created)
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
func (api *CardsApi) QueryRadical(c echo.Context) error {
	name := c.Param("name")

	ctx := c.Request().Context()
	queried, err := api.cards.QueryRadical(ctx, name)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, queried)
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
func (api *CardsApi) UpdateRadical(c echo.Context) error {
	name := c.Param("name")

	radical := new(cards.UpdateRadicalRequest)
	if err := c.Bind(radical); err != nil {
		return err
	}

	ctx := c.Request().Context()
	updated, err := api.cards.UpdateRadical(ctx, name, *radical)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, updated)
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
func (api *CardsApi) DeleteRadical(c echo.Context) error {
	name := c.Param("name")

	ctx := c.Request().Context()
	if err := api.cards.DeleteRadical(ctx, name); err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.NoContent(http.StatusNoContent)
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
func (api *CardsApi) AllRadicals(c echo.Context) error {
	radicals, err := api.cards.AllRadicals(c.Request().Context())
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, radicals)
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
func (api *CardsApi) QueryRadicalKanjis(c echo.Context) error {
	name := c.Param("name")
	kanjis, err := api.cards.QueryRadicalKanjis(c.Request().Context(), name)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, kanjis)
}
