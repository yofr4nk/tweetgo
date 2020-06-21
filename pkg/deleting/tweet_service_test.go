package deleting_test

import (
	"errors"
	"testing"
	"tweetgo/pkg/deleting"
)

type tweetRepositoryMock struct {
	ShouldFailDeleteTweet bool
}

func (trm tweetRepositoryMock) DeleteTweet(ID string, UserID string) error {
	if trm.ShouldFailDeleteTweet {
		return errors.New("error deleting tweet")
	}

	return nil
}

func TestTweetService_DeleteTweetShouldFailDeletingTweet(t *testing.T) {
	trm := tweetRepositoryMock{
		ShouldFailDeleteTweet: true,
	}

	ts := deleting.NewTweetService(trm)
	err := ts.DeleteTweet("12345", "fakeId")
	if err == nil {
		t.Errorf("Expected error deleting tweets but got nil")
	}
}

func TestTweetService_DeleteTweetShouldNotResponseErrorDeletingTweetInDB(t *testing.T) {
	trm := tweetRepositoryMock{}

	ts := deleting.NewTweetService(trm)
	err := ts.DeleteTweet("12345", "fakeId")
	if err != nil {
		t.Errorf("Expected error nil, but got: %v", err)
	}
}
