package refmiddlewares

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tweetgo/pkg/domain"
)

type getTweets func(userID string, page int64) (domain.Tweets, error)

// GetTweets find tweets data from userId
func GetTweets(getTweets getTweets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		userID := r.URL.Query().Get("userId")
		pageString := r.URL.Query().Get("page")

		if len(userID) == 0 {
			http.Error(w, "userId param is required", http.StatusBadRequest)

			return
		}

		if len(pageString) < 1 {
			http.Error(w, "page param is required", http.StatusBadRequest)

			return
		}

		page, err := strconv.ParseInt(pageString, 10, 64)
		if err != nil {
			http.Error(w, "Something went wrong getting page value "+err.Error(), http.StatusBadRequest)

			return
		}

		if page < 1 {
			http.Error(w, "page param must be an number greater than 0", http.StatusBadRequest)

			return
		}

		response, err := getTweets(userID, page)
		if err != nil {
			http.Error(w, "Something went wrong getting tweets "+err.Error(), http.StatusBadRequest)

			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
