package handlers

import (
	"context"
	"github.com/gorilla/csrf"
	"net/http"
)

func CreateCsrfTokenHandler(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		w.WriteHeader(http.StatusOK)
	}
}
