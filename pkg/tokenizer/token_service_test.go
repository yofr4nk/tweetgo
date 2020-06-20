package tokenizer_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/tokenizer"
)

func TestGenerateTokenSignPayloadAndResponseAToken(t *testing.T) {
	u := domain.User{
		ID:    primitive.ObjectID{},
		Name:  "TestName",
		Email: "fakeEmail",
	}
	tks := tokenizer.NewTokenService("key")
	token, err := tks.GenerateToken(u)
	if err != nil {
		t.Errorf("Expected error nil but got, %v", err)
	}

	if token == "" {
		t.Errorf("Expected a token value but got, %v", token)
	}
}

func TestGetAndValidateTokenDataShouldFailValidatingWhenTokenHasWrongFormat(t *testing.T) {
	tks := tokenizer.NewTokenService("key")
	_, isValid, err := tks.GetAndValidateTokenData("fakeToken")
	if err == nil {
		t.Errorf("Expected error %v, validating token but got nil", errors.New("invalid token format"))
	}

	if isValid == true {
		t.Errorf("Expected not valid token but got a valid token")
	}
}

func TestGetAndValidateTokenDataShouldFailValidatingWhenTokenIsNotValid(t *testing.T) {
	tks := tokenizer.NewTokenService("key")
	_, isValid, err := tks.GetAndValidateTokenData("Bearer fakeToken")
	if err == nil {
		t.Errorf("Expected token structure error but got nil")
	}

	if isValid == true {
		t.Errorf("Expected not valid token but got a valid token")
	}
}

func TestGetAndValidateTokenDataShouldValidateAndReturnUserData(t *testing.T) {
	u := domain.User{
		ID:    primitive.ObjectID{},
		Name:  "TestName",
		Email: "fakeEmail",
	}
	tks := tokenizer.NewTokenService("key")
	token, err := tks.GenerateToken(u)
	if err != nil {
		t.Errorf("Expected error nil but got, %v", err)

		return
	}

	if token == "" {
		t.Errorf("Expected a token value but got an empty token")

		return
	}

	userData, isValid, err := tks.GetAndValidateTokenData("Bearer " + token)
	if err != nil {
		t.Errorf("Expected a nil error but got %v, ", err)
	}

	if isValid == false {
		t.Errorf("Expected a valid token but got an invalid")
	}

	if userData.Email != u.Email {
		t.Errorf("Expected email %v, but got %v", u.Email, userData.Email)
	}

	if userData.ID != u.ID {
		t.Errorf("Expected id %v, but got %v", u.ID, userData.ID)
	}
}
