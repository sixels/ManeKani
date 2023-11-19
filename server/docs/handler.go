package docs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DocsHandler struct {
}

func New() *DocsHandler {
	return &DocsHandler{}
}

func (api *DocsHandler) ServiceName() string {
	return "docs"
}
func (docs *DocsHandler) SetupRoutes(router *echo.Echo) {
	handler := router.Group("/docs")
	//dir, err := fs.FS("docs/manekani")
	//if
	handler.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "docs/manekani",
		Browse: false,
	}))
}
