package middleware_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/http/middleware"
)

type getTweetServiceMock struct {
	shouldFail bool
	tweets     domain.Tweets
}

func (gts getTweetServiceMock) getTweetMock(userID string, page int64) (domain.Tweets, error) {
	if gts.shouldFail {
		return domain.Tweets{}, errors.New("get tweet error mock")
	}

	return gts.tweets, nil
}

func TestGetTweetsShouldFailGettingUserIdParam(t *testing.T) {
	gts := getTweetServiceMock{}
	mw := middleware.GetTweets(gts.getTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "GET")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetTweetsShouldFailGettingPageParam(t *testing.T) {
	gts := getTweetServiceMock{}
	mw := middleware.GetTweets(gts.getTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=12345", "GET")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetTweetsShouldFailParsingPageParamToInt(t *testing.T) {
	gts := getTweetServiceMock{}
	mw := middleware.GetTweets(gts.getTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=12345&page=fakeParam", "GET")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetTweetsShouldFailWhenPageIsLessThanOne(t *testing.T) {
	gts := getTweetServiceMock{}
	mw := middleware.GetTweets(gts.getTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=12345&page=0", "GET")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetTweetsShouldFailGettingUser(t *testing.T) {
	gts := getTweetServiceMock{
		shouldFail: true,
	}
	mw := middleware.GetTweets(gts.getTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=12345&page=1", "GET")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetTweetsShouldResponseWithStatusOk(t *testing.T) {
	gts := getTweetServiceMock{
		tweets: domain.Tweets{
			domain.Tweet{
				ID:      primitive.ObjectID{},
				UserID:  "12345",
				Message: "mock tweet",
				Date:    0,
			},
		},
	}
	mw := middleware.GetTweets(gts.getTweetMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "?userId=12345&page=1", "GET")

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got: %v", http.StatusOK, r.Code)
	}
}
