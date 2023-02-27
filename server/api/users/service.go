package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/services/ent/users"
	mkjwt "github.com/sixels/manekani/services/jwt"
)

type UserApi struct {
	users *users.UsersRepository
	jwt   *mkjwt.JWTService
}

func New(repo *users.UsersRepository, jwtService *mkjwt.JWTService) *UserApi {
	return &UserApi{users: repo, jwt: jwtService}
}

func (api *UserApi) SetupRoutes(router *gin.Engine) {
	router.GET("/user", api.RequiresUser(), api.GetBasicUserInfo())
	// router.GET("/user/decks", api.RequiresUser(), api.GetSRSInfo())
}
