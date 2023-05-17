package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) bindRoutes() {
	server.router.GET("/health", healthCheck)

	for _, service := range server.services {
		service.SetupRoutes(server.router)
	}
}

// HealthCheck godoc
//
//	@Id				get-health-check
//	@Summary		Show the status of the server.
//	@Description	Get the status of the server.
//	@Tags			health
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": "server is online",
	})
}
