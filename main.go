package main

import (
	"log"
	"net/http"
)

// Home handler that writes a hello message
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Initialize a servemux to route handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Start new web server on port 4000
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
