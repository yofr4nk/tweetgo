package refmiddlewares

import (
	"log"
	"net/http"
)

func SaveUser(us userSaver, getUserFromCtx getUserFromCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := getUserFromCtx(r.Context())
		if err != nil {
			http.Error(w, "something went wrong getting user from context: "+err.Error(), 400)

			return
		}

		if len(u.Email) == 0 {
			http.Error(w, "the email is required", 400)

			return
		}

		if len(u.Password) < 6 {
			http.Error(w, "the password is required", 400)

			return
		}

		status, err := us.SaveUser(u)
		if err != nil {
			log.Print("something went wrong saving user " + err.Error())
			http.Error(w, "something went wrong saving user "+err.Error(), 400)

			return
		}

		if status == false {
			log.Print("cannot save user")
			http.Error(w, "cannot save user", 400)

			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
