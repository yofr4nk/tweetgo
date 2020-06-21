package middleware

import (
	"encoding/json"
	"net/http"
	"tweetgo/pkg/domain"
)

func UpdateProfile(getUserFromCtx getUserFromCtx, updateUser updateUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		var u domain.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "Invalid provided data "+err.Error(), http.StatusBadRequest)

			return
		}

		usrCtx, err := getUserFromCtx(r.Context())
		if err != nil {
			http.Error(w, "something went wrong getting user from context: "+err.Error(), http.StatusBadRequest)

			return
		}

		status, err := updateUser(u, usrCtx.ID.Hex())
		if err != nil {
			http.Error(w, "Something went wrong updating profile: "+err.Error(), http.StatusBadRequest)

			return
		}

		if status == false {
			http.Error(w, "It could not possible to update profile", http.StatusBadGateway)

			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
