package routers

import (
	"encoding/json"
	"net/http"
	"github.com/yofr4nk/tweetgo/models"
)

// UserRegister get the data from request and save it
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var u models.User
	
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "the data received has errors "+err.Error(), 400)

		return
	}

	if len(u.Name) == 0 {
		http.Error(w, "The name is required", 400)
	}

	if len(u.Email) == 0 {
		http.Error(w, "The email is required", 400)
	}

	w.WriteHeader(http.StatusCreated)
}


