package domain_test

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"tweetgo/pkg/domain"
)

func TestGetUserFromCtxShouldFailWhenContextIsEmpty(t *testing.T) {
	_, err := domain.GetUserFromCtx(nil)
	if err.Error() != "context is empty" {
		t.Errorf("Expected context is empty message error, but got: %v", err.Error())
	}
}

func TestGetUserFromCtxReturnErrorIfKeyNotPresent(t *testing.T) {
	_, err := domain.GetUserFromCtx(context.TODO())
	if err.Error() != "user not found in context" {
		t.Errorf("Expected context is empty message error, but got: %v", err.Error())
	}
}

func Test_FileSummaryShouldReturnValueIfIsPresent(t *testing.T) {
	usr := domain.User{
		ID:   primitive.ObjectID{},
		Name: "userFake",
	}
	ctx := domain.SetUserToContext(context.TODO(), usr)

	userCtx, err := domain.GetUserFromCtx(ctx)
	if err != nil {
		t.Errorf("Expecting no errors, but got: %v", err.Error())
	}

	if userCtx.Name != usr.Name {
		t.Errorf("Expecting %v, but got: %v", usr.Name, userCtx.Name)
	}
}
