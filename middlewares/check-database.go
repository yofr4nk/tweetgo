package middlewares

import (
	"log"
	"net/http"
	"tweetgo/database"
)

//CheckDatabase validate if database is connected before continue
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckingConnection() == false {
			log.Print("Missing database connection")
			http.Error(w, "Sorry, something went wrong", 500)

			return
		}

		next.ServeHTTP(w, r)
	}
}
