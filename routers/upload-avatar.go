package routers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"tweetgo/database"
	"tweetgo/models"
	"tweetgo/multimedia"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadAvatar allow to upload avatar image
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	userID := w.Header().Get("Id")
	email := w.Header().Get("Email")

	var u models.User

	file, fileHeader, err := r.FormFile("avatar")

	if err != nil {
		http.Error(w, "Something went wrong getting file ", http.StatusBadRequest)
	}

	defer file.Close()

	filePathName := "avatars/" + primitive.NewObjectID().Hex() + filepath.Ext(fileHeader.Filename)

	fileLocation, uErr := multimedia.UploadFile(filePathName, file, fileHeader)

	if uErr != nil {
		http.Error(w, "Something went wrong uploading avatar ", http.StatusBadRequest)

		return
	}

	u.Avatar = fileLocation

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
