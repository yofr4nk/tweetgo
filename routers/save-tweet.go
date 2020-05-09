package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/yofr4nk/tweetgo/database"
	"github.com/yofr4nk/tweetgo/models"
)

// SaveTweet allow to save a new tweet
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var t models.Tweet

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "the received data has errors "+err.Error(), 400)

		return
	}

	userID := w.Header().Get("Id")

	tweetMessage := models.Tweet{
		UserID:  userID,
		Message: t.Message,
		Date:    time.Now().Unix(),
	}

	_, status, err := database.SaveTweet(tweetMessage)

	if err != nil {
		log.Fatal("Something went wront saving tweet " + err.Error())
		http.Error(w, "Something went wront saving tweet "+err.Error(), 400)

		return
	}

	if status == false {
		log.Fatal("Something went wront saving tweet ")
		http.Error(w, "Something went wront saving tweet ", 400)

		return
	}

	w.Header().Del("Id")
	w.WriteHeader(http.StatusCreated)
}
