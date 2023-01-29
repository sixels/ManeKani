package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"

	_ "sixels.io/manekani/docs/manekani"

	authApi "sixels.io/manekani/server/api/auth"
	filesApi "sixels.io/manekani/server/api/files"
	usersApi "sixels.io/manekani/server/api/users"
	v1CardsApi "sixels.io/manekani/server/api/v1/cards"

	"sixels.io/manekani/services/auth"
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/ent"
	"sixels.io/manekani/services/files"
	"sixels.io/manekani/services/users"
)

type Service interface {
	SetupRoutes(router *gin.Engine)
}

type Server struct {
	router *gin.Engine

	services []Service
}

func New() *Server {
	entClient, err := ent.Connect()
	if err != nil {
		log.Fatalf("Could not connect with ManeKani database: %v", err)
	}
	cardsRepo := cards.NewRepository(entClient)
	usersRepo := users.NewRepository(entClient)

	if err := auth.StartAuthenticator(usersRepo); err != nil {
		log.Fatalf("Could not start the authenticator: %v", err)
	}

	filesRepo, err := files.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Files repository: %v", err)
	}

	cardsV1 := v1CardsApi.New(cardsRepo, filesRepo)
	filesApi := filesApi.New(filesRepo)
	usersApi := usersApi.New(usersRepo)

	router := gin.Default()

	return &Server{
		router:   router,
		services: []Service{cardsV1, filesApi, usersApi},
	}
}

func (server *Server) Start(logFile io.Writer) {
	// cors
	hostname, _ := os.Hostname()
	corsConfig := cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:8082", fmt.Sprintf("http://%s:8082", hostname), "http://192.168.15.9:8082"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
	}
	server.router.Use(cors.New(corsConfig))

	// SuperTokens
	server.router.Use(authApi.SupertokensMiddleware)

	server.bindRoutes()

	log.Fatal(server.router.Run(":8081"))
}
