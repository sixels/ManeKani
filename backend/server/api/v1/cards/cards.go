package v1

import (
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/files"

	"github.com/labstack/echo/v4"
)

type CardsApi struct {
	cards *cards.CardsService
	files *files.FilesService
}

func New(cardsService *cards.CardsService, filesService *files.FilesService) *CardsApi {
	return &CardsApi{
		cards: cardsService,
		files: filesService,
	}
}

func (api *CardsApi) SetupRoutes(router *echo.Echo) {
	apiV1 := router.Group("/api/v1")

	apiV1.POST("/radical", api.CreateRadical, api.UploadRadicalImage)
	apiV1.GET("/radical/:name", api.QueryRadical)
	apiV1.PATCH("/radical/:name", api.UpdateRadical)
	apiV1.DELETE("/radical/:name", api.DeleteRadical)
	apiV1.GET("/radical", api.AllRadicals)
	apiV1.GET("/radical/:name/kanji", api.QueryRadicalKanjis)

	apiV1.POST("/kanji", api.CreateKanji)
	apiV1.GET("/kanji/:symbol", api.QueryKanji)
	apiV1.PATCH("/kanji/:symbol", api.UpdateKanji)
	apiV1.DELETE("/kanji/:symbol", api.DeleteKanji)
	apiV1.GET("/kanji", api.AllKanji)
	apiV1.GET("/kanji/:symbol/radicals", api.QueryKanjiRadicals)
	apiV1.GET("/kanji/:symbol/vocabularies", api.QueryKanjiVocabularies)

	apiV1.POST("/vocabulary", api.CreateVocabulary)
	apiV1.GET("/vocabulary/:word", api.QueryVocabulary)
	apiV1.PATCH("/vocabulary/:word", api.UpdateVocabulary)
	apiV1.DELETE("/vocabulary/:word", api.DeleteVocabulary)
	apiV1.GET("/vocabulary", api.AllVocabularies)
	apiV1.GET("/vocabulary/:word/kanji", api.QueryVocabularyKanjis)

	apiV1.GET("/level", api.AllLevels)
}
