package refmiddlewares_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"testing"
	"tweetgo/pkg/domain"
	refmiddlewares "tweetgo/pkg/http/middlewares"
)

type UserLoginMock struct {
	shouldFailGetUser         bool
	shouldFailComparePassword bool
	shouldFailGenerateToken   bool
	usr                       domain.User
	token                     string
}

func getUser(ulm UserLoginMock) func(email string) (domain.User, error) {
	return func(email string) (domain.User, error) {
		if ulm.shouldFailGetUser {
			return domain.User{}, errors.New("getUserFail")
		}

		return ulm.usr, nil
	}
}

func comparePassword(ulm UserLoginMock) func(password string, passwordHashed string) error {
	return func(password string, passwordHashed string) error {
		if ulm.shouldFailComparePassword {
			return errors.New("FailComparePassword")
		}

		return nil
	}
}

func (ulm UserLoginMock) GenerateToken(u domain.User) (string, error) {
	if ulm.shouldFailGenerateToken {
		return "", errors.New("FailGenerateToken")
	}

	return ulm.token, nil
}

func TestLoginShouldFailParsingBody(t *testing.T) {
	us := UserLoginMock{}
	mw := refmiddlewares.Login(getUser(us), comparePassword(us), us)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestLoginShouldFailWhenEmailIsEmpty(t *testing.T) {
	us := UserLoginMock{}
	mw := refmiddlewares.Login(getUser(us), comparePassword(us), us)
	body := strings.NewReader(`{"email": "" }`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestLoginShouldFailGettingUserFromDB(t *testing.T) {
	us := UserLoginMock{
		shouldFailGetUser: true,
	}
	mw := refmiddlewares.Login(getUser(us), comparePassword(us), us)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestLoginShouldFailComparingPassword(t *testing.T) {
	us := UserLoginMock{
		shouldFailComparePassword: true,
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeEmail",
		},
	}
	mw := refmiddlewares.Login(getUser(us), comparePassword(us), us)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestLoginShouldFailCreatingToken(t *testing.T) {
	us := UserLoginMock{
		shouldFailGenerateToken: true,
		usr: domain.User{
			ID:       primitive.ObjectID{},
			Email:    "fakeEmail",
			Password: "123456",
		},
	}
	mw := refmiddlewares.Login(getUser(us), comparePassword(us), us)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestLoginShouldResponseStatusOK(t *testing.T) {
	us := UserLoginMock{
		token: "fakeToken",
		usr: domain.User{
			ID:       primitive.ObjectID{},
			Email:    "fakeEmail",
			Password: "123456",
		},
	}
	mw := refmiddlewares.Login(getUser(us), comparePassword(us), us)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}
