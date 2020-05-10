package routers

import (
	"net/http"
	"path/filepath"

	"github.com/yofr4nk/tweetgo/multimedia"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadAvatar allow to upload avatar image
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("avatar")

	if err != nil {
		http.Error(w, "Something went wrong getting file ", http.StatusBadRequest)
	}

	defer file.Close()

	filePathName := "avatars/" + primitive.NewObjectID().Hex() + filepath.Ext(fileHeader.Filename)

	uErr := multimedia.UploadFile(filePathName, file, fileHeader)

	if uErr != nil {
		http.Error(w, "Something went wrong uploading avatar ", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}
