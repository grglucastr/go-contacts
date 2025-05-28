package main

import (
	"fmt"
	"log"
	"net/http"
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
