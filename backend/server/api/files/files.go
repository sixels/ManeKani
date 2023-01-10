package files

import (
	"github.com/gin-gonic/gin"
	"sixels.io/manekani/services/files"
)

type FilesApi struct {
	files *files.FilesRepository
}

func New(filesService *files.FilesRepository) *FilesApi {
	return &FilesApi{
		files: filesService,
	}
}

func (api *FilesApi) SetupRoutes(router *gin.Engine) {
	r := router.Group("/files")

	r.GET("/:kind/:namespace/:name", api.QueryFile())
}
