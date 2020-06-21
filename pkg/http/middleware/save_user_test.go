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

type userSavingMock struct {
	shouldFail bool
	status     bool
}

func (usm *userSavingMock) SaveUser(u domain.User) (bool, error) {
	if usm.shouldFail {
		return false, errors.New("cannot save user")
	}

	return usm.status, nil
}

func TestSaveUserShouldFailGettingUserFromCtx(t *testing.T) {
	ucm := userCtxMock{
		shouldFail: true,
	}
	usm := userSavingMock{}

	mw := middleware.SaveUser(&usm, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveUserShouldFailWhenEmailIsEmpty(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{},
	}
	usm := userSavingMock{}

	mw := middleware.SaveUser(&usm, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveUserShouldFailWhenPasswordIsEmpty(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeEmail",
		},
	}
	usm := userSavingMock{}

	mw := middleware.SaveUser(&usm, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveUserShouldFailSavingUserInDB(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:       primitive.ObjectID{},
			Email:    "fakeEmail",
			Password: "fakePass",
		},
	}
	usm := userSavingMock{
		shouldFail: true,
	}

	mw := middleware.SaveUser(&usm, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveUserShouldFailSavingUserWhenGetStatusFalseFromDB(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:       primitive.ObjectID{},
			Email:    "fakeEmail",
			Password: "fakePass",
		},
	}
	usm := userSavingMock{
		status: false,
	}

	mw := middleware.SaveUser(&usm, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestSaveUserShouldResponseStatusCreated(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:       primitive.ObjectID{},
			Email:    "fakeEmail",
			Password: "fakePass",
		},
	}
	usm := userSavingMock{
		status: true,
	}

	mw := middleware.SaveUser(&usm, getUserFromCtxMock(ucm))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, but got: %v", r.Code)
	}
}
