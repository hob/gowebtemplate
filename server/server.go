package main

import (
	"context"
	"github.com/gorilla/csrf"
	"log"
	"net/http"
	"os"
	"os/signal"
	"spillane.farm/gowebtemplate/internal/cache"
	"spillane.farm/gowebtemplate/internal/db"
	"spillane.farm/gowebtemplate/internal/router"
	"spillane.farm/gowebtemplate/internal/system"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// server entry - reads config, sets up environment, sets up router, and starts server on a default port
func main() {
	config := system.GetConfig()
	logrus.SetOutput(logrus.StandardLogger().Out)
	mainRouter, uiRouter := router.Configure(context.Background())
	db.Init(config)
	db.Migrate()
	defer db.Close()
	cache.Init(config)

	// finally, start server
	srv := &http.Server{
		Handler:      mainRouter,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		csrfMiddleware := csrf.Protect([]byte(config.CsrfTokenKey), csrf.Path("/api"))
		uiRouter.Use(csrfMiddleware)
		if err := http.ListenAndServe(":8080", mainRouter); err != nil {
			logrus.WithError(err).Fatal("failed to start http server")
			panic(err)
		}
	}()
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
