package hashing

import "golang.org/x/crypto/bcrypt"

// HashPassword encrypt the string password
func HashPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
