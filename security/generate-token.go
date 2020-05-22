package security

import (
	"log"
	"os"
	"time"
	"tweetgo/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// GenerateToken create a token based on user info
func GenerateToken(u models.User) (string, error) {
	//Checking environment before load .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading security environment variables", err.Error())
		}
	}

	securityKey := os.Getenv("SECURITY_KEY")

	securityKeyBytes := []byte(securityKey)

	payload := jwt.MapClaims{
		"email":        u.Email,
		"name":         u.Name,
		"lastname":     u.LastName,
		"userbirthday": u.UserBirthday,
		"biography":    u.Biography,
		"location":     u.Location,
		"website":      u.WebSite,
		"_id":          u.ID.Hex(),
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenToSign := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := tokenToSign.SignedString(securityKeyBytes)

	if err != nil {
		return token, err
	}

	return token, nil
}
