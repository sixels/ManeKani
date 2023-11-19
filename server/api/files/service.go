package files

import (
	"github.com/labstack/echo/v4"
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

func (api *FilesApi) SetupRoutes(router *echo.Echo) {
	r := router.Group("/files")

	r.GET("/*path", api.QueryFile)
}
