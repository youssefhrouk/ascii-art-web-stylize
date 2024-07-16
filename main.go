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
	stylz := http.FileServer(http.Dir("stylize"))
	http.Handle("/css/", http.StripPrefix("/css", stylz))

	// Start the server

	fmt.Println("Server starting on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
