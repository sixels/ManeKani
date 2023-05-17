package server

import (
	"context"
	"encoding/gob"
	"io"
	"log"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"

	"github.com/sixels/manekani/core/adapters/cards"
	"github.com/sixels/manekani/core/adapters/tokens"
	_ "github.com/sixels/manekani/docs/manekani"

	cards_api "github.com/sixels/manekani/server/api/cards"
	files_api "github.com/sixels/manekani/server/api/files"
	tokens_api "github.com/sixels/manekani/server/api/tokens"
	users_api "github.com/sixels/manekani/server/api/users"
	auth_api "github.com/sixels/manekani/server/auth"
	"github.com/sixels/manekani/server/docs"

	"github.com/sixels/manekani/services/ent"
	cards_repo "github.com/sixels/manekani/services/ent/cards"
	"github.com/sixels/manekani/services/ent/token"
	"github.com/sixels/manekani/services/ent/users"
	"github.com/sixels/manekani/services/files"
)

type Service interface {
	SetupRoutes(router *gin.Engine)
}

type Server struct {
	router *gin.Engine

	services []Service
}

func New() *Server {
	gob.Register(user.User{})

	entRepo, err := ent.NewRepository()
	if err != nil {
		log.Fatalf("Could not connect with ManeKani database: %v", err)
	}

	tokensService := tokens.CreateAdapter(token.NewRepository(entRepo))
	authService := auth_api.NewAuthService(&tokensService)

	cardsRepo := cards_repo.NewRepository(entRepo)
	usersRepo := users.NewRepository(entRepo)
	if err != nil {
		log.Fatal(err)
	}

	filesRepo, err := files.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not setup the 'file' repository: %v", err)
	}

	tokensAPI := tokens_api.New(&tokensService, authService)
	cardsV1 := cards_api.New(cards.CreateAdapter(cardsRepo, filesRepo), filesRepo, usersRepo, authService)
	filesAPI := files_api.New(filesRepo)
	usersAPI := users_api.New(usersRepo, authService)

	docsHandler := docs.New("./docs/swagger.json")

	router := gin.Default()

	return &Server{
		router:   router,
		services: []Service{cardsV1, filesAPI, usersAPI, tokensAPI, authService, docsHandler},
	}
}

func (server *Server) Start(logFile io.Writer) {
	// clientURL := os.Getenv("MANEKANI_CLIENT_URL")
	// clientPort := os.Getenv("MANEKANI_CLIENT_PORT")
	serverPort := os.Getenv("MANEKANI_SERVER_PORT")

	// trustedOrigins := []string{
	// 	clientURL + ":" + clientPort,
	// 	"http://localhost:" + clientPort,
	// }

	// corsConfig := cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     trustedOrigins,
	// 	AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS", "PATCH"},
	// 	AllowHeaders: append([]string{"Content-Type"},
	// 		supertokens.GetAllCORSHeaders()...),
	// }
	// server.router.Use(cors.New(corsConfig))

	// // SuperTokens
	// server.router.Use(auth_api.SupertokensMiddleware)

	server.bindRoutes()

	log.Fatal(server.router.Run(":" + serverPort))
}
