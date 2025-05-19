package main

import (
	"html/template"
	"log"
	"net/http"
)

type Contact struct {
	ID    int32
	Name  string
	Phone string
	Email string
}

type templateData struct {
	Contacts []Contact
}

func postNewContact(w http.ResponseWriter, r *http.Request) {

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

	contact := Contact{ID: 123, Name: "asfsadfas", Email: "sadfsadf", Phone: "21312321"}
	contacts := []Contact{}
	contacts = append(contacts, contact)

	data := &templateData{
		Contacts: contacts,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
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
