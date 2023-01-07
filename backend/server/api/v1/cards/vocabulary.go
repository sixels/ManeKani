package v1

import (
	"net/http"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"

	"github.com/labstack/echo/v4"
)

// CreateVocabulary godoc
// @Id post-vocabulary-create
// @Summary Create a new vocabulary
// @Description Creates a vocabulary with the given values
// @Tags cards, vocabulary
// @Accept json
// @Produce json
// @Param vocabulary body cards.CreateVocabularyRequest true "The vocabulary to be created"
// @Success 201 {object} cards.Vocabulary
// @Router /api/v1/vocabulary [post]
func (api *CardsApi) CreateVocabulary(c echo.Context) error {
	vocabulary := new(cards.CreateVocabularyRequest)
	if err := c.Bind(vocabulary); err != nil {
		return err
	}

	ctx := c.Request().Context()
	created, err := api.cards.CreateVocabulary(ctx, *vocabulary)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}

	return c.JSON(http.StatusCreated, created)
}

// QueryVocabulary godoc
// @Id get-vocabulary-query
// @Summary Query a vocabulary
// @Description Search a vocabulary by its word
// @Tags cards, vocabulary
// @Accept */*
// @Produce json
// @Param word path string true "Vocabulary word"
// @Success 200 {object} cards.Vocabulary
// @Router /api/v1/vocabulary/{word} [get]
func (api *CardsApi) QueryVocabulary(c echo.Context) error {
	word := c.Param("word")

	ctx := c.Request().Context()
	queried, err := api.cards.QueryVocabulary(ctx, word)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, queried)
}

// UpdateVocabulary godoc
// @Id patch-vocabulary-update
// @Summary Update a vocabulary
// @Description Update a vocabulary with the given values
// @Tags cards, vocabulary
// @Accept json
// @Produce json
// @Param word path string true "Vocabulary word"
// @Param vocabulary body cards.UpdateVocabularyRequest true "Vocabulary fields to update"
// @Success 200 {object} cards.Vocabulary
// @Router /api/v1/vocabulary/{word} [patch]
func (api *CardsApi) UpdateVocabulary(c echo.Context) error {
	word := c.Param("word")

	vocabulary := new(cards.UpdateVocabularyRequest)
	if err := c.Bind(vocabulary); err != nil {
		return err
	}

	ctx := c.Request().Context()
	updated, err := api.cards.UpdateVocabulary(ctx, word, *vocabulary)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteVocabulary godoc
// @Id delete-vocabulary-delete
// @Summary Delete a vocabulary
// @Description Delete a vocabulary by its word
// @Tags cards, vocabulary
// @Accept */*
// @Produce json
// @Param word path string true "Vocabulary word"
// @Success 200
// @Router /api/v1/vocabulary/{word} [delete]
func (api *CardsApi) DeleteVocabulary(c echo.Context) error {
	word := c.Param("word")

	ctx := c.Request().Context()
	if err := api.cards.DeleteVocabulary(ctx, word); err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.NoContent(http.StatusNoContent)
}

// AllVocabulary godoc
// @Id get-vocabulary-all
// @Summary Query all vocabulary
// @Description Return a list of all vocabulary
// @Tags cards, vocabulary
// @Accept */*
// @Produce json
// @Success 200 {array} cards.PartialVocabularyResponse
// @Router /api/v1/vocabulary [get]
func (api *CardsApi) AllVocabularies(c echo.Context) error {
	filters := new(cards.QueryAllVocabularyRequest)
	if err := c.Bind(filters); err != nil {
		return err
	}

	vocabularies, err := api.cards.AllVocabularies(c.Request().Context(), *filters)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, vocabularies)
}

// QueryVocabularyKanjis godoc
// @Id get-vocabulary-kanji
// @Summary Query kanjis from vocabulary
// @Description Return a list of all kanji that composes the given vocabulary
// @Tags cards, vocabulary
// @Accept */*
// @Produce json
// @Param word path string true "Vocabulary word"
// @Success 200 {array} cards.PartialKanjiResponse
// @Router /api/v1/vocabulary/{:word}/kanji [get]
func (api *CardsApi) QueryVocabularyKanjis(c echo.Context) error {
	word := c.Param("word")
	kanjis, err := api.cards.QueryVocabularyKanjis(c.Request().Context(), word)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, kanjis)
}
