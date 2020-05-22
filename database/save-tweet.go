package database

import (
	"context"
	"time"
	"tweetgo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveTweet save the new tweet
func SaveTweet(t models.Tweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := DBConnect.Database("tweetgo")
	tweetCollection := database.Collection("tweet")

	tweet := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := tweetCollection.InsertOne(ctx, tweet)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
