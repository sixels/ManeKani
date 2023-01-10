package v1

import (
	"sixels.io/manekani/server/middlewares"
	"sixels.io/manekani/services/auth"
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/files"

	"github.com/gin-gonic/gin"
)

type CardsApi struct {
	cards         *cards.CardsRepository
	files         *files.FilesRepository
	authenticator *auth.Authenticator
}

func New(
	cardsService *cards.CardsRepository,
	filesService *files.FilesRepository,
	authService *auth.Authenticator,
) *CardsApi {
	return &CardsApi{
		cards:         cardsService,
		files:         filesService,
		authenticator: authService,
	}
}

func (api *CardsApi) SetupRoutes(router *gin.Engine) {
	r := router.Group("/api/v1")
	loginRequired := middlewares.LoginRequired(*api.authenticator)

	r.
		POST("/radical", api.CreateRadical(), loginRequired, api.UploadRadicalImage()).
		GET("/radical/:name", api.QueryRadical()).
		PATCH("/radical/:name", api.UpdateRadical(), loginRequired).
		DELETE("/radical/:name", api.DeleteRadical(), loginRequired).
		GET("/radical", api.AllRadicals(), loginRequired).
		GET("/radical/:name/kanji", api.QueryRadicalKanjis())
	r.
		POST("/kanji", api.CreateKanji(), loginRequired).
		GET("/kanji/:symbol", api.QueryKanji()).
		PATCH("/kanji/:symbol", api.UpdateKanji(), loginRequired).
		DELETE("/kanji/:symbol", api.DeleteKanji(), loginRequired).
		GET("/kanji", api.AllKanji()).
		GET("/kanji/:symbol/radicals", api.QueryKanjiRadicals()).
		GET("/kanji/:symbol/vocabularies", api.QueryKanjiVocabularies())
	r.
		POST("/vocabulary", api.CreateVocabulary(), loginRequired).
		GET("/vocabulary/:word", api.QueryVocabulary()).
		PATCH("/vocabulary/:word", api.UpdateVocabulary(), loginRequired).
		DELETE("/vocabulary/:word", api.DeleteVocabulary(), loginRequired).
		GET("/vocabulary", api.AllVocabularies()).
		GET("/vocabulary/:word/kanji", api.QueryVocabularyKanjis())
	r.
		GET("/level", api.AllLevels())
}
