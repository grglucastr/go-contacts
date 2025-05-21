package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/grglucastr/go-contacts/internal/models"
)

func (app *application) postNewContact(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	phone := r.PostForm.Get("phone")

	log.Println("name", name, "email", email, "phone", phone)

	app.contacts.Insert(name, phone, email)

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func showPageIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func showPageListContacts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/contacts.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	contact := models.Contact{ID: 123, Name: "asfsadfas", Email: "sadfsadf", Phone: "21312321"}
	contacts := []models.Contact{}
	contacts = append(contacts, contact)

	data := &templateData{
		Contacts: contacts,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func showPageFormContacts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/fcontacts.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
