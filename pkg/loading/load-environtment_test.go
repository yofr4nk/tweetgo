package loading_test

import (
	"os"
	"testing"
	"tweetgo/pkg/loading"
)

func TestGetEnvironmentKeysShouldFailLoadingEnvironmentVars(t *testing.T) {
	_, err := loading.GetEnvironmentKeys()

	if err == nil {
		t.Errorf("Expected error loading environment vars, but got nil")

	}
}

func TestGetEnvironmentKeysShouldSetNewEnvironmentVars(t *testing.T) {
	securityKey := "fakeSecurityKey"
	bucket := "fakeBucket"

	os.Setenv("APP_ENV", "production")
	os.Setenv("SECURITY_KEY", securityKey)
	os.Setenv("BUCKET", bucket)

	envKeys, err := loading.GetEnvironmentKeys()
	if err != nil {
		t.Errorf("Expected error nil, but got: %v", err)
	}

	if envKeys.SecurityKey != securityKey {
		t.Errorf("Expected securitykey %v, but got: %v", securityKey, envKeys.SecurityKey)
	}

	if envKeys.Bucket != bucket {
		t.Errorf("Expected bucket %v, but got: %v", bucket, envKeys.Bucket)
	}
}
