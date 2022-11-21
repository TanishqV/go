package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting appilcation on %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
