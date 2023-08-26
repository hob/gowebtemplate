package handlers

import (
	"context"
	"net/http"
)

func HelloWorldHandler(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		w.WriteHeader(http.StatusOK)
	}
}