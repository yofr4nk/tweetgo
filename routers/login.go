package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/models"
	"github.com/yofr4nk/tweetgo/security"
)

// Login validate user data passed to create a token
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Invalid User/Password "+err.Error(), 400)

		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "The email is required", 400)

		return
	}

	userData, userFound := database.UserLogin(u.Email, u.Password)

	if userFound == false {
		http.Error(w, "Cannot find User "+u.Email, 400)

		return
	}

	tokenKey, err := security.GenerateToken(userData)

	if err != nil {
		http.Error(w, "Something went wrong "+err.Error(), 400)
	}

	response := models.UserLoginResponse{
		Token: tokenKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
