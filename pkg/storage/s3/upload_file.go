package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

type UploadStorage struct {
	session *session.Session
	bucket  string
}

func NewUploadStorage(accessKey string, secretKey string, bucket string) (*UploadStorage, error) {
	sess, err := CreateAwsSession(accessKey, secretKey)
	if err != nil {
		return nil, err
	}

	return &UploadStorage{
		session: sess,
		bucket:  bucket,
	}, nil
}

func (storage *UploadStorage) UploadFile(filePathName string, contentType string, body *bytes.Reader) (string, error) {
	uploader := s3manager.NewUploader(storage.session)

	uploadInfo, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(storage.bucket),
		Key:         aws.String(filePathName),
		ContentType: aws.String(contentType),
		Body:        body,
	})

	if err != nil {
		log.Print("uploading error " + err.Error())

		return "", err
	}

	return uploadInfo.Location, nil
}
