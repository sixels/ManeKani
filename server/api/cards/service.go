package cards

import (
	"github.com/sixels/manekani/core/adapters/cards"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/services/ent/users"
	"github.com/sixels/manekani/services/files"

	auth_api "github.com/sixels/manekani/server/auth"

	"github.com/gin-gonic/gin"

	cards1 "github.com/sixels/manekani/server/api/cards/v1"
)

type CardsApi struct {
	V1   cards1.CardsApiV1
	auth *auth_api.AuthService
}

func New(
	cardsService cards.CardsAdapter,
	filesService *files.FilesRepository,
	usersService *users.UsersRepository,
	authMiddleware *auth_api.AuthService,
) *CardsApi {
	return &CardsApi{
		V1: cards1.CardsApiV1{
			Cards: cardsService,
			Users: usersService,
			Files: filesService,
		},
		auth: authMiddleware,
	}
}

func (api *CardsApi) SetupRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	v1Router := apiRouter.Group("/v1")

	// session
	apiRouter.
		GET("/study-session",
			api.auth.EnsureLogin(),
			api.StudySession())

	// subject
	v1Router.
		POST("/subject",
			api.auth.EnsurePermissions(tokens.TokenPermissionSubjectCreate),
			api.V1.CreateSubject()).
		GET("/subject/:id", api.V1.QuerySubject()).
		PATCH("/subject/:id",
			api.auth.EnsurePermissions(tokens.TokenPermissionSubjectUpdate),
			api.V1.UpdateSubject()).
		DELETE("/subject/:name",
			api.auth.EnsurePermissions(tokens.TokenPermissionSubjectDelete),
			api.V1.DeleteSubject()).
		GET("/subject", api.V1.AllSubjects())

	// deck
	v1Router.
		GET("/deck/:id", api.V1.QueryDeck()).
		GET("/deck", api.V1.AllDecks()).
		PUT("/deck/:id/subscribe",
			api.auth.EnsurePermissions(tokens.TokenPermissionUserUpdate),
			api.V1.SubscribeUserToDeck()).
		DELETE("/deck/:id/subscribe",
			api.auth.EnsurePermissions(tokens.TokenPermissionUserUpdate),
			api.V1.UnsubscribeUserFromDeck())

	// cards
	v1Router.
		GET("/card",
			api.auth.EnsurePermissions(),
			api.V1.AllUserCards())

	// reviews
	v1Router.
		GET("/review", api.auth.EnsurePermissions(),
			api.V1.AllUserReviews()).
		POST("/review",
			api.auth.EnsurePermissions(tokens.TokenPermissionReviewCreate),
			api.V1.CreateReview())
}
