package main

import (
	"fmt"
	"log"
	"os"
	"path"

	server "github.com/sixels/manekani/server"

	"github.com/joho/godotenv"
)

// @title ManeKani API
// @version 1.0
// @description ManeKani API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email sixels@protonmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @tag.name health
// @tag.description API status

// @tag.name cards
// @tag.description Cards API related

// @tag.name kanji
// @tag.description Kanji cards actions

// @tag.name radical
// @tag.description Radical cards actions

// @tag.name vocabulary
// @tag.description Vocabulary cards actions

// @host localhost:8081
// @BasePath /
// @schemes http
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warn: could not load the .env file")
	}
	logFile := setLogFile()
	defer logFile.Close()

	fmt.Println("Starting the server")

	server.New().
		Start(logFile)
}

func setLogFile() *os.File {
	// TODO: check XDG_DATA_HOME directory too
	dataDir := os.Getenv("MANEKANI_DATA_HOME")
	if dataDir == "" {
		dataDir = "/data/"
	}

	logFile := path.Join(dataDir, "manekani.log")

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	log.SetOutput(f)
	return f
}
