package main

import (
	"log"

	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/handlers"
)

func main() {
	if database.CheckingConnection() == false {
		log.Fatal("Missing database connection")

		return
	}

	handlers.MainManagement()
}
