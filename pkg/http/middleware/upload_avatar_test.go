package middleware_test

import (
	"bytes"
	"errors"
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

type uploadFileService struct {
	shouldFail bool
}

func (ufs uploadFileService) uploadFileMock(filePathName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	if ufs.shouldFail {
		return "", errors.New("error uploadingMock")
	}

	return "", nil
}

func TestUploadAvatarShouldFailGettingUserFromCtx(t *testing.T) {
	ucm := userCtxMock{
		shouldFail: true,
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileService{}

	mw := middleware.UploadAvatar(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadAvatarShouldFailGettingFileFromRequest(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileService{}

	mw := middleware.UploadAvatar(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	body := strings.NewReader(``)
	r := mockServerHTTP(mw, body, "", "POST")

	if r.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got: %v", r.Code)
	}
}

func TestUploadAvatarShouldFailUploadingFile(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileService{
		shouldFail: true,
	}

	mw := middleware.UploadAvatar(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("avatar", filepath.Base(path))
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

func TestUploadAvatarShouldFailUpdatingUser(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{
		shouldFail: true,
	}
	ufs := uploadFileService{}

	mw := middleware.UploadAvatar(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("avatar", filepath.Base(path))
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

func TestUploadAvatarShouldFailUpdatingUserWithStatusFalse(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{}
	ufs := uploadFileService{}

	mw := middleware.UploadAvatar(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("avatar", filepath.Base(path))
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

func TestUploadAvatarShouldResponseWithStatusOk(t *testing.T) {
	ucm := userCtxMock{
		usr: domain.User{
			ID:    primitive.ObjectID{},
			Email: "fakeName",
		},
	}
	usm := updateUserServiceMock{
		updateStatus: true,
	}
	ufs := uploadFileService{}

	mw := middleware.UploadAvatar(getUserFromCtxMock(ucm), updateUserMock(usm), ufs.uploadFileMock)
	path := "../../../assets/mocks/middleware/avatar_mock.jpg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	p, err := writer.CreateFormFile("avatar", filepath.Base(path))
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
