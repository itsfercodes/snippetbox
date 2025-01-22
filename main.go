package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for create a snippet"))
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Saving new snippet..."))
}

func main() {
	// A ServeMux is a router
	mux := http.NewServeMux()

	// All unknown routes will lead to the "/" route if "{$}" is not added
	mux.HandleFunc("GET /{$}", home) // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Print("starting server on :4000")

	// Start a new server with the address and the handler
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
