package users

import (
	"github.com/gin-gonic/gin"
	"sixels.io/manekani/services/users"
)

type UserApi struct {
	users *users.UsersRepository
}

func New(repo *users.UsersRepository) *UserApi {
	return &UserApi{users: repo}
}

func (api *UserApi) SetupRoutes(router *gin.Engine) {
	router.GET("/user", api.GetBasicUserInfo())
	router.GET("/user/srs", api.RequiresUser(), api.GetSRSInfo())
	router.GET("/user/srs/reset", api.RequiresUser(), api.ResetSRSData())
}
