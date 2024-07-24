package main

import (
	"fmt"
	"log"
	"net/http"

	ascii "ascii/Handlers"
)

func main() {
	http.HandleFunc("/", ascii.IndexHandler)
	http.HandleFunc("/ascii-art", ascii.AsciiHandler)
	// This creates a file server to serve static files from the stylize directory.
	stylize := http.FileServer(http.Dir("stylize"))
	// this serve static files and it strips the /css prefix and serves the static file from that directory
	http.Handle("/css/", http.StripPrefix("/css", stylize))

	// Start the server

	fmt.Println("Server starting on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
