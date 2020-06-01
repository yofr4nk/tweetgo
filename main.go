package main

import (
	"log"
	"net/http"
	"os"
	"tweetgo/database"
	"tweetgo/pkg/finding"
	"tweetgo/pkg/http/rest"
	"tweetgo/pkg/saving"
	"tweetgo/pkg/storage/mongodb"
)

func main() {
	if database.CheckingConnection() == false {
		log.Fatal("Missing database connection")

		return
	}

	//Storage
	dbClient := mongodb.DBConnection()
	userStorage := mongodb.NewUserStorage(dbClient)

	//Services
	savingUserService := saving.NewUserService(userStorage)
	findingUserService := finding.NewUserService(userStorage)

	r := rest.RouterManagement(savingUserService, findingUserService)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
