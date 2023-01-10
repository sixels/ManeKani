package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (api *CardsApi) AllLevels() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := new(cards.FilterLevel)
		if err := c.Bind(filters); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		kanjis, err := queryAll(ctx, api.cards.AllKanji, filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
			return
		}
		radicals, err := queryAll(ctx, api.cards.AllRadicals, filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
			return
		}
		vocabularies, err := queryAll(ctx, api.cards.AllVocabularies, filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
			return
		}

		c.JSON(http.StatusOK, cards.Level{
			Kanji:      kanjis,
			Radical:    radicals,
			Vocabulary: vocabularies,
		})
	}
}

type filterByLevel interface {
	~struct{ cards.FilterLevel }
}

func queryAll[T filterByLevel, A any](ctx context.Context, q func(context.Context, T) ([]*A, error), filters *cards.FilterLevel) ([]*A, error) {
	return q(ctx, T{*filters})
}
