package handlers

import (
	"log"
	"net/http"
	"os"
	"tweetgo/middlewares"
	"tweetgo/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// MainManagement set the main config for routers
func MainManagement() {
	router := mux.NewRouter()

	router.HandleFunc("/user-register", middlewares.CheckDatabase(routers.UserRegister)).Methods("POST")
	router.HandleFunc("/user-login", middlewares.CheckDatabase(routers.Login)).Methods("POST")
	router.HandleFunc("/get-profile", middlewares.CheckDatabase(middlewares.CheckToken(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/update-profile", middlewares.CheckDatabase(middlewares.CheckToken(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/save-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/get-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.GetTweet))).Methods("GET")
	router.HandleFunc("/delete-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/upload-avatar", middlewares.CheckDatabase(middlewares.CheckToken(routers.UploadAvatar))).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
