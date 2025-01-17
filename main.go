package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing some snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for create a snippet"))
}

func main() {
	mux := http.NewServeMux()

	// All unknown routes will lead to the "/" route if "{$}" is not added
	mux.HandleFunc("/{$}", home) // Restrict this route to exact matches on / only.
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// Start a new server with the address and the handler
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
