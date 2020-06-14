package refmiddlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/finding"
)

func ValidateUserExist(fus *finding.UserService, next http.HandlerFunc) http.HandlerFunc {
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

		userExist, findErr := fus.FindUserExists(u.Email)

		if userExist == true {
			http.Error(w, "The user already exist", 400)

			return
		}

		if findErr != nil {
			http.Error(w, "Something went wrong searching user: "+findErr.Error(), 400)

			return
		}

		ctx := context.WithValue(r.Context(), domain.UserCtxKey, u)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
