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

type tokenValidationMock struct {
	shouldFail bool
	isValid    bool
	usr        domain.User
}

func (tvm tokenValidationMock) GetAndValidateTokenData(token string) (domain.User, bool, error) {
	if tvm.shouldFail {
		return domain.User{}, false, errors.New("token validation error")
	}

	if tvm.isValid == false {
		return domain.User{}, false, nil
	}

	return tvm.usr, true, nil
}

func TestCheckTokenShouldFailWhenValidationTokenFail(t *testing.T) {
	tokenService := tokenValidationMock{
		shouldFail: true,
	}
	mw := refmiddlewares.CheckToken(domain.SetUserToContext, tokenService, mockHandler)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestCheckTokenShouldFailWhenTokenIsNotValid(t *testing.T) {
	tokenService := tokenValidationMock{
		isValid: false,
	}
	mw := refmiddlewares.CheckToken(domain.SetUserToContext, tokenService, mockHandler)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestCheckTokenShouldResponseStatusOkWhenTokenIsValid(t *testing.T) {
	tokenService := tokenValidationMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeEmail",
		},
		isValid: true,
	}
	mw := refmiddlewares.CheckToken(domain.SetUserToContext, tokenService, mockHandler)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got: %v", r.Code)
	}
}
