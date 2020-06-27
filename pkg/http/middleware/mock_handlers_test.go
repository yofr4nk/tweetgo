package middleware_test

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"tweetgo/pkg/domain"
)

type userCtxMock struct {
	shouldFail bool
	usr        domain.User
}

type updateUserServiceMock struct {
	shouldFail   bool
	updateStatus bool
}

type uploadFileMockService struct {
	shouldFail bool
}

func mockServerHTTP(mw http.HandlerFunc, body *strings.Reader, params string, method string) *httptest.ResponseRecorder {
	w := httptest.NewRequest(method, "/"+params, body)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, w)

	return rr
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getUserFromCtxMock(usm userCtxMock) func(ctx context.Context) (domain.User, error) {
	return func(ctx context.Context) (domain.User, error) {
		if usm.shouldFail {
			return domain.User{}, errors.New("error getting user")
		}

		return usm.usr, nil
	}
}

func updateUserMock(usm updateUserServiceMock) func(u domain.User, ID string) (bool, error) {
	return func(u domain.User, ID string) (bool, error) {
		if usm.shouldFail {
			return usm.updateStatus, errors.New("update user error mock")
		}

		return usm.updateStatus, nil
	}
}

func (ufs uploadFileMockService) uploadFileMock(filePathName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	if ufs.shouldFail {
		return "", errors.New("error uploadingMock")
	}

	return "", nil
}
