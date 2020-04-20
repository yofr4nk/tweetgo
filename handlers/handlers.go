package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/yofr4nk/tweetgo/middlewares"
	"github.com/yofr4nk/tweetgo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// MainManagement set the main config for routers
func MainManagement() {
	router := mux.NewRouter()

	router.HandleFunc("/user-register", middlewares.CheckDatabase(routers.UserRegister)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
