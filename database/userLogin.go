package database

import (
	"github.com/yofr4nk/tweetgo/models"
	"golang.org/x/crypto/bcrypt"
)

// UserLogin make sure the user exists and check password
func UserLogin(email string, password string) (models.User, bool) {
	userData, findErr := FindUser(email)

	if findErr != nil {
		return userData, false
	}

	passwordBytes := []byte(password)
	dbPassword := []byte(userData.Password)

	err := bcrypt.CompareHashAndPassword(dbPassword, passwordBytes)

	if err != nil {
		return userData, false
	}

	return userData, true
}
