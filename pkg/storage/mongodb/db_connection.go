package mongodb

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// DBConnection returns the client connection
func DBConnection() *mongo.Client {

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
		log.Print("Database Connection Error, " + err.Error())
		return client
	}

	checkError := client.Ping(context.TODO(), nil)

	if checkError != nil {
		log.Print("Checking Connection Error, " + checkError.Error())
		return client
	}

	fmt.Println("Connection successful")

	return client

}
