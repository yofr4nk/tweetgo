package hashing_test

import (
	"testing"
	"tweetgo/pkg/hashing"
)

func TestHashPasswordShouldReturnErrorNil(t *testing.T) {
	_, err := hashing.HashPassword("pass")

	if err != nil {
		t.Errorf("Error should be nil but got: %v", err)
	}
}
