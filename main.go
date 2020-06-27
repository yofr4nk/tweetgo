package main

import (
	"log"
	"net/http"
	"os"
	"tweetgo/pkg/deleting"
	"tweetgo/pkg/finding"
	"tweetgo/pkg/http/rest"
	"tweetgo/pkg/loading"
	"tweetgo/pkg/saving"
	"tweetgo/pkg/storage/mongodb"
	"tweetgo/pkg/storage/s3"
	"tweetgo/pkg/tokenizer"
	"tweetgo/pkg/uploading"
)

func main() {
	envKeys, err := loading.GetEnvironmentKeys()
	if err != nil {
		log.Fatal(err)

		return
	}

	//Storage
	dbClient := mongodb.DBConnection()

	userStorage := mongodb.NewUserStorage(dbClient)
	tweetStorage := mongodb.NewTweetStorage(dbClient)
	mediaStorage, err := s3.NewUploadStorage(envKeys.AwsAccessKey, envKeys.AwsSecretKey, envKeys.Bucket)
	if err != nil {
		log.Fatal(err)

		return
	}

	//Services
	savingUserService := saving.NewUserService(userStorage)
	findingUserService := finding.NewUserService(userStorage)
	tokenizerService := tokenizer.NewTokenService(envKeys.SecurityKey)
	savingTweetService := saving.NewTweetService(tweetStorage)
	findingTweetService := finding.NewTweetService(tweetStorage)
	deletingTweetService := deleting.NewTweetService(tweetStorage)
	uploadFileService := uploading.NewUploadFileService(mediaStorage)

	r := rest.RouterManagement(savingUserService, findingUserService, tokenizerService, savingTweetService, findingTweetService, deletingTweetService, uploadFileService)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
