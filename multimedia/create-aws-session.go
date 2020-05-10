package multimedia

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// CreateAwsSession build session to allow actions in aws client
func CreateAwsSession() (*session.Session, error) {
	awsAccessKey := os.Getenv("ACCESS_KEY")
	awsSecretKey := os.Getenv("SECRET_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})

	if err != nil {
		log.Fatal("something went wrong creating aws session " + err.Error())

		return nil, err
	}

	return sess, nil
}
