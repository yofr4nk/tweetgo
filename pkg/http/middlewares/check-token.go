package refmiddlewares

import (
	"net/http"
	"tweetgo/pkg/domain"
)

type tokenizerService interface {
	GetAndValidateTokenData(token string) (domain.User, bool, error)
}

// CheckToken validate provided token
func CheckToken(setUserToCtx setUserToCtx,
	tokenizerService tokenizerService,
	next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		u, isValid, err := tokenizerService.GetAndValidateTokenData(authorization)
		if err != nil {
			http.Error(w, "Token Error "+err.Error(), http.StatusBadRequest)

			return
		}

		if isValid == false {
			http.Error(w, "Invalid token", http.StatusBadRequest)

			return
		}

		ctx := setUserToCtx(r.Context(), u)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
