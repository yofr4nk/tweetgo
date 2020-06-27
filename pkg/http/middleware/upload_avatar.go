package middleware

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"tweetgo/pkg/domain"
)

type uploadFile func(filePathName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)

func UploadAvatar(getUserFromCtx getUserFromCtx, updateUser updateUser, uploadFile uploadFile) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var u domain.User
		usrCtx, err := getUserFromCtx(r.Context())
		if err != nil {
			http.Error(w, "something went wrong getting user from context: "+err.Error(), 400)

			return
		}

		userID := usrCtx.ID.Hex()
		email := usrCtx.Email

		file, fileHeader, err := r.FormFile("avatar")
		if err != nil {
			http.Error(w, "Something went wrong getting file ", http.StatusBadRequest)

			return
		}

		defer file.Close()

		filePathName := "avatars/" + primitive.NewObjectID().Hex() + filepath.Ext(fileHeader.Filename)

		fileLocation, err := uploadFile(filePathName, file, fileHeader)
		if err != nil {
			http.Error(w, "Something went wrong uploading avatar ", http.StatusBadRequest)

			return
		}

		u.Avatar = fileLocation

		status, err := updateUser(u, userID)
		if err != nil {
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
}
