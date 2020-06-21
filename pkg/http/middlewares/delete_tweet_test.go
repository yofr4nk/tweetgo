package refmiddlewares_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	rmiddlewares "tweetgo/pkg/http/middlewares"
)

type deleteTweetServiceMock struct {
	shouldFail bool
}

func (gts deleteTweetServiceMock) DeleteTweetMock(ID string, UserID string) error {
	if gts.shouldFail {
		return errors.New("delete tweet error mock")
	}

	return nil
}

func TestDeleteTweetShouldFailWhenUserIdIsEmpty(t *testing.T) {
	dts := deleteTweetServiceMock{}

	mw := rmiddlewares.DeleteTweet(dts.DeleteTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "DELETE")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestDeleteTweetShouldFailWhenTweetIdIsEmpty(t *testing.T) {
	dts := deleteTweetServiceMock{}

	mw := rmiddlewares.DeleteTweet(dts.DeleteTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=fakeUserId", "DELETE")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestDeleteTweetShouldFailDeletingTweet(t *testing.T) {
	dts := deleteTweetServiceMock{
		shouldFail: true,
	}

	mw := rmiddlewares.DeleteTweet(dts.DeleteTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=fakeUserId&id=fakeId", "DELETE")

	if r.Code != http.StatusBadGateway {
		t.Errorf("Expected status code %v, but got: %v", http.StatusBadGateway, r.Code)
	}
}

func TestDeleteTweetShouldResponseWithStatusOk(t *testing.T) {
	dts := deleteTweetServiceMock{}

	mw := rmiddlewares.DeleteTweet(dts.DeleteTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=fakeUserId&id=fakeId", "DELETE")

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got: %v", http.StatusOK, r.Code)
	}
}
