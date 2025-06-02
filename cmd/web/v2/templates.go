package main

import "github.com/grglucastr/go-contacts/internal/models/v2"

type templateData struct {
	Contact       models.Contact
	Contacts      []models.Contact
	Relationships []models.Relationship
}
