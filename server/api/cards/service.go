package cards

import (
	"github.com/labstack/echo/v4"
	"github.com/sixels/manekani/core/adapters/cards"

	cards1 "github.com/sixels/manekani/server/api/cards/v1"
)

type CardsApi struct {
	V1 cards1.CardsApiV1
}

func (api *CardsApi) ServiceName() string {
	return "subjects"
}

func New(
	cardsService cards.CardsAdapter,
) *CardsApi {
	return &CardsApi{
		V1: cards1.CardsApiV1{
			Cards: cardsService,
		},
	}
}

func (api *CardsApi) SetupRoutes(router *echo.Echo) {
	cards1.RegisterHandlers(router, &api.V1)

	// apiRouter := router.Group("/api")
	// v1Router := apiRouter.Group("/v1")

	// // session
	// apiRouter.
	// 	GET("/study-session",
	// 		// api.auth.EnsureLogin(),
	// 		api.StudySession())

	// // subject
	// v1Router.
	// 	POST("/subjects",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionSubjectCreate),
	// 		api.V1.CreateSubject()).
	// 	GET("/subjects/:id", api.V1.QuerySubject()).
	// 	PATCH("/subjects/:id",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionSubjectUpdate),
	// 		api.V1.UpdateSubject()).
	// 	DELETE("/subjects/:name",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionSubjectDelete),
	// 		api.V1.DeleteSubject()).
	// 	GET("/subjects", api.V1.AllSubjects())

	// // deck
	// v1Router.
	// 	POST("/decks",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionDeckCreate),
	// 		api.V1.CreateDeck()).
	// 	GET("/decks/:id", api.V1.QueryDeck()).
	// 	GET("/decks", api.V1.AllDecks()).
	// 	PUT("/decks/:id/subscribe",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionUserUpdate),
	// 		api.V1.SubscribeUserToDeck()).
	// 	DELETE("/decks/:id/subscribe",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionUserUpdate),
	// 		api.V1.UnsubscribeUserFromDeck())

	// // cards
	// v1Router.
	// 	GET("/card",
	// 		// api.auth.EnsurePermissions(),
	// 		api.V1.AllUserCards())

	// // reviews
	// v1Router.
	// 	GET("/review",
	// 		// api.auth.EnsurePermissions(),
	// 		api.V1.AllUserReviews()).
	// 	POST("/review",
	// 		// api.auth.EnsurePermissions(tokens.TokenPermissionReviewCreate),
	// 		api.V1.CreateReview())
}
