package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnect returns the client connection
var DBConnect = ConnectDatabaBase()

// ConnectDatabaBase returns the client connection
func ConnectDatabaBase() *mongo.Client {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading environment variables", err.Error())
		}
	}

	userDatabase := os.Getenv("USER_DATABASE")
	password := os.Getenv("PASSWORD")
	databasePath := os.Getenv("DATABASE_PATH")

	var clientOptions = options.Client().ApplyURI("mongodb+srv://" + userDatabase + ":" + password + "@" + databasePath)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Database Connection Error, " + err.Error())
		return client
	}

	checkError := client.Ping(context.TODO(), nil)

	if checkError != nil {
		log.Fatal("Checking Connection Error, " + checkError.Error())
		return client
	}

	fmt.Println("Connection successful")

	return client

}

// CheckingConnection check if the database connection is available
func CheckingConnection() bool {
	checkError := DBConnect.Ping(context.TODO(), nil)

	if checkError != nil {
		return false
	}

	return true
}
