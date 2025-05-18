package main

import (
	"log"
	"net/http"
)

func showPageListContacts(w http.ResponseWriter, r *http.Request) {
	log.Println("show page to list contacts")
}

func showPageFormContacts(w http.ResponseWriter, r *http.Request) {
	log.Println("Show page to add or edit contact")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /contacts", showPageListContacts)
	mux.HandleFunc("GET /fcontacts", showPageFormContacts)

	log.Println("Starting server on 4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
