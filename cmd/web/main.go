package main

import (
	"database/sql"
	"time"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

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


func openDB(){

	db, err := sql.Open("mysql", "web_go_contacts@/go_contacts")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	

}