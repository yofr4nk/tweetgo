package uploading

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

type UploadFileRepository interface {
	UploadFile(filePathName string, contentType string, body *bytes.Reader) (string, error)
}

type UploadFileService struct {
	repository UploadFileRepository
}

func NewUploadFileService(repository UploadFileRepository) *UploadFileService {
	return &UploadFileService{repository: repository}
}

func (ufs UploadFileService) UploadFile(filePathName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	body := bytes.NewReader(buffer)
	contentType := http.DetectContentType(buffer)

	locationPath, err := ufs.repository.UploadFile(filePathName, contentType, body)
	if err != nil {
		return "", err
	}

	return locationPath, nil
}
