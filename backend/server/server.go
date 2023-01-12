package server

import (
	"context"
	"io"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"

	_ "sixels.io/manekani/docs/manekani"

	authApi "sixels.io/manekani/server/api/auth"
	filesApi "sixels.io/manekani/server/api/files"
	v1CardsApi "sixels.io/manekani/server/api/v1/cards"

	"sixels.io/manekani/services/auth"
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/files"
)

type Service interface {
	SetupRoutes(router *gin.Engine)
}

type Server struct {
	router *gin.Engine

	services []Service
}

func New() *Server {
	if err := auth.StartAuthenticator(); err != nil {
		log.Fatalf("Could not start the authenticator: %v", err)
	}

	cardsRepo, err := cards.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Cards repository: %v", err)
	}

	filesRepo, err := files.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Files repository: %v", err)
	}

	cardsV1 := v1CardsApi.New(cardsRepo, filesRepo)
	filesApi := filesApi.New(filesRepo)

	router := gin.Default()

	return &Server{
		router:   router,
		services: []Service{cardsV1, filesApi},
	}
}

func (server *Server) Start(logFile io.Writer) {
	// cors
	server.router.Use(cors.New(cors.Config{
		AllowWildcard:    true,
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
	}))

	// SuperTokens
	server.router.Use(authApi.SupertokensMiddleware)

	server.bindRoutes()

	log.Fatal(server.router.Run("0.0.0.0:8081"))
}
