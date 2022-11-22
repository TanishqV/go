package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tanishqv/mywebapp-go/pkg/config"
	"github.com/tanishqv/mywebapp-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Creating HTTP handler, often called a "mux" or "multiplexer"
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
