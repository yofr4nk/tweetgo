package tokenizer_test

import (
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
