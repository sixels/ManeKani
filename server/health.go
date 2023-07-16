package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthService struct {
}

func (s *HealthService) ServiceName() string {
	return "health"
}

func (s *HealthService) SetupRoutes(router *gin.Engine) {
	router.GET("/health", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": "server is online",
	})
}
