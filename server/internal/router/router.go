package router

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"spillane.farm/gowebtemplate/internal/handlers"
)

func Configure(ctx context.Context) (mainRouter *mux.Router, uiRouter *mux.Router) {
	mainRouter = mux.NewRouter()

	//Sub-router for UI to support CSRF auth
	uiRouter = mainRouter.PathPrefix("/api").Subrouter()
	uiRouter.HandleFunc("/hello-world", handlers.HelloWorldHandler(ctx)).Methods(http.MethodGet)
	//Special UI route that can't have csrf protections on it because it sets up the client for further csrf protected calls
	uiRouter.HandleFunc("/csrftoken", handlers.CreateCsrfTokenHandler(ctx)).Methods(http.MethodGet)

	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		mainRouter.ServeHTTP(writer, request)
	})
	return mainRouter, uiRouter
}
