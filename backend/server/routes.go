package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (server *Server) bindRoutes() {
	// TODO: don't expose swagger routes in production
	if true {
		server.router.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	server.router.GET("/health", healthCheck)

	for _, service := range server.services {
		service.SetupRoutes(server.router)
	}
}

// HealthCheck godoc
// @Id get-health-check
// @Summary Show the status of the server.
// @Description Get the status of the server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "server is online",
	})
}
