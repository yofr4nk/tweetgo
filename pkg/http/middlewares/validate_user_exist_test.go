package refmiddlewares_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"tweetgo/pkg/domain"
	refmiddlewares "tweetgo/pkg/http/middlewares"
)

type UserServiceMock struct {
	shouldFail bool
	userExist  bool
}

func (usm *UserServiceMock) FindUserExists(email string) (bool, error) {
	if usm.shouldFail {
		return false, errors.New("mock error")
	}

	if usm.userExist {
		return true, nil
	}

	return false, nil
}

func TestValidateUserExistShouldFailParsingBody(t *testing.T) {
	us := UserServiceMock{}
	mw := refmiddlewares.ValidateUserExist(&us, domain.SetUserToContext, mockHandler)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestValidateUserExistShouldFailObtainingEmail(t *testing.T) {
	us := UserServiceMock{}
	mw := refmiddlewares.ValidateUserExist(&us, domain.SetUserToContext, mockHandler)
	body := strings.NewReader(`{"email": "" }`)
	r := mockServerHTTP(mw, body)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestValidateUserExistShouldFailWhenEmailAlreadyExist(t *testing.T) {
	us := UserServiceMock{
		userExist: true,
	}
	mw := refmiddlewares.ValidateUserExist(&us, domain.SetUserToContext, mockHandler)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestValidateUserExistShouldFailValidatingIfUserExist(t *testing.T) {
	us := UserServiceMock{
		shouldFail: true,
	}
	mw := refmiddlewares.ValidateUserExist(&us, domain.SetUserToContext, mockHandler)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestValidateUserExistShouldResponseWhitStatusOk(t *testing.T) {
	us := UserServiceMock{}
	mw := refmiddlewares.ValidateUserExist(&us, domain.SetUserToContext, mockHandler)
	body := strings.NewReader(`{"email": "fakeEmail" }`)
	r := mockServerHTTP(mw, body)

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got: %v", r.Code)
	}
}
