package server

import (
	"context"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"

	_ "sixels.io/manekani/docs/manekani"

	auth_api "sixels.io/manekani/server/api/auth"
	cards_api "sixels.io/manekani/server/api/cards"
	files_api "sixels.io/manekani/server/api/files"
	users_api "sixels.io/manekani/server/api/users"

	"sixels.io/manekani/services/auth"
	"sixels.io/manekani/services/ent"
	"sixels.io/manekani/services/ent/cards"
	"sixels.io/manekani/services/ent/users"
	"sixels.io/manekani/services/files"
	"sixels.io/manekani/services/jwt"
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
	cardsRepo := cards.NewRepository(entRepo)
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

	cardsV1 := cards_api.New(cardsRepo, filesRepo, usersRepo, jwtService)
	filesAPI := files_api.New(filesRepo)
	usersAPI := users_api.New(usersRepo, jwtService)

	router := gin.Default()

	return &Server{
		router:   router,
		services: []Service{cardsV1, filesAPI, usersAPI},
	}
}

func (server *Server) Start(logFile io.Writer) {
	// cors
	hostname, _ := os.Hostname()
	corsConfig := cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:8082", fmt.Sprintf("http://%s:8082", hostname), "http://192.168.15.9:8082"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS", "PATCH"},
		AllowHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
	}
	server.router.Use(cors.New(corsConfig))

	// SuperTokens
	server.router.Use(auth_api.SupertokensMiddleware)

	server.bindRoutes()

	log.Fatal(server.router.Run(":8081"))
}
