package main

import "github.com/grglucastr/go-contacts/internal/models/v2"

type templateData struct {
	Contact       models.Contact
	Info          models.Info
	Contacts      []models.Contact
	Relationships []models.Relationship
	Types         []models.Type
	Infos         []models.Info
}

func (td *templateData) AddSelectedInfo(info models.Info) {
	td.Info = info
}

func (td *templateData) AddInfos(infos []models.Info) {
	td.Infos = infos
}

func (td *templateData) AddContact(c models.Contact) {
	td.Contact = c
}
