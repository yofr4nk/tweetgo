package middleware_test

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/http/middleware"
)

func TestUploadBannerShouldFailGettingUserFromCtx(t *testing.T) {
	ucm := userCtxMock{
		shouldFail: true,
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileMockService{}

	mw := middleware.UploadBanner(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadBannerShouldFailGettingFileFromRequest(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileMockService{}

	mw := middleware.UploadBanner(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadBannerShouldFailUploadingFile(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileMockService{
		shouldFail: true,
	}

	mw := middleware.UploadBanner(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("banner", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(p, file)
	writer.Close()

	w, _ := http.NewRequest("POST", "/", body)
	w.Header.Set("Content-Type", writer.FormDataContentType())

	r := httptest.NewRecorder()
	mw.ServeHTTP(r, w)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadBannerShouldFailUpdatingUser(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{
		shouldFail: true,
	}
	ufs := uploadFileMockService{}

	mw := middleware.UploadBanner(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("banner", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(p, file)
	writer.Close()

	w, _ := http.NewRequest("POST", "/", body)
	w.Header.Set("Content-Type", writer.FormDataContentType())

	r := httptest.NewRecorder()
	mw.ServeHTTP(r, w)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadBannerShouldFailUpdatingUserWithStatusFalse(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileMockService{}

	mw := middleware.UploadBanner(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("banner", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(p, file)
	writer.Close()

	w, _ := http.NewRequest("POST", "/", body)
	w.Header.Set("Content-Type", writer.FormDataContentType())

	r := httptest.NewRecorder()
	mw.ServeHTTP(r, w)

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadBannerShouldResponseWithStatusOk(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{
		updateStatus: true,
	}
	ufs := uploadFileMockService{}

	mw := middleware.UploadBanner(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("banner", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(p, file)
	writer.Close()

	w, _ := http.NewRequest("POST", "/", body)
	w.Header.Set("Content-Type", writer.FormDataContentType())

	r := httptest.NewRecorder()
	mw.ServeHTTP(r, w)

	if r.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got: %v", r.Code)
	}
}
