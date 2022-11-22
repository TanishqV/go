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

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting appilcation on %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
