package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FindUserExists get the user by email
func FindUserExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	database := DBConnect.Database("tweetgo")
	userCollection := database.Collection("users")

	condition := bson.M{"email": email}

	count, err := userCollection.CountDocuments(ctx, condition)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
