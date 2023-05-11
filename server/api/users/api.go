package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/server/api/users/hooks"
	"github.com/sixels/manekani/server/auth"
)

type UserApi struct {
	users ports.UserRepository
	auth  *auth.AuthService
}

func New(repo ports.UserRepository, auth *auth.AuthService) *UserApi {
	return &UserApi{users: repo, auth: auth}
}

func (api *UserApi) SetupRoutes(router *gin.Engine) {
	hookHandler := router.Group("/hook/user")
	hookHandler.POST("/register", hooks.RegisterUser(api.users))

	router.GET("/user", api.RequiresUser(), api.GetBasicUserInfo())
}
