package loading

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetSecurityKey() string {
	//Checking environment before load .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading security environment variables", err.Error())
		}
	}

	securityKey := os.Getenv("SECURITY_KEY")

	return securityKey
}
