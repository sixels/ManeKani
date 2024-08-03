package main

import (
	"log"

	"github.com/sixels/manekani/services/user-management/db"
	"github.com/sixels/manekani/services/user-management/service"
)

type DB struct {
	// database connection pgx
}

func main() {
	// TODO: get configuration from env
	db, err := db.New("postgres://manekani:secret@postgres-manekani/manekani-test?sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	state := service.NewState(db)

	log.Fatal(service.New(state).Listen(":9998"))
}
