package server

import (
	"context"
	"io"
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
	router *echo.Echo

	services []Service
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
		router:   router,
		services: []Service{cardsV1, filesApi},
	}
}

func (server *Server) Start(logFile io.Writer) {
	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Output = logFile
	server.router.Use(middleware.LoggerWithConfig(loggerConfig))

	server.router.Use(middleware.Recover())

	server.bindRoutes()

	log.Fatal(server.router.Start("0.0.0.0:8081"))
}
