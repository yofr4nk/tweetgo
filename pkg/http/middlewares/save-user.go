package refmiddlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"tweetgo/pkg/saving"
)

func SaveUser(sus *saving.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u saving.User

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

		_, status, err := sus.SaveUser(u)
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
}
