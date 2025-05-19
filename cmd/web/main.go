package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", showPageIndex)
	mux.HandleFunc("GET /contacts", showPageListContacts)
	mux.HandleFunc("GET /fcontacts", showPageFormContacts)

	// api V1
	mux.HandleFunc("POST /api/v1/contacts", postNewContact)

	log.Println("Starting server on 4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
