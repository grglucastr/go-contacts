package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	mux := mux.NewRouter()

	//view routes
	mux.HandleFunc("/", app.index).Methods("GET")
	mux.HandleFunc("/contacts", app.showPageListContacts).Methods("GET")
	mux.HandleFunc("/contacts/{id}", app.showPageContactInfos).Methods("GET")
	mux.HandleFunc("/fcontacts", app.showPageFormContact).Methods("GET")
	mux.HandleFunc("/fcontacts/{id}", app.showPageFormContact).Methods("GET")
	mux.HandleFunc("/fcontacts/{id}/infos/{infoId}", app.showPageFormContact).Methods("GET")
	mux.HandleFunc("/dcontacts/{contactId}/infos/{infoId}", app.deleteContactInfo).Methods("GET")

	// rest routes
	mux.HandleFunc("/api/v2/contacts", app.listContacts).Methods("GET")
	mux.HandleFunc("/api/v2/contacts/{id}", app.getContactById).Methods("GET")
	mux.HandleFunc("/api/v2/contacts/{id}", app.getContactById).Methods("DELETE")
	mux.HandleFunc("/api/v2/contacts", app.addContact).Methods("POST")
	mux.HandleFunc("/api/v2/contacts/{id}/details", app.addContactDetails).Methods("POST")

	return mux
}
