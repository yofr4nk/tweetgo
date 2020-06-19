package finding_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/finding"
)

type UserRepositoryMock struct {
	ShouldFailFindUserExist bool
	ShouldFailFindUser      bool
	count                   int64
	usr                     domain.User
}

func (srm UserRepositoryMock) FindUserExists(email string) (int64, error) {
	if srm.ShouldFailFindUserExist {
		return 0, errors.New("find user fail")
	}

	return srm.count, nil
}

func (srm UserRepositoryMock) FindUser(email string) (domain.User, error) {
	if srm.ShouldFailFindUser {
		return domain.User{}, errors.New("find user fail")
	}

	return srm.usr, nil
}

func TestUserService_GetUserShouldFailFindingUser(t *testing.T) {
	srm := UserRepositoryMock{
		ShouldFailFindUser: true,
	}
	us := finding.NewUserService(srm)

	_, err := us.GetUser("fakeEmail")
	if err == nil {
		t.Errorf("Expected error getting user, but got nil")
	}
}

func TestUserService_GetUserShouldReturnAnUserFound(t *testing.T) {
	u := domain.User{
		ID:       primitive.ObjectID{},
		Name:     "Fake Name",
		Email:    "fakeEmail",
		Password: "fakePassword",
	}
	srm := UserRepositoryMock{
		usr: u,
	}
	us := finding.NewUserService(srm)

	usr, err := us.GetUser("fakeEmail")
	if err != nil {
		t.Errorf("Expected error nil, but got: %v", err)
	}

	if usr.Name != u.Name {
		t.Errorf("Expected user name %v, but got: %v", u.Name, usr.Name)
	}
}

func TestUserService_FindUserExistsShouldFailValidatingIfExistsSomeUser(t *testing.T) {
	srm := UserRepositoryMock{
		ShouldFailFindUserExist: true,
	}
	us := finding.NewUserService(srm)

	_, err := us.FindUserExists("fakeEmail")
	if err == nil {
		t.Errorf("Expected error getting user, but got nil")
	}
}

func TestUserService_FindUserExistsShouldReturnTrueWhenUserExists(t *testing.T) {
	srm := UserRepositoryMock{
		count: 1,
	}
	us := finding.NewUserService(srm)

	userExists, err := us.FindUserExists("fakeEmail")
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}

	if userExists == false {
		t.Errorf("userExist should be true")
	}
}

func TestUserService_FindUserExistsShouldReturnFalseWhenUserDoesNotExist(t *testing.T) {
	srm := UserRepositoryMock{
		count: 0,
	}
	us := finding.NewUserService(srm)

	userExists, err := us.FindUserExists("fakeEmail")
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}

	if userExists {
		t.Errorf("userExist should be false")
	}
}
