package middlewares

import (
	"net/http"
	"tweetgo/security"
)

// CheckToken validate provided token
func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		claims, _, _, err := security.GetTokenData(authorization)

		if err != nil {
			http.Error(w, "Token Error "+err.Error(), http.StatusBadRequest)

			return
		}

		w.Header().Set("Email", claims.Email)
		w.Header().Set("Id", claims.ID.Hex())
		next.ServeHTTP(w, r)
	}
}
