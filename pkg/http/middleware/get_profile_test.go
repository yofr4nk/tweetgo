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

type usrServiceMock struct {
	shouldFail bool
	usr        domain.User
}

func GetUserMock(usm usrServiceMock) func(email string) (domain.User, error) {
	return func(email string) (domain.User, error) {
		if usm.shouldFail {
			return domain.User{}, errors.New("usr service error mock")
		}

		return usm.usr, nil
	}
}

func TestGetProfileShouldFailGettingUserFromCtx(t *testing.T) {
	ucm := userCtxMock{
		shouldFail: true,
	}
	usm := usrServiceMock{}
	mw := middleware.GetProfile(getUserFromCtxMock(ucm), GetUserMock(usm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetProfileShouldFailGettingUserFroDB(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:       primitive.ObjectID{},
			Email:    "fakeEmail",
			Password: "fakePassword",
		},
	}
	usm := usrServiceMock{
		shouldFail: true,
	}
	mw := middleware.GetProfile(getUserFromCtxMock(ucm), GetUserMock(usm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestGetProfileShouldResponseWithStatusOk(t *testing.T) {
	u := domain.User{
		ID:       primitive.ObjectID{},
		Email:    "fakeEmail",
		Password: "fakePassword",
	}
	ucm := userCtxMock{
		usr: u,
	}
	usm := usrServiceMock{
		usr: u,
	}
	mw := middleware.GetProfile(getUserFromCtxMock(ucm), GetUserMock(usm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got: %v", r.Code)
	}
}
