package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"tweetgo/pkg/domain"
)

type TweetStorage struct {
	collection *mongo.Collection
}

const (
	tweetCollection string = "tweet"
)

func NewTweetStorage(db *mongo.Client) *TweetStorage {
	database := db.Database(dbName)
	tc := database.Collection(tweetCollection)

	return &TweetStorage{collection: tc}
}

// SaveTweet save the new tweet
func (storage *TweetStorage) SaveTweet(t domain.Tweet) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tweet := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	_, err := storage.collection.InsertOne(ctx, tweet)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (storage *TweetStorage) GetTweets(userID string, page int64) (domain.Tweets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var tweets domain.Tweets

	filter := bson.M{
		"userid": userID,
	}

	queryOptions := options.Find()
	queryOptions.SetLimit(20)
	queryOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	queryOptions.SetSkip((page - 1) * 20)

	data, err := storage.collection.Find(ctx, filter, queryOptions)
	if err != nil {
		log.Print("Something went wrong getting tweets " + err.Error())

		return tweets, err
	}

	for data.Next(context.TODO()) {
		var tweetRow domain.Tweet

		err := data.Decode(&tweetRow)
		if err != nil {
			log.Fatal(err.Error())

			return tweets, err
		}

		tweets = append(tweets, tweetRow)
	}

	return tweets, nil
}
