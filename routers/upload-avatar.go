package routers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/models"
	"github.com/yofr4nk/tweetgo/multimedia"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadAvatar allow to upload avatar image
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	userID := w.Header().Get("Id")
	email := w.Header().Get("Email")
	bucketURL := os.Getenv("AWS_BASE_URL")
	var u models.User

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

	u.Avatar = bucketURL + filePathName

	status, updateErr := database.UpdateUser(u, userID)

	if updateErr != nil {
		http.Error(w, "Something went wrong uploading avatar"+err.Error(), 400)

		return
	}

	if status == false {
		http.Error(w, "It could not possible to upload avatar", http.StatusBadRequest)

		return
	}

	objID, _ := primitive.ObjectIDFromHex(userID)

	u.Email = email
	u.ID = objID

	w.Header().Set("Content-type", "application/json")
	w.Header().Del("Id")
	w.Header().Del("Email")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
