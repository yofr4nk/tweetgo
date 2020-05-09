package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteTweet remove a tweet from a tweet ID and UserId
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	database := DBConnect.Database("tweetgo")
	tweetCollection := database.Collection("tweet")

	objTweetID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		log.Fatal("Something went wront parsing tweetID " + err.Error())

		return err
	}

	filter := bson.M{
		"_id":    objTweetID,
		"userid": UserID,
	}

	result, DeleteErr := tweetCollection.DeleteOne(ctx, filter)

	if result.DeletedCount == 0 {
		return errors.New("It could not possible to find tweet")
	}

	return DeleteErr

}
