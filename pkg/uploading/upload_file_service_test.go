package uploading_test

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"tweetgo/pkg/uploading"
)

type uploadFileRepository struct {
	shouldFile   bool
	locationPath string
}

func (ufr uploadFileRepository) UploadFile(filePathName string, contentType string, body *bytes.Reader) (string, error) {
	if ufr.shouldFile {
		return "", errors.New("error uploading file")
	}

	return ufr.locationPath, nil
}

func TestUploadFileService_UploadFileShouldFailUploadingFileToStorage(t *testing.T) {
	ufl := uploadFileRepository{
		shouldFile: true,
	}

	ufs := uploading.NewUploadFileService(ufl)
	path := "../../assets/mocks/middleware/avatar_mock.jpg"
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
	mockFile, mockHeader, err := w.FormFile("avatar")
	if err != nil {
		t.Error(err)

		return
	}

	_, err = ufs.UploadFile("pathName", mockFile, mockHeader)
	if err == nil {
		t.Errorf("Expected error uploading file, but got nil")
	}

}

func TestUploadFileService_UploadFileShouldReturnLocationPathAfterUploadFile(t *testing.T) {
	ufl := uploadFileRepository{
		locationPath: "fakePath",
	}

	ufs := uploading.NewUploadFileService(ufl)
	path := "../../assets/mocks/middleware/avatar_mock.jpg"
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
	mockFile, mockHeader, err := w.FormFile("avatar")
	if err != nil {
		t.Error(err)

		return
	}

	locationPath, err := ufs.UploadFile("pathName", mockFile, mockHeader)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	if locationPath != "fakePath" {
		t.Errorf("Expected fakePath as locationn path, but got: %v", locationPath)
	}

}
