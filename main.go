package main

import (
	"log"
	"tweetgo/database"
	"tweetgo/handlers"
)

func main() {
	if database.CheckingConnection() == false {
		log.Fatal("Missing database connection")

		return
	}

	handlers.MainManagement()
}
