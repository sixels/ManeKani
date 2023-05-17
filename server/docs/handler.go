package docs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DocsHandler struct {
	schemaURL string
}

func New(schemaURL string) *DocsHandler {
	return &DocsHandler{schemaURL}
}

func (docs *DocsHandler) SetupRoutes(router *gin.Engine) {
	handler := router.Group("/docs")
	handler.StaticFS("/", http.Dir("docs/manekani"))
}
