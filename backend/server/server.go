package server

import (
	"context"
	"encoding/gob"
	"io"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

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
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatalf("Could not create the authenticator: %v", err)
	}

	cardsRepo, err := cards.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Cards repository: %v", err)
	}

	filesRepo, err := files.NewRepository(context.Background())
	if err != nil {
		log.Fatalf("Could not create the Files repository: %v", err)
	}

	cardsV1 := v1CardsApi.New(cardsRepo, filesRepo, authenticator)
	filesApi := filesApi.New(filesRepo)
	authApi := authApi.New(authenticator)

	router := gin.Default()

	return &Server{
		router:   router,
		services: []Service{cardsV1, filesApi, authApi},
	}
}

func (server *Server) Start(logFile io.Writer) {
	// session
	redisUrl := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisStore, err := redis.NewStore(10000, "tcp", redisUrl, redisPassword, []byte("TODO: SECRET KEY"))
	if err != nil {
		log.Fatalf("could not connect with the redis session: %v", err)
	}

	gob.Register(oauth2.Token{})
	server.router.Use(sessions.SessionsMany(
		[]string{"auth-session", "user-session"},
		redisStore,
	))

	// cors
	server.router.Use(cors.New(cors.Config{
		AllowWildcard:    true,
		AllowCredentials: true,
		AllowOrigins:     []string{"localhost"},
		AllowMethods:     []string{"GET", "PATCH", "POST", "DELETE"},
	}))

	server.bindRoutes()

	log.Fatal(server.router.Run("0.0.0.0:8081"))
}
