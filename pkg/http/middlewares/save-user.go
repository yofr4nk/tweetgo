package refmiddlewares

import (
	"log"
	"net/http"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/saving"
)

func SaveUser(sus *saving.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(domain.UserCtxKey).(domain.User)

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
			log.Print("Something went wrong saving user " + err.Error())
			http.Error(w, "Something went wrong saving user "+err.Error(), 400)

			return
		}

		if status == false {
			log.Print("Something went wrong saving user ")
			http.Error(w, "Something went wrong saving user ", 400)

			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
