package refmiddlewares

import (
	"net/http"
)

type deleteTweet func(ID string, UserID string) error

func DeleteTweet(deleteTweet deleteTweet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		userID := r.URL.Query().Get("userId")
		ID := r.URL.Query().Get("id")

		if len(userID) == 0 {
			http.Error(w, "userId param is required", http.StatusBadRequest)

			return
		}

		if len(ID) == 0 {
			http.Error(w, "id param is required", http.StatusBadRequest)

			return
		}

		err := deleteTweet(ID, userID)
		if err != nil {
			http.Error(w, "Something went wrong removing tweet: "+err.Error(), http.StatusBadGateway)

			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
