package docs

import (
	"github.com/gin-gonic/gin"
)

type DocsHandler struct {
}

func New() *DocsHandler {
	return &DocsHandler{}
}

func (api *DocsHandler) ServiceName() string {
	return "docs"
}
func (docs *DocsHandler) SetupRoutes(router *gin.Engine) {
	handler := router.Group("/docs")
	handler.StaticFS("/", gin.Dir("docs/manekani", true))
}
