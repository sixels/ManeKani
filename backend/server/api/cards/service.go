package cards

import (
	"sixels.io/manekani/services/ent/cards"
	"sixels.io/manekani/services/ent/users"
	"sixels.io/manekani/services/files"

	auth_api "sixels.io/manekani/server/api/auth"
	mkjwt "sixels.io/manekani/services/jwt"

	"github.com/gin-gonic/gin"

	cards1 "sixels.io/manekani/server/api/cards/v1"
)

type CardsApi struct {
	V1  cards1.CardsApiV1
	jwt *mkjwt.JWTService
}

func New(
	cardsService *cards.CardsRepository,
	filesService *files.FilesRepository,
	usersService *users.UsersRepository,
	jwtService *mkjwt.JWTService,
) *CardsApi {
	return &CardsApi{
		V1: cards1.CardsApiV1{
			Cards: cardsService,
			Users: usersService,
			Files: filesService,
		},
		jwt: jwtService,
	}
}

func (api *CardsApi) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	// session
	router.
		GET("/study-session",
			auth_api.EnsureLogin(),
			api.StudySession())

	// subject
	v1.
		POST("/subject",
			auth_api.EnsureCapabilities(api.jwt, mkjwt.TokenCapabilitySubjectCreate),
			api.V1.CreateSubject()).
		GET("/subject/:id", api.V1.QuerySubject()).
		PATCH("/subject/:id",
			auth_api.EnsureCapabilities(api.jwt, mkjwt.TokenCapabilitySubjectUpdate),
			api.V1.UpdateSubject()).
		DELETE("/subject/:name",
			auth_api.EnsureCapabilities(api.jwt, mkjwt.TokenCapabilitySubjectDelete),
			api.V1.DeleteSubject()).
		GET("/subject", api.V1.AllSubjects())

	// deck
	v1.
		GET("/deck/:id", api.V1.QueryDeck()).
		GET("/deck", api.V1.AllDecks()).
		PATCH("/deck/:id/subscribe",
			auth_api.EnsureCapabilities(api.jwt, mkjwt.TokenCapabilityUserUpdate),
			api.V1.SubscribeUserToDeck()).
		DELETE("/deck/:id/subscribe",
			auth_api.EnsureCapabilities(api.jwt, mkjwt.TokenCapabilityUserUpdate),
			api.V1.UnsubscribeUserFromDeck())

	// cards
	v1.
		GET("/card",
			auth_api.EnsureCapabilities(api.jwt),
			api.V1.AllUserCards())

	// reviews
	v1.
		GET("/review", auth_api.EnsureCapabilities(api.jwt),
			api.V1.AllUserReviews()).
		POST("/review",
			auth_api.EnsureCapabilities(api.jwt, mkjwt.TokenCapabilityReviewCreate),
			api.V1.CreateReview())
}
