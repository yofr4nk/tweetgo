package database

import (
	"context"
	"time"
	"tweetgo/models"

	"go.mongodb.org/mongo-driver/bson"
)

// FindUser get the user by email
func FindUser(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	database := DBConnect.Database("tweetgo")
	userCollection := database.Collection("users")

	condition := bson.M{"email": email}

	var u models.User

	err := userCollection.FindOne(ctx, condition).Decode(&u)

	if err != nil {
		return u, err
	}

	return u, nil
}
