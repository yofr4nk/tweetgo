package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetProfile find user data
func GetProfile(getUserFromCtx getUserFromCtx, getUser getUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		usrCtx, err := getUserFromCtx(r.Context())
		if err != nil {
			http.Error(w, "something went wrong getting user from context: "+err.Error(), http.StatusBadRequest)

			return
		}

		u, err := getUser(usrCtx.Email)
		if err != nil {
			log.Print("Cannot get profile error, " + err.Error())
			http.Error(w, "Cannot get profile", http.StatusBadRequest)

			return
		}

		u.Password = ""

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)

	}
}
