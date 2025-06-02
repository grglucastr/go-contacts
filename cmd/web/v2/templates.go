package main

import "github.com/grglucastr/go-contacts/internal/models/v2"

type templateData struct {
	Contacts      []models.Contact
	Relationships []models.Relationship
}
