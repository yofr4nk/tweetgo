package middlewares

import (
	"log"
	"net/http"

	"github.com/yofr4nk/tweetgo/database"
)

//CheckDatabase validate if database is connected before continue
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckingConnection() == false {
			log.Fatal("Missing database connection")
			http.Error(w, "Sorry, something went wrong", 500)

			return
		}

		next.ServeHTTP(w, r)
	}
}
