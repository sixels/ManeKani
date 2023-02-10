package cards

import (
	"sixels.io/manekani/services/ent/cards"
	"sixels.io/manekani/services/files"

	auth_api "sixels.io/manekani/server/api/auth"
	users_api "sixels.io/manekani/server/api/users"
	mkjwt "sixels.io/manekani/services/jwt"

	"github.com/gin-gonic/gin"
)

type CardsApi struct {
	cards *cards.CardsRepository
	files *files.FilesRepository
	jwt   *mkjwt.JWTService
}

func New(
	cardsService *cards.CardsRepository,
	filesService *files.FilesRepository,
	jwtService *mkjwt.JWTService,
) *CardsApi {
	return &CardsApi{
		cards: cardsService,
		files: filesService,
		jwt:   jwtService,
	}
}

func (api *CardsApi) SetupRoutes(router *gin.Engine) {
	r := router.Group("/api/v1")

	r.
		POST("/subject",
			// auth_api.VerifySession(&sessmodels.VerifySessionOptions{SessionRequired: &flase}),
			users_api.EnsureCapabilities(api.jwt, mkjwt.TokenCababiltySubjectCreate),
			api.CreateSubject()).
		GET("/subject/:id", api.QuerySubject()).
		PATCH("/subject/:id",
			auth_api.VerifySession(nil),
			users_api.EnsureCapabilities(api.jwt, mkjwt.TokenCababiltySubjectUpdate),
			api.UpdateSubject()).
		DELETE("/subject/:name",
			auth_api.VerifySession(nil),
			users_api.EnsureCapabilities(api.jwt, mkjwt.TokenCababiltySubjectDelete),
			api.DeleteSubject()).
		GET("/subject", api.AllSubjects())

	r.
		GET("/deck/:id", api.QueryDeck()).
		GET("/deck", api.AllDecks())
}
