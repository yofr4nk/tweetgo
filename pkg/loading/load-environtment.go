package loading

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type EnvironmentBody struct {
	SecurityKey  string
	Bucket       string
	AwsAccessKey string
	AwsSecretKey string
}

func GetEnvironmentKeys() (EnvironmentBody, error) {
	//Checking environment before load .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return EnvironmentBody{}, errors.New("error loading environment variables " + err.Error())
		}
	}

	envBody := EnvironmentBody{
		SecurityKey:  os.Getenv("SECURITY_KEY"),
		Bucket:       os.Getenv("BUCKET"),
		AwsAccessKey: os.Getenv("ACCESS_KEY"),
		AwsSecretKey: os.Getenv("SECRET_KEY"),
	}

	return envBody, nil
}
