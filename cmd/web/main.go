package main

import (
	"fmt"
	"net/http"

	"github.com/tanishqv/mywebapp-go/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting appilcation on %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
