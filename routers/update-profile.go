package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/models"
)

// UpdateProfile modify user data
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Invalid provided data "+err.Error(), 400)

		return
	}

	userID := w.Header().Get("Id")

	status, updateErr := database.UpdateUser(u, userID)

	if updateErr != nil {
		http.Error(w, "Something went wrong updating profile"+err.Error(), 400)

		return
	}

	if status == false {
		http.Error(w, "It could not possible to update profile", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Del("Id")
	w.Header().Del("Email")
	w.WriteHeader(http.StatusCreated)
}
