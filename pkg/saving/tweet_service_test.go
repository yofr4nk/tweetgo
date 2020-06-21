package saving_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/saving"
)

type tweetServiceMock struct {
	saveTweetShouldFail bool
	saveTweetStatus     bool
}

func (tsm tweetServiceMock) SaveTweet(t domain.Tweet) (bool, error) {
	if tsm.saveTweetShouldFail {
		return tsm.saveTweetStatus, errors.New("error test")
	}

	return tsm.saveTweetStatus, nil
}

func TestTweetService_SaveTweetShouldFailSavingInDB(t *testing.T) {
	tweetData := domain.Tweet{
		ID:      primitive.ObjectID{},
		UserID:  "fakeId",
		Message: "fakeMessage",
		Date:    0,
	}
	tsm := tweetServiceMock{
		saveTweetShouldFail: true,
	}
	us := saving.NewTweetService(tsm)

	_, err := us.SaveTweet(tweetData)
	if err == nil {
		t.Errorf("Expected error saving tweet, but got %v", err)
	}
}

func TestTweetService_SaveTweetShouldFailWithStatusFalseSavingInDB(t *testing.T) {
	tweetData := domain.Tweet{
		ID:      primitive.ObjectID{},
		UserID:  "fakeId",
		Message: "fakeMessage",
		Date:    0,
	}
	tsm := tweetServiceMock{}
	us := saving.NewTweetService(tsm)

	status, _ := us.SaveTweet(tweetData)
	if status {
		t.Errorf("Expected false status saving tweet, but got true")
	}
}

func TestTweetService_SaveTweetShouldReturnStatusTrueSavingInDB(t *testing.T) {
	tweetData := domain.Tweet{
		ID:      primitive.ObjectID{},
		UserID:  "fakeId",
		Message: "fakeMessage",
		Date:    0,
	}
	tsm := tweetServiceMock{
		saveTweetStatus: true,
	}
	us := saving.NewTweetService(tsm)

	status, _ := us.SaveTweet(tweetData)
	if status == false {
		t.Errorf("Expected true status saving tweet, but got false")
	}
}
