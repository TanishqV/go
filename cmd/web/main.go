package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tanishqv/mywebapp-go/pkg/config"
	"github.com/tanishqv/mywebapp-go/pkg/handlers"
	"github.com/tanishqv/mywebapp-go/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
