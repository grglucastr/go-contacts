package main

import "github.com/grglucastr/go-contacts/internal/models/v1"

type templateData struct {
	Contacts []models.Contact
	Contact  models.Contact
}
