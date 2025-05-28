package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type application struct {
}

func main() {

	app := &application{}

	fmt.Println("Loading server on port :4000")
	err := http.ListenAndServe(":4000", app.routes())

	if err != nil {
		log.Panic(err.Error())
	}
}

func showPageListContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show page to list contacts")
}

func showPageContactInfos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show page of contact details")
}

func showPageFormContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cId := vars["id"]

	if len(cId) > 0 {
		fmt.Fprintln(w, "Edit contact form")
		return
	}

	fmt.Fprintln(w, "New contact form")
}
