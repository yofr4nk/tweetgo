package refmiddlewares_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"tweetgo/pkg/domain"
)

type userCtxMock struct {
	shouldFail bool
	usr        domain.User
}

func mockServerHTTP(mw http.HandlerFunc, body *strings.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRequest("POST", "/", body)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, w)

	return rr
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getUserFromCtxMock(usm userCtxMock) func(ctx context.Context) (domain.User, error) {
	return func(ctx context.Context) (domain.User, error) {
		if usm.shouldFail {
			return domain.User{}, errors.New("error getting user")
		}

		return usm.usr, nil
	}
}
