package database

import (
	"context"
	"log"
	"time"

	"github.com/yofr4nk/tweetgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTweets obtain the list of tweets related to an userId
func GetTweets(userID string, page int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	database := DBConnect.Database("tweetgo")
	tweetCollection := database.Collection("tweet")

	var result []*models.Tweet

	filter := bson.M{
		"userid": userID,
	}

	queryOtions := options.Find()
	queryOtions.SetLimit(20)
	queryOtions.SetSort(bson.D{{Key: "date", Value: -1}})
	queryOtions.SetSkip((page - 1) * 20)

	data, err := tweetCollection.Find(ctx, filter, queryOtions)

	if err != nil {
		log.Fatal("Something went wront getting tweets " + err.Error())

		return result, false
	}

	for data.Next(context.TODO()) {
		var tweetRow models.Tweet

		err := data.Decode(&tweetRow)

		if err != nil {
			log.Fatal(err.Error())

			return result, false
		}

		result = append(result, &tweetRow)
	}

	return result, true
}
