package middlewares

import (
	"net/http"

	"github.com/yofr4nk/tweetgo/security"
)

// CheckToken validate provided token
func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, _, _, err := security.GetTokenData(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Token Error "+err.Error(), http.StatusBadRequest)

			return
		}

		w.Header().Set("Email", claims.Email)
		next.ServeHTTP(w, r)
	}
}
