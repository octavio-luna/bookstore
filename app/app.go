package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/octavio-luna/bookstore/datasources"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()
	datasources.ConnectDB()
	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
