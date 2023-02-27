package server

import (
	"context"
	"encoding/gob"
	"io"
	"log"
	"os"
	"os/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"

	"github.com/sixels/manekani/core/services/cards"
	_ "github.com/sixels/manekani/docs/manekani"

	auth_api "github.com/sixels/manekani/server/api/auth"
	cards_api "github.com/sixels/manekani/server/api/cards"
	files_api "github.com/sixels/manekani/server/api/files"
	users_api "github.com/sixels/manekani/server/api/users"

	"github.com/sixels/manekani/services/auth"
	"github.com/sixels/manekani/services/ent"
	cards_repo "github.com/sixels/manekani/services/ent/cards"
	"github.com/sixels/manekani/services/ent/users"
	"github.com/sixels/manekani/services/files"
	"github.com/sixels/manekani/services/jwt"
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

	jwtService := jwt.CreateService(os.Getenv("TOKEN_SIGN_KEY"))
	entRepo, err := ent.NewRepository()
	if err != nil {
		log.Fatalf("Could not connect with ManeKani database: %v", err)
	}
	cardsRepo := cards_repo.NewRepository(entRepo)
	usersRepo, err := users.NewRepository(context.Background(), entRepo, jwtService)
	if err != nil {
		log.Fatal(err)
	}

	if err := auth.StartAuthenticator(usersRepo); err != nil {
		log.Fatalf("Could setup the 'user' repository: %v", err)
	}

	filesRepo, err := files.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not setup the 'file' repository: %v", err)
	}

	cardsV1 := cards_api.New(cards.NewService(cardsRepo, filesRepo), filesRepo, usersRepo, jwtService)
	filesAPI := files_api.New(filesRepo)
	usersAPI := users_api.New(usersRepo, jwtService)

	router := gin.Default()

	return &Server{
		router:   router,
		services: []Service{cardsV1, filesAPI, usersAPI},
	}
}

func (server *Server) Start(logFile io.Writer) {
	clientURL := os.Getenv("MANEKANI_CLIENT_URL")
	clientPort := os.Getenv("MANEKANI_CLIENT_URL")
	serverPort := os.Getenv("MANEKANI_SERVER_PORT")

	trustedOrigins := []string{
		clientURL + ":" + clientPort,
		"http://localhost:" + clientPort,
	}

	corsConfig := cors.Config{
		AllowCredentials: true,
		AllowOrigins:     trustedOrigins,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS", "PATCH"},
		AllowHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
	}
	server.router.Use(cors.New(corsConfig))

	// SuperTokens
	server.router.Use(auth_api.SupertokensMiddleware)

	server.bindRoutes()

	log.Fatal(server.router.Run(":" + serverPort))
}
