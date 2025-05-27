package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/grglucastr/go-contacts/internal/models/v1"
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
	pId := r.PostForm.Get("id")

	log.Println("name", name, "email", email, "phone", phone)

	if len(pId) == 0 {
		app.contacts.Insert(name, phone, email)
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	log.Println("id ", pId)
	id, err := strconv.Atoi(pId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	app.contacts.Update(name, phone, email, id)
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func showPageIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *application) showPageListContacts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/v1/base.html",
		"./ui/html/v1/pages/contacts.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	contacts, err := app.contacts.ListAll()
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

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

func (app *application) showPageFormContacts(w http.ResponseWriter, r *http.Request) {

	pId := r.URL.Query().Get("id")

	var contact models.Contact
	if len(pId) > 0 {
		id, err := strconv.Atoi(pId)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		contact, err = app.contacts.FindById(id)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	} else {
		contact = models.Contact{}
	}

	data := &templateData{
		Contact: contact,
	}

	files := []string{
		"./ui/html/v1/base.html",
		"./ui/html/v1/pages/fcontacts.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) deleteNewContact(w http.ResponseWriter, r *http.Request) {

	pId := r.URL.Query().Get("id")

	id, err := strconv.Atoi(pId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = app.contacts.DeleteById(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
