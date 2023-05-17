package main

import (
	"fmt"
	"log"
	"os"
	"path"

	server "github.com/sixels/manekani/server"

	"github.com/joho/godotenv"
)

//	@title			ManeKani API
//	@version		1.0
//	@description	ManeKani API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	sixels@protonmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@tag.name			health
//	@tag.description	API status

//	@tag.name			cards
//	@tag.description	Cards API related

//	@tag.name			kanji
//	@tag.description	Kanji cards actions

//	@tag.name			radical
//	@tag.description	Radical cards actions

//	@tag.name			vocabulary
//	@tag.description	Vocabulary cards actions

//	@tag.name			user
//	@tag.description	User related

//	@tag.name			tokens
//	@tag.description	API token related

//	@securitydefinitions.apikey	Login
//	@in							cookie
//	@name						ory_kratos_session
//	@description				Login at http://127.0.0.1:4455/login and copy the contents of the 'ory_kratos_session' cookie

//	@securityDefinitions.apikey	ApiKey
//	@in							header
//	@name						Authorization
//	@description				API key authentication

//	@host		127.0.0.1:8080
//	@BasePath	/
//	@schemes	http
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warn: could not load the .env file")
	}
	logFile := setLogFile()
	if logFile != nil {
		defer logFile.Close()
	}

	fmt.Println("Starting the server")

	server.New().
		Start(logFile)
}

func setLogFile() *os.File {
	// TODO: check XDG_DATA_HOME directory too
	dataDir := os.Getenv("MANEKANI_DATA_HOME")
	if dataDir != "" {
		logFile := path.Join(dataDir, "manekani.log")
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening log file: %v", err)
		}
		log.SetOutput(f)
		return f
	}
	return nil
}
