package users

import (
	"github.com/labstack/echo/v4"
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

func (api *UserApi) SetupRoutes(router *echo.Echo) {
	hookHandler := router.Group("/hook/user")
	hookHandler.POST("/register", hooks.RegisterUser(api.users))

	//router.GET("/user", api.GetBasicUserInfo, api.RequiresUser)
}
