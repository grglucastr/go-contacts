package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/grglucastr/go-contacts/internal/models/v1"
)

type application struct {
	contacts *models.ContactModel
}

func main() {

	db, err := openDB()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer db.Close()

	app := &application{contacts: &models.ContactModel{DB: db}}
	log.Println("App loading", app)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", showPageIndex)
	mux.HandleFunc("GET /contacts", app.showPageListContacts)
	mux.HandleFunc("GET /fcontacts", app.showPageFormContacts)

	// api V1
	mux.HandleFunc("POST /api/v1/contacts", app.postNewContact)
	mux.HandleFunc("GET /api/v1/dcontacts", app.deleteNewContact)

	log.Println("Starting server on 4000")
	err = http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}

func openDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "web_go_contacts:pass123@/go_contacts")

	if err != nil {
		return nil, err
	}

	return db, nil
}

// dude nothing special here
