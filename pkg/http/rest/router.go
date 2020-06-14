package rest

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"tweetgo/middlewares"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/finding"
	rmiddlewares "tweetgo/pkg/http/middlewares"
	"tweetgo/pkg/saving"
	"tweetgo/routers"
)

// RouterManagement set the main config for routers
func RouterManagement(sus *saving.UserService, fus *finding.UserService) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/user-register", rmiddlewares.ValidateUserExist(fus, domain.SetUserToContext, rmiddlewares.SaveUser(sus, domain.GetUserFromCtx))).Methods("POST")
	router.HandleFunc("/user-login", middlewares.CheckDatabase(routers.Login)).Methods("POST")
	router.HandleFunc("/get-profile", middlewares.CheckDatabase(middlewares.CheckToken(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/update-profile", middlewares.CheckDatabase(middlewares.CheckToken(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/save-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/get-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.GetTweet))).Methods("GET")
	router.HandleFunc("/delete-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/upload-avatar", middlewares.CheckDatabase(middlewares.CheckToken(routers.UploadAvatar))).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	return handler
}
