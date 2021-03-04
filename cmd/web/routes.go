package main

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/lcsval/gohtml/pkg/config"
	"github.com/lcsval/gohtml/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))

	return mux
}
