package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"tweetgo/database"
	"tweetgo/models"
)

// UserRegister get the data from request and save it
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "the received data has errors "+err.Error(), 400)

		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "The email is required", 400)

		return
	}

	if len(u.Password) < 6 {
		http.Error(w, "The password is required", 400)

		return
	}

	userExist, findErr := database.FindUserExists(u.Email)

	if userExist == true {
		http.Error(w, "The user already exist", 400)

		return
	}

	if findErr != nil {
		http.Error(w, "Something went wront searching user"+findErr.Error(), 400)

		return
	}

	_, status, err := database.SaveUser(u)

	if err != nil {
		log.Print("Something went wront saving user " + err.Error())
		http.Error(w, "Something went wront saving user "+err.Error(), 400)

		return
	}

	if status == false {
		log.Print("Something went wront saving user ")
		http.Error(w, "Something went wront saving user ", 400)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
