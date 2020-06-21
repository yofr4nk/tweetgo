package saving_test

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/saving"
)

type UserServiceMock struct {
	shouldFailSaveUser   bool
	shouldFailUpdateUser bool
	saveUserStatus       bool
}

func (usm UserServiceMock) SaveUser(u domain.User) (bool, error) {
	if usm.shouldFailSaveUser {
		return usm.saveUserStatus, errors.New("saving user error")
	}

	return usm.saveUserStatus, nil
}

func (usm UserServiceMock) UpdateUser(usrData domain.UserDataContainer, ID string) error {
	if usm.shouldFailUpdateUser {
		return errors.New("updating user error")
	}

	return nil
}

func TestUserService_SaveUserShouldFailSavingUserToDB(t *testing.T) {
	usm := UserServiceMock{
		shouldFailSaveUser: true,
	}
	u := domain.User{
		ID:           primitive.ObjectID{},
		Name:         "testName",
		LastName:     "fakeLastName",
		UserBirthday: "1998/06/01",
		Email:        "fakeEmail",
		Avatar:       "avatar/fake/url",
		Banner:       "banner/fake/url",
		Biography:    "testBio",
		Location:     "loc",
		WebSite:      "webSiteTest",
	}

	us := saving.NewUserService(usm)

	_, err := us.SaveUser(u)
	if err == nil {
		t.Errorf("Expected error saving user, but got nil")
	}
}

func TestUserService_SaveUserShouldFailWithStatusFalseSavingUserToDB(t *testing.T) {
	usm := UserServiceMock{
		saveUserStatus: false,
	}
	u := domain.User{
		ID:    primitive.ObjectID{},
		Name:  "testName",
		Email: "fakeEmail",
	}

	us := saving.NewUserService(usm)

	status, err := us.SaveUser(u)
	if err == nil || status {
		t.Errorf("Expected error saving user, but got nil or status true")
	}
}

func TestUserService_SaveUserShouldGetStatusTrueSavingUserInDB(t *testing.T) {
	usm := UserServiceMock{
		saveUserStatus: true,
	}
	u := domain.User{
		ID:           primitive.ObjectID{},
		Name:         "testName",
		LastName:     "fakeLastName",
		UserBirthday: "1998/06/01",
		Email:        "fakeEmail",
		Avatar:       "avatar/fake/url",
		Banner:       "banner/fake/url",
		Biography:    "testBio",
		Location:     "loc",
		WebSite:      "webSiteTest",
	}

	us := saving.NewUserService(usm)

	status, err := us.SaveUser(u)
	if err != nil || status == false {
		t.Errorf("Expected error nil but got %v or status false", err)
	}
}

func TestUserService_UpdateUserShouldFailSavingUserToDB(t *testing.T) {
	usm := UserServiceMock{
		shouldFailUpdateUser: true,
	}
	u := domain.User{
		ID:    primitive.ObjectID{},
		Email: "fakeEmail",
	}

	us := saving.NewUserService(usm)

	_, err := us.UpdateUser(u, "fakeID")
	if err == nil {
		t.Errorf("Expected error updating user, but got nil")
	}
}

func TestUserService_UpdateUserShouldUpdateUserInDB(t *testing.T) {
	usm := UserServiceMock{}
	u := domain.User{
		ID:           primitive.ObjectID{},
		Name:         "testName",
		LastName:     "fakeLastName",
		UserBirthday: "1998/06/01",
		Email:        "fakeEmail",
		Avatar:       "avatar/fake/url",
		Banner:       "banner/fake/url",
		Biography:    "testBio",
		Location:     "loc",
		WebSite:      "webSiteTest",
	}

	us := saving.NewUserService(usm)

	status, err := us.UpdateUser(u, "fakeID")
	if err != nil || status == false {
		t.Errorf("Expected error updating user, but got nil")
	}
}
