package security

import (
	"errors"
	"log"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/models"
)

// TokenData save info from tokenInfo
type TokenData struct {
	Email  string
	UserID string
}

var tokenInfo TokenData

// GetTokenData validate token provided
func GetTokenData(t string) (*models.Claim, bool, string, error) {
	//Checking environment before load .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading security environment variables", err.Error())
		}
	}

	securityKey := os.Getenv("SECURITY_KEY")
	claims := &models.Claim{}

	splitToken := strings.Split(t, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Token format invalid")
	}

	tokenToValidate := strings.TrimSpace(splitToken[1])

	tokn, err := jwt.ParseWithClaims(tokenToValidate, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(securityKey), nil
	})

	if err == nil {
		userFound, findErr := database.FindUserExists(claims.Email)

		if findErr != nil {
			return claims, false, string(""), findErr
		}

		if userFound == true {
			tokenInfo.Email = claims.Email
			tokenInfo.UserID = claims.ID.Hex()

			return claims, userFound, tokenInfo.UserID, nil
		}
	}

	if tokn.Valid == false {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
