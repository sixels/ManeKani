package server

import (
	"encoding/gob"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sixels/manekani/server/api/apicommon"
	"net/http"
	"os/user"
)

type Service interface {
	ServiceName() string
	SetupRoutes(router *echo.Echo)
}

type Server struct {
	addr   string
	router *echo.Echo
}

func New(addr string) *Server {
	gob.Register(user.User{})

	log.EnableColor()
	log.SetLevel(log.DEBUG)

	router := echo.New()
	router.Use(DefaultMiddlewares()...)

	//router.RedirectTrailingSlash = false
	// router.RemoveExtraSlash = true

	router.RouteNotFound("/api/*", func(c echo.Context) error {
		return apicommon.Error(http.StatusNotFound, errors.New("API route not found"))
	})

	return &Server{
		router: router,
		addr:   addr,
	}
}

func (s *Server) WithService(service Service) *Server {
	log.Infof("Registering service: %s\n", service.ServiceName())
	service.SetupRoutes(s.router)
	return s
}

func (s *Server) WithMiddleware(mw ...echo.MiddlewareFunc) *Server {
	s.router.Use(mw...)
	return s
}

func (s *Server) Start() {
	s.WithService(&HealthService{})
	log.Fatal(s.router.Start(s.addr))
}

func (s *Server) Router() *echo.Echo {
	return s.router
}

func DefaultMiddlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{middleware.Logger(), middleware.Recover(), middleware.RemoveTrailingSlash()}
}
