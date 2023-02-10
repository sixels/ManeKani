package users

import (
	"github.com/gin-gonic/gin"
	"sixels.io/manekani/services/ent/users"
	mkjwt "sixels.io/manekani/services/jwt"
)

type UserApi struct {
	users *users.UsersRepository
	jwt   *mkjwt.JWTService
}

func New(repo *users.UsersRepository, jwtService *mkjwt.JWTService) *UserApi {
	return &UserApi{users: repo, jwt: jwtService}
}

func (api *UserApi) SetupRoutes(router *gin.Engine) {
	router.GET("/user", api.GetBasicUserInfo())
	// router.GET("/user/decks", api.RequiresUser(), api.GetSRSInfo())
	router.GET("/user/decks/:deck-id", api.RequiresUser(), api.GetSRSInfo())
	router.DELETE("/user/decks/:deck-id", api.RequiresUser(), api.UnsubscribeUserFromDeck())
}
