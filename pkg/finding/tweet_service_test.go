package finding_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/finding"
)

type tweetRepositoryMock struct {
	ShouldFailGetTweets bool
	tweets              domain.Tweets
}

func (trm tweetRepositoryMock) GetTweets(userID string, page int64) (domain.Tweets, error) {
	if trm.ShouldFailGetTweets {
		return domain.Tweets{}, errors.New("get tweets error mock")
	}

	return trm.tweets, nil
}

func TestTweetService_GetTweetsShouldFailGettingTweetsFromDB(t *testing.T) {
	trm := tweetRepositoryMock{
		ShouldFailGetTweets: true,
	}

	ts := finding.NewTweetService(trm)
	_, err := ts.GetTweets("", 1)
	if err == nil {
		t.Errorf("Expected error getting tweets, but got nil")
	}
}

func TestTweetService_GetTweetsShouldReturnTweets(t *testing.T) {
	tweets := domain.Tweets{
		domain.Tweet{
			ID:      primitive.ObjectID{},
			UserID:  "12345",
			Message: "mock tweet",
			Date:    0,
		},
	}
	trm := tweetRepositoryMock{
		tweets: tweets,
	}

	ts := finding.NewTweetService(trm)
	tweetsFromService, err := ts.GetTweets("", 1)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	if len(tweetsFromService) != 1 {
		t.Errorf("Expected one tweet, but got: %v", tweetsFromService)
	}
}
