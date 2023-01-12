package v1

import (
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/files"

	authApi "sixels.io/manekani/server/api/auth"

	"github.com/gin-gonic/gin"
)

type CardsApi struct {
	cards *cards.CardsRepository
	files *files.FilesRepository
}

func New(
	cardsService *cards.CardsRepository,
	filesService *files.FilesRepository,
) *CardsApi {
	return &CardsApi{
		cards: cardsService,
		files: filesService,
	}
}

func (api *CardsApi) SetupRoutes(router *gin.Engine) {
	r := router.Group("/api/v1")

	r.
		POST("/radical", authApi.VerifySession(nil), api.UploadRadicalImage(), api.CreateRadical()).
		GET("/radical/:name", api.QueryRadical()).
		PATCH("/radical/:name", authApi.VerifySession(nil), api.UpdateRadical()).
		DELETE("/radical/:name", authApi.VerifySession(nil), api.DeleteRadical()).
		GET("/radical", api.AllRadicals()).
		GET("/radical/:name/kanji", api.QueryRadicalKanjis())
	r.
		POST("/kanji", authApi.VerifySession(nil), api.CreateKanji()).
		GET("/kanji/:symbol", api.QueryKanji()).
		PATCH("/kanji/:symbol", authApi.VerifySession(nil), api.UpdateKanji()).
		DELETE("/kanji/:symbol", authApi.VerifySession(nil), api.DeleteKanji()).
		GET("/kanji", api.AllKanji()).
		GET("/kanji/:symbol/radicals", api.QueryKanjiRadicals()).
		GET("/kanji/:symbol/vocabularies", api.QueryKanjiVocabularies())
	r.
		POST("/vocabulary", authApi.VerifySession(nil), api.CreateVocabulary()).
		GET("/vocabulary/:word", api.QueryVocabulary()).
		PATCH("/vocabulary/:word", authApi.VerifySession(nil), api.UpdateVocabulary()).
		DELETE("/vocabulary/:word", authApi.VerifySession(nil), api.DeleteVocabulary()).
		GET("/vocabulary", api.AllVocabularies()).
		GET("/vocabulary/:word/kanji", api.QueryVocabularyKanjis())
	r.
		GET("/level", api.AllLevels())
}
