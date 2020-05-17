package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yofr4nk/tweetgo/database"
)

// GetTweet find tweets data
func GetTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	userID := r.URL.Query().Get("userId")
	pageString := r.URL.Query().Get("page")

	if len(userID) < 1 {
		http.Error(w, "userId param is required", http.StatusBadRequest)

		return
	}

	if len(pageString) < 1 {
		http.Error(w, "page param is required", http.StatusBadRequest)

		return
	}

	page, err := strconv.ParseInt(pageString, 10, 64)

	if page < 1 {
		http.Error(w, "page param must be an number greater than 0", http.StatusBadRequest)

		return
	}

	if err != nil {
		http.Error(w, "Something went wrong getting page value "+err.Error(), http.StatusBadRequest)

		return
	}

	response, status := database.GetTweets(userID, page)

	if status == false {
		http.Error(w, "Something went wrong getting tweets ", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
