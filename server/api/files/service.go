package files

import (
	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/services/files"
)

type FilesApi struct {
	files *files.FilesRepository
}

func (api *FilesApi) ServiceName() string {
	return "files"
}

func New(filesService *files.FilesRepository) *FilesApi {
	return &FilesApi{
		files: filesService,
	}
}

func (api *FilesApi) SetupRoutes(router *gin.Engine) {
	r := router.Group("/files")

	r.GET("/*path", api.QueryFile())
}
