package middleware

import (
	"encoding/json"
	"net/http"
	"tweetgo/pkg/domain"
)

func ValidateUserExist(fus userFinder, setUserToCtx setUserToCtx, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u domain.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "the received data has errors "+err.Error(), 400)

			return
		}

		if len(u.Email) == 0 {
			http.Error(w, "The email is required", 400)

			return
		}

		userExist, err := fus.FindUserExists(u.Email)

		if userExist {
			http.Error(w, "The user already exist", 400)

			return
		}

		if err != nil {
			http.Error(w, "Something went wrong searching user: "+err.Error(), 400)

			return
		}

		ctx := setUserToCtx(r.Context(), u)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
