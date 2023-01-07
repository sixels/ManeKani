package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
)

// AllLevels godoc
// @Id get-level-all
// @Summary Query all radicals, kanji and vocabularies by level
// @Description Return a list of all radicals, kanji and vocabularies by the given level
// @Tags cards, kanji, radical, vocabulary
// @Accept */*
// @Produce json
// @Success 200 {array} cards.Level
// @Router /api/v1/level [get]
func (api *CardsApi) AllLevels(c echo.Context) error {
	filters := new(cards.FilterLevel)
	if err := c.Bind(filters); err != nil {
		return err
	}

	ctx := c.Request().Context()

	kanjis, err := api.cards.AllKanji(ctx, cards.QueryAllKanjiRequest{FilterLevel: *filters})
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	radicals, err := api.cards.AllRadicals(ctx, cards.QueryAllRadicalRequest{FilterLevel: *filters})
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}
	vocabularies, err := api.cards.AllVocabularies(ctx, cards.QueryAllVocabularyRequest{FilterLevel: *filters})
	if err != nil {
		return c.JSON(err.(*errors.Error).Status, err)
	}

	return c.JSON(http.StatusOK, cards.Level{
		Kanji:      kanjis,
		Radical:    radicals,
		Vocabulary: vocabularies,
	})
}
