package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// HTML routes
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreateForm)

	// CRUD routes
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
