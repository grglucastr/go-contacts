package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *application) showPageListContacts(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/v2/base.html",
		"./ui/html/v2/pages/contacts.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	contacts, err := app.ContactModel.ListAllContacts()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	templateData := &templateData{
		Contacts: contacts,
	}

	err = ts.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) showPageContactInfos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show page of contact details")
}

func (app *application) showPageFormContact(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/v2/pages/fcontacts.html",
		"./ui/html/v2/base.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	cId := vars["id"]

	relationships, err := app.RelationshipModel.ListAll()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &templateData{
		Relationships: relationships,
	}

	if len(cId) > 0 {
		fmt.Fprintln(w, "Edit contact form", data)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) listContacts(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getContactById(w http.ResponseWriter, r *http.Request) {

}

func (app *application) addContact(w http.ResponseWriter, r *http.Request) {

}
