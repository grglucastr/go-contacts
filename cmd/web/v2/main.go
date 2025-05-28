package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	v2 "github.com/grglucastr/go-contacts/internal/models/v2"
)

type application struct {
	ContactModel *v2.ContactModel
}

func main() {

	db, err := openDB()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		ContactModel: &v2.ContactModel{DB: db},
	}

	fmt.Println("Loading server on port :4000")
	err = http.ListenAndServe(":4000", app.routes())

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func openDB() (*sql.DB, error) {
	conn, err := sql.Open("mysql", "gowebv2:webv2@/go_contacts_v2")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
