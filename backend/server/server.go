package server

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "sixels.io/manekani/docs/manekani"
	filesApi "sixels.io/manekani/server/api/files"
	v1CardsApi "sixels.io/manekani/server/api/v1/cards"
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/files"
)

type Service interface {
	SetupRoutes(router *echo.Echo)
}

type Server struct {
	router  *echo.Echo
	v1Cards Service
	files   Service
}

func New() *Server {
	cardsRepo, err := cards.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Cards repository: %v", err)
	}

	filesRepo, err := files.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Files repository: %v", err)
	}

	cardsService := cards.NewService(cardsRepo)
	filesService := files.NewService(filesRepo)

	cardsV1 := v1CardsApi.New(cardsService, filesService)
	filesApi := filesApi.New(filesService)

	router := echo.New()

	return &Server{
		router:  router,
		v1Cards: cardsV1,
		files:   filesApi,
	}
}

func (server *Server) UseLogger() *Server {
	server.router.Use(middleware.Logger())
	server.router.Use(middleware.Recover())
	return server
}

func (server *Server) Start() {
	server.bindRoutes()
	log.Fatal(server.router.Start("localhost:8081"))
}
