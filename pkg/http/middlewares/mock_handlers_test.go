package refmiddlewares_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

func mockServerHTTP(mw http.HandlerFunc, body *strings.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRequest("POST", "/", body)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, w)

	return rr
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
