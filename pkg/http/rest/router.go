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
	"tweetgo/pkg/tokenizer"
	"tweetgo/pkg/validating"
	"tweetgo/routers"
)

// RouterManagement set the main config for routers
func RouterManagement(sus *saving.UserService, fus *finding.UserService, tks tokenizer.TokenService) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/user-register", rmiddlewares.ValidateUserExist(fus, domain.SetUserToContext, rmiddlewares.SaveUser(sus, domain.GetUserFromCtx))).Methods("POST")
	router.HandleFunc("/user-login", rmiddlewares.Login(fus.GetUser, validating.ComparePassword, tks)).Methods("POST")
	router.HandleFunc("/get-profile", rmiddlewares.CheckToken(domain.SetUserToContext, tks, rmiddlewares.GetProfile(domain.GetUserFromCtx, fus.GetUser))).Methods("GET")
	router.HandleFunc("/update-profile", rmiddlewares.CheckToken(domain.SetUserToContext, tks, rmiddlewares.UpdateProfile(domain.GetUserFromCtx, sus.UpdateUser))).Methods("PUT")
	router.HandleFunc("/save-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/get-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.GetTweet))).Methods("GET")
	router.HandleFunc("/delete-tweet", middlewares.CheckDatabase(middlewares.CheckToken(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/upload-avatar", middlewares.CheckDatabase(middlewares.CheckToken(routers.UploadAvatar))).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	return handler
}
