package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthService struct {
}

func (s *HealthService) ServiceName() string {
	return "health"
}

func (s *HealthService) SetupRoutes(router *echo.Echo) {
	router.GET("/health", s.healthCheck)
}

func (s *HealthService) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "server is online",
	})
}
