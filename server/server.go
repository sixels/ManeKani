package server

import (
	"encoding/gob"
	"log"
	"os/user"

	"github.com/gin-gonic/gin"
)

type Service interface {
	ServiceName() string
	SetupRoutes(router *gin.Engine)
}

type Server struct {
	addr   string
	router *gin.Engine
}

func New(addr string) *Server {
	gob.Register(user.User{})

	router := gin.New()
	router.Use(DefaultMiddlewares()...)
	router.RedirectTrailingSlash = false
	// router.RemoveExtraSlash = true

	return &Server{
		router: router,
		addr:   addr,
	}
}

func (s *Server) WithService(service Service) *Server {
	log.Printf("Registering service: %s\n", service.ServiceName())
	service.SetupRoutes(s.router)
	return s
}

func (s *Server) WithMiddleware(mw ...gin.HandlerFunc) *Server {
	s.router.Use(mw...)
	return s
}

func (s *Server) Start() {
	s.WithService(&HealthService{})
	log.Fatal(s.router.Run(s.addr))
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func DefaultMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{gin.Logger(), gin.Recovery()}
}
