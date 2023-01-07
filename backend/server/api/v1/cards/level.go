package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
)

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
