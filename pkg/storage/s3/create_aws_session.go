package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func CreateAwsSession(awsAccessKey string, awsSecretKey string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})

	if err != nil {
		log.Print("something went wrong creating aws session " + err.Error())

		return nil, err
	}

	return sess, nil
}
