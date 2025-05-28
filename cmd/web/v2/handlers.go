package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) showPageListContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show page to list contacts")
}

func (app *application) showPageContactInfos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show page of contact details")
}

func (app *application) showPageFormContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cId := vars["id"]

	if len(cId) > 0 {
		fmt.Fprintln(w, "Edit contact form")
		return
	}

	fmt.Fprintln(w, "New contact form")
}


func (app *application) listContacts(w http.ResponseWriter, r *http.Request){

}

func (app *application) getContactById(w http.ResponseWriter, r *http.Request){

}

func (app *application) addContact(w http.ResponseWriter, r *http.Request){

}