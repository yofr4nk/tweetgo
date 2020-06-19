package validating

import "golang.org/x/crypto/bcrypt"

func ComparePassword(password string, hashedPassword string) error {
	passwordToCompare := []byte(password)
	dbPassword := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(dbPassword, passwordToCompare)

	return err
}
