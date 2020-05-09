package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/models"
)

// GetProfile find user data
func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var u models.User

	userEmail := w.Header().Get("Email")

	u, err := database.FindUser(userEmail)

	if err != nil {
		log.Fatal("Cannot get profile error, " + err.Error())
		http.Error(w, "Cannot get profile", http.StatusBadRequest)

		return
	}

	u.Password = ""

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)

}
