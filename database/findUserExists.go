package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/yofr4nk/tweetgo/models"
)

// FindUserExists get the user by email
func FindUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	database := DBConnect.Database("tweetgo")
	userCollection := database.Collection("users")

	condition := bson.M{"email":email}

	var result models.User

	err := userCollection.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}