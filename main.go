package main

import (
	"log"
	"net/http"
	"os"
	"tweetgo/database"
	"tweetgo/pkg/finding"
	"tweetgo/pkg/http/rest"
	"tweetgo/pkg/loading"
	"tweetgo/pkg/saving"
	"tweetgo/pkg/storage/mongodb"
	"tweetgo/pkg/tokenizer"
)

func main() {
	if database.CheckingConnection() == false {
		log.Fatal("Missing database connection")

		return
	}

	securityKey := loading.GetSecurityKey()

	//Storage
	dbClient := mongodb.DBConnection()
	userStorage := mongodb.NewUserStorage(dbClient)
	tweetStorage := mongodb.NewTweetStorage(dbClient)

	//Services
	savingUserService := saving.NewUserService(userStorage)
	findingUserService := finding.NewUserService(userStorage)
	tokenizerService := tokenizer.NewTokenService(securityKey)
	savingTweetService := saving.NewTweetService(tweetStorage)
	findingTweetService := finding.NewTweetService(tweetStorage)

	r := rest.RouterManagement(savingUserService, findingUserService, tokenizerService, savingTweetService, findingTweetService)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
