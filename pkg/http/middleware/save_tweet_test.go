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

type tweetServiceMock struct {
	shouldFail  bool
	tweetStatus bool
}

func (usm *tweetServiceMock) SaveTweet(t domain.Tweet) (bool, error) {
	if usm.shouldFail {
		return false, errors.New("mock error")
	}

	return usm.tweetStatus, nil
}

func TestSaveTweetShouldFailParsingBody(t *testing.T) {
	ucm := userCtxMock{}
	tsm := tweetServiceMock{}

	mw := middleware.SaveTweet(tsm.SaveTweet, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveTweetShouldFailGettingUserFromCtx(t *testing.T) {
	ucm := userCtxMock{
		shouldFail: true,
	}
	tsm := tweetServiceMock{}

	mw := middleware.SaveTweet(tsm.SaveTweet, getUserFromCtxMock(ucm))
	body := strings.NewReader(`{"message": "fakeMessage"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveTweetShouldFailSavingTweet(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeEmail",
		},
	}
	tsm := tweetServiceMock{
		shouldFail: true,
	}

	mw := middleware.SaveTweet(tsm.SaveTweet, getUserFromCtxMock(ucm))
	body := strings.NewReader(`{"message": "fakeMessage"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveTweetShouldFailWithStatusFalseSavingTweet(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeEmail",
		},
	}
	tsm := tweetServiceMock{}

	mw := middleware.SaveTweet(tsm.SaveTweet, getUserFromCtxMock(ucm))
	body := strings.NewReader(`{"message": "fakeMessage"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveTweetShouldResponseWithStatusCreatedSavingTweet(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeEmail",
		},
	}
	tsm := tweetServiceMock{
		tweetStatus: true,
	}

	mw := middleware.SaveTweet(tsm.SaveTweet, getUserFromCtxMock(ucm))
	body := strings.NewReader(`{"message": "fakeMessage"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusCreated {
		t.Errorf("Expected status code %v, but got: %v", http.StatusCreated, r.Code)
	}
}
