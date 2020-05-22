package routers

import (
	"net/http"
	"tweetgo/database"
)

// DeleteTweet remove tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	userID := r.URL.Query().Get("userId")
	ID := r.URL.Query().Get("id")

	if len(userID) < 1 {
		http.Error(w, "userId param is required", http.StatusBadRequest)

		return
	}

	if len(ID) < 1 {
		http.Error(w, "id param is required", http.StatusBadRequest)

		return
	}

	err := database.DeleteTweet(ID, userID)

	if err != nil {
		http.Error(w, "Something went wrong removing tweet: "+err.Error(), http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
