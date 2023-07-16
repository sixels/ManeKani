package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/server/api/users/hooks"
)

type UserApi struct {
	users ports.UserRepository
}

func (api *UserApi) ServiceName() string {
	return "users"
}

func New(repo ports.UserRepository) *UserApi {
	return &UserApi{users: repo}
}

func (api *UserApi) SetupRoutes(router *gin.Engine) {
	hookHandler := router.Group("/hook/user")
	hookHandler.POST("/register", hooks.RegisterUser(api.users))

	router.GET("/user", api.RequiresUser(), api.GetBasicUserInfo())
}
