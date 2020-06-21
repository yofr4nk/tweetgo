package refmiddlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"tweetgo/models"
	"tweetgo/pkg/domain"
)

type saveTweet func(t domain.Tweet) (bool, error)

func SaveTweet(saveTweet saveTweet, getUserFromCtx getUserFromCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t models.Tweet

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			http.Error(w, "the received data has errors "+err.Error(), 400)

			return
		}

		u, err := getUserFromCtx(r.Context())
		if err != nil {
			http.Error(w, "something went wrong getting user from context: "+err.Error(), http.StatusBadRequest)

			return
		}

		tweetMessage := domain.Tweet{
			UserID:  u.ID.Hex(),
			Message: t.Message,
			Date:    time.Now().Unix(),
		}

		status, err := saveTweet(tweetMessage)
		if err != nil {
			log.Print("Something went wrong saving tweet " + err.Error())
			http.Error(w, "Something went wrong saving tweet "+err.Error(), http.StatusBadRequest)

			return
		}

		if status == false {
			log.Print("Something went wrong saving tweet ")
			http.Error(w, "Something went wrong saving tweet ", http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
