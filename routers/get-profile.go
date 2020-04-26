package routers

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	u.Password = ""

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(u)

}
