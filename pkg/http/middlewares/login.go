package refmiddlewares

import (
	"encoding/json"
	"net/http"
	"tweetgo/pkg/domain"
)

type comparePassword func(password string, passwordHashed string) error
type generateToken func(u domain.User) (string, error)

type loginPayload struct {
	Token string
}

func Login(getUser getUser, comparePassword comparePassword, generateToken generateToken) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var u domain.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "Invalid User/Password "+err.Error(), 400)

			return
		}

		if len(u.Email) == 0 {
			http.Error(w, "The email is required", 400)

			return
		}

		userData, err := getUser(u.Email)
		if err != nil {
			http.Error(w, "Cannot find User "+err.Error(), 400)

			return
		}

		matchingFail := comparePassword(u.Password, userData.Password)
		if matchingFail != nil {
			http.Error(w, "Error comparing password "+matchingFail.Error(), 400)

			return
		}

		tokenKey, err := generateToken(userData)
		if err != nil {
			http.Error(w, "Something went wrong creating token "+err.Error(), 400)

			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loginPayload{
			Token: tokenKey,
		})
	}
}
