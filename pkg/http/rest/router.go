package rest

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"tweetgo/pkg/deleting"
	"tweetgo/pkg/domain"
	"tweetgo/pkg/finding"
	"tweetgo/pkg/http/middleware"
	"tweetgo/pkg/saving"
	"tweetgo/pkg/tokenizer"
	"tweetgo/pkg/uploading"
	"tweetgo/pkg/validating"
)

// RouterManagement set the main config for routers
func RouterManagement(sus *saving.UserService,
	fus *finding.UserService,
	tks *tokenizer.TokenService,
	sts *saving.TweetService,
	fts *finding.TweetService,
	dts *deleting.TweetService,
	ufs *uploading.UploadFileService,
) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/user-register", middleware.ValidateUserExist(fus, domain.SetUserToContext, middleware.SaveUser(sus, domain.GetUserFromCtx))).Methods("POST")
	router.HandleFunc("/user-login", middleware.Login(fus.GetUser, validating.ComparePassword, tks)).Methods("POST")
	router.HandleFunc("/get-profile", middleware.CheckToken(domain.SetUserToContext, tks, middleware.GetProfile(domain.GetUserFromCtx, fus.GetUser))).Methods("GET")
	router.HandleFunc("/update-profile", middleware.CheckToken(domain.SetUserToContext, tks, middleware.UpdateProfile(domain.GetUserFromCtx, sus.UpdateUser))).Methods("PUT")
	router.HandleFunc("/save-tweet", middleware.CheckToken(domain.SetUserToContext, tks, middleware.SaveTweet(sts.SaveTweet, domain.GetUserFromCtx))).Methods("POST")
	router.HandleFunc("/get-tweet", middleware.CheckToken(domain.SetUserToContext, tks, middleware.GetTweets(fts.GetTweets))).Methods("GET")
	router.HandleFunc("/delete-tweet", middleware.CheckToken(domain.SetUserToContext, tks, middleware.DeleteTweet(dts.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/upload-avatar", middleware.CheckToken(domain.SetUserToContext, tks, middleware.UploadAvatar(domain.GetUserFromCtx, sus.UpdateUser, ufs.UploadFile))).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	return handler
}
