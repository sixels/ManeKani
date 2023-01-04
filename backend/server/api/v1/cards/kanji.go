package v1

import (
	"net/http"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"

	"github.com/labstack/echo/v4"
)

// CreateKanji godoc
// @Id post-kanji-create
// @Summary Create a new kanji
// @Description Creates a kanji with the given values
// @Tags cards, kanji
// @Accept json
// @Produce json
// @Param kanji body cards.CreateKanjiRequest true "The kanji to be created"
// @Success 201 {object} cards.Kanji
// @Router /api/v1/kanji [post]
func (api *CardsApi) CreateKanji(c echo.Context) error {
	kanji := new(cards.CreateKanjiRequest)
	if err := c.Bind(kanji); err != nil {
		return err
	}

	ctx := c.Request().Context()
	created, err := api.cards.CreateKanji(ctx, *kanji)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}

	return c.JSON(http.StatusCreated, created)
}

// QueryKanji godoc
// @Id get-kanji-query
// @Summary Query a kanji
// @Description Search a kanji by its symbol
// @Tags cards, kanji
// @Accept */*
// @Produce json
// @Param symbol path string true "Kanji symbol"
// @Success 200 {object} cards.Kanji
// @Router /api/v1/kanji/{symbol} [get]
func (api *CardsApi) QueryKanji(c echo.Context) error {
	symbol := c.Param("symbol")

	ctx := c.Request().Context()
	queried, err := api.cards.QueryKanji(ctx, symbol)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, queried)
}

// UpdateKanji godoc
// @Id patch-kanji-update
// @Summary Update a kanji
// @Description Update a kanji with the given values
// @Tags cards, kanji
// @Accept json
// @Produce json
// @Param symbol path string true "Kanji symbol"
// @Param kanji body cards.UpdateKanjiRequest true "Kanji fields to update"
// @Success 200 {object} cards.Kanji
// @Router /api/v1/kanji/{symbol} [patch]
func (api *CardsApi) UpdateKanji(c echo.Context) error {
	symbol := c.Param("symbol")

	kanji := new(cards.UpdateKanjiRequest)
	if err := c.Bind(kanji); err != nil {
		return err
	}

	ctx := c.Request().Context()
	updated, err := api.cards.UpdateKanji(ctx, symbol, *kanji)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteKanji godoc
// @Id delete-kanji-delete
// @Summary Delete a kanji
// @Description Delete a kanji by its symbol
// @Tags cards, kanji
// @Accept */*
// @Produce json
// @Param symbol path string true "Kanji symbol"
// @Success 200
// @Router /api/v1/kanji/{symbol} [delete]
func (api *CardsApi) DeleteKanji(c echo.Context) error {
	symbol := c.Param("symbol")

	ctx := c.Request().Context()
	if err := api.cards.DeleteKanji(ctx, symbol); err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.NoContent(http.StatusNoContent)
}

// AllKanji godoc
// @Id get-kanji-all
// @Summary Query all kanji
// @Description Return a list of all kanji
// @Tags cards, kanji
// @Accept */*
// @Produce json
// @Success 200 {array} cards.PartialKanjiResponse
// @Router /api/v1/kanji [get]
func (api *CardsApi) AllKanji(c echo.Context) error {
	ctx := c.Request().Context()
	kanjis, err := api.cards.AllKanji(ctx)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, kanjis)
}

// QueryKanjiRadicals godoc
// @Id get-kanji-radicals
// @Summary Query a kanji's radicals
// @Description Return a list of all radicals that composes the given kanji
// @Tags cards, kanji
// @Accept */*
// @Produce json
// @Param symbol path string true "Kanji symbol"
// @Success 200 {array} cards.PartialRadicalResponse
// @Router /api/v1/kanji/{:symbol}/radicals [get]
func (api *CardsApi) QueryKanjiRadicals(c echo.Context) error {
	symbol := c.Param("symbol")
	radicals, err := api.cards.QueryKanjiRadicals(c.Request().Context(), symbol)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, radicals)
}

// QueryKanjiVocabularies godoc
// @Id get-kanji-vocabularies
// @Summary Query a kanji's vocabularies
// @Description Return a list of all vocabularies that are composed by the given kanji
// @Tags cards, kanji
// @Accept */*
// @Produce json
// @Param symbol path string true "Kanji symbol"
// @Success 200 {array} cards.PartialVocabularyResponse
// @Router /api/v1/kanji/{:symbol}/vocabularies [get]
func (api *CardsApi) QueryKanjiVocabularies(c echo.Context) error {
	symbol := c.Param("symbol")
	vocabularies, err := api.cards.QueryKanjiVocabularies(c.Request().Context(), symbol)
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	return c.JSON(http.StatusOK, vocabularies)
}
