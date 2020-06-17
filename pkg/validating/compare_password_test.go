package validating_test

import (
	"testing"
	"tweetgo/pkg/hashing"
	"tweetgo/pkg/validating"
)

func TestComparePasswordShouldReturnErrorWhenPasswordsAreNotEquals(t *testing.T) {
	passwordHashed, _ := hashing.HashPassword("fakePass")
	passwordToCompare := "fakePassword"

	err := validating.ComparePassword(passwordToCompare, passwordHashed)

	if err == nil {
		t.Errorf("passwords should not be equals")
	}
}

func TestComparePasswordShouldReturnANilErrorWhenPasswordsAreEquals(t *testing.T) {
	passwordHashed, _ := hashing.HashPassword("fakePass")
	passwordToCompare := "fakePass"

	err := validating.ComparePassword(passwordToCompare, passwordHashed)

	if err != nil {
		t.Errorf("passwords should be equals")
	}
}
