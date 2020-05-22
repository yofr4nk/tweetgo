package multimedia

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadFile get file info to upload in s3
func UploadFile(filePathName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	awsBucket := os.Getenv("BUCKET")

	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	sess, errSession := CreateAwsSession()

	if errSession != nil {
		log.Print("something went wrong creating aws session " + errSession.Error())

		return "", errSession
	}

	uploader := s3manager.NewUploader(sess)

	uploadInfo, errUploader := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(awsBucket),
		Key:         aws.String(filePathName),
		ContentType: aws.String(http.DetectContentType(buffer)),
		Body:        bytes.NewReader(buffer),
	})

	if errUploader != nil {
		log.Print("uploading error " + errUploader.Error())

		return "", errUploader
	}

	fmt.Println("Uploaded")

	return uploadInfo.Location, nil
}
