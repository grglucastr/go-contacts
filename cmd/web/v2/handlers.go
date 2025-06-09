package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	data := app.LoadTemplateData()

	vars := mux.Vars(r)
	cId := vars["id"]
	infId := vars["infoId"]

	if len(cId) > 0 {
		conId := app.ConvertToUnsignedInt(cId)
		contact := app.LoadContact(conId)
		infos := app.LoadInfosByContactId(conId)

		data.AddContact(contact)
		data.AddInfos(infos)

		if infId != "" {
			infoId := app.ConvertToUnsignedInt(infId)
			selectedInfo := app.LoadSingleInfo(infoId)
			data.AddSelectedInfo(selectedInfo)
		}
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

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")
	pixKey := r.Form.Get("pix_key")
	relationship_id := r.Form.Get("relationship_id")

	rel_id, err := strconv.Atoi(relationship_id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = app.ContactModel.Insert(name, pixKey, int32(rel_id))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *application) addContactDetails(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := r.Form.Get("info_id")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	tpId := r.Form.Get("type")

	typeId, err := strconv.Atoi(tpId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	cId := vars["id"]

	contactId, err := strconv.Atoi(cId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if id != "0" {
		infoId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, err = app.InfoModel.UpdateInfo(infoId, email, phone, typeId)
		http.Redirect(w, r, fmt.Sprintf("/fcontacts/%d", contactId), http.StatusSeeOther)
		return
	}

	_, err = app.InfoModel.Insert(email, phone, typeId, contactId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/fcontacts/%d", contactId), http.StatusSeeOther)

}

func (app *application) deleteContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cId := vars["id"]

	if cId == "" {
		log.Fatalln("Invalid contact ID")
		http.Error(w, "Invalid contact ID", http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(cId)
	if err != nil {
		log.Fatalln(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = app.ContactModel.DeleteContactById(int32(id))
	if err != nil {
		log.Fatalln(err.Error())
		http.Error(w, "Error when delete contact", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *application) deleteContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactId := vars["contactId"]
	infoId := vars["infoId"]

	infId, err := strconv.Atoi(infoId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = app.InfoModel.DeleteById(infId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/fcontacts/%s", contactId), http.StatusSeeOther)
}
