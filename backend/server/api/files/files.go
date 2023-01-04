package files

import (
	"sixels.io/manekani/services/files"

	"github.com/labstack/echo/v4"
)

type FilesApi struct {
	files *files.FilesService
}

func New(filesService *files.FilesService) *FilesApi {
	return &FilesApi{
		files: filesService,
	}
}

func (api *FilesApi) SetupRoutes(router *echo.Echo) {
	filesRoute := router.Group("/files")

	filesRoute.GET("/:kind/:namespace/:name", api.QueryFile)
}
