package main

import (
	"context"
	"log"
	"os"

	subjectsapi "github.com/sixels/manekani/server/api/cards"
	fileapi "github.com/sixels/manekani/server/api/files"
	tokenapi "github.com/sixels/manekani/server/api/tokens"
	userapi "github.com/sixels/manekani/server/api/users"
	"github.com/sixels/manekani/server/auth"
	"github.com/sixels/manekani/server/docs"

	"github.com/joho/godotenv"
	"github.com/sixels/manekani/core/adapters/cards"
	"github.com/sixels/manekani/core/adapters/tokens"
	server "github.com/sixels/manekani/server"
	"github.com/sixels/manekani/services/ent"
	card "github.com/sixels/manekani/services/ent/cards"
	"github.com/sixels/manekani/services/ent/token"
	"github.com/sixels/manekani/services/ent/users"
	file "github.com/sixels/manekani/services/files"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warn: could not load the .env file")
	}

	serverPort := os.Getenv("MANEKANI_SERVER_PORT")
	if serverPort == "" {
		serverPort = "10010"
	}

	entClient, err := ent.NewRepository()
	if err != nil {
		panic(err)
	}
	filesRepository, err := file.NewRepository(context.Background())
	if err != nil {
		panic(err)
	}

	tokenProvider := tokens.CreateAdapter(token.NewRepository(entClient))
	subjectProvider := cards.CreateAdapter(card.NewRepository(entClient), filesRepository)

	authenticator := auth.NewAuthService(&tokenProvider)
	authmws, err := authenticator.Middlewares()
	if err != nil {
		panic(err)
	}

	s := server.New(":" + serverPort).
		WithMiddleware(authmws...).
		WithService(docs.New()).
		WithService(tokenapi.New(tokenProvider)).
		WithService(subjectsapi.New(subjectProvider)).
		WithService(fileapi.New(filesRepository)).
		WithService(userapi.New(users.NewRepository(entClient)))

	// api.RegisterHandlers(s.Router(), s.API())

	log.Printf("Starting the server at :%s\n", serverPort)
	s.
		Start()
}
