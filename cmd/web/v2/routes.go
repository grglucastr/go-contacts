package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	mux := mux.NewRouter()

	//view routes
	mux.HandleFunc("/contacts", showPageListContacts).Methods("GET")
	mux.HandleFunc("/contacts/{id}", showPageContactInfos).Methods("GET")
	mux.HandleFunc("/fcontacts", showPageFormContact).Methods("GET")
	mux.HandleFunc("/fcontacts/{id}", showPageFormContact).Methods("GET")

	return mux
}
