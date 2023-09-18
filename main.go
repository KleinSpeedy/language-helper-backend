package main

import (
	"log"
	"os"
	"strconv"

	"github.com/KleinSpeedy/language-helper-backend/api"
	"github.com/KleinSpeedy/language-helper-backend/database"
)

func main() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")

	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		log.Fatalf("Error converting: %s", err.Error())
	}

	dbctr := database.NewController(name, user, pw, port)
	err, ok := dbctr.OpenConnection()

	if !ok {
		log.Fatalf(err.Error())
	} else {
		log.Println("Connected to database")
	}

	// create server and run
	server := api.NewServer(8000)
	log.Fatal(server.Start())
}
