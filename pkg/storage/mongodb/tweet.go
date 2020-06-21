package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
