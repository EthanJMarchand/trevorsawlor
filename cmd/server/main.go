package main

import (
	"fmt"
	"github.com/ethanjmarchand/trevorsawlor/internal/handlers"
	"log"
	"net/http"
)

const portNumber = ":3000"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
