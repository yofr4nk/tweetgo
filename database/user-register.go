package database

import (
	"context"
	"time"
	"tweetgo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveUser save the new user
func SaveUser(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := DBConnect.Database("tweetgo")
	userCollection := database.Collection("users")

	u.Password, _ = HashPassword(u.Password)

	result, err := userCollection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
