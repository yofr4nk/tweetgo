package middleware_test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/http/middleware"
)

func TestUpdateProfileShouldFailDecodingBody(t *testing.T) {
	us := updateUserServiceMock{}
	mw := middleware.UpdateProfile(domain.GetUserFromCtx, updateUserMock(us))
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUpdateProfileShouldFailGettingUserFromCtx(t *testing.T) {
	us := updateUserServiceMock{}
	usm := userCtxMock{
		shouldFail: true,
	}
	mw := middleware.UpdateProfile(getUserFromCtxMock(usm), updateUserMock(us))
	body := strings.NewReader(`{"name": "mockName"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUpdateProfileShouldFailUpdatingUser(t *testing.T) {
	u := domain.User{
		ID:    primitive.ObjectID{},
		Email: "fakeEmail",
	}
	us := updateUserServiceMock{
		shouldFail: true,
	}
	usm := userCtxMock{
		usr: u,
	}
	mw := middleware.UpdateProfile(getUserFromCtxMock(usm), updateUserMock(us))
	body := strings.NewReader(`{"name": "mockName"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUpdateProfileShouldFailUpdatingUserWhenUpdateStatusIsFalse(t *testing.T) {
	u := domain.User{
		ID:    primitive.ObjectID{},
		Email: "fakeEmail",
	}
	us := updateUserServiceMock{}
	usm := userCtxMock{
		usr: u,
	}
	mw := middleware.UpdateProfile(getUserFromCtxMock(usm), updateUserMock(us))
	body := strings.NewReader(`{"name": "mockName"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadGateway {
		t.Errorf("Expected status code %v, but got: %v", http.StatusBadGateway, r.Code)
	}
}

func TestUpdateProfileShouldResponseWithStatusOk(t *testing.T) {
	u := domain.User{
		ID:    primitive.ObjectID{},
		Email: "testEmail",
	}
	us := updateUserServiceMock{
		updateStatus: true,
	}
	usm := userCtxMock{
		usr: u,
	}
	mw := middleware.UpdateProfile(getUserFromCtxMock(usm), updateUserMock(us))
	body := strings.NewReader(`{"name": "mockName"}`)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got: %v", http.StatusOK, r.Code)
	}
}
