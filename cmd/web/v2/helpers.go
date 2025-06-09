package main

import (
	"log"
	"strconv"

	"github.com/grglucastr/go-contacts/internal/models/v2"
)

func (app *application) LoadRelationships() []models.Relationship {
	relationships, err := app.RelationshipModel.ListAll()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return relationships
}

func (app *application) LoadTypes() []models.Type {
	types, err := app.TypeModel.ListAll()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return types
}

func (app *application) LoadTemplateData() *templateData {
	return &templateData{
		Relationships: app.LoadRelationships(),
		Types:         app.LoadTypes(),
	}
}

func (app *application) LoadInfosByContactId(cId int) []models.Info {
	infos, err := app.InfoModel.ListAllByContactsId(cId)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return infos
}

func (app *application) LoadSingleInfo(infoId int) models.Info {
	info, err := app.InfoModel.GetById(infoId)
	if err != nil {
		return models.Info{}
	}
	return info
}

func (app *application) LoadContact(cId int) models.Contact {
	contact, err := app.ContactModel.GetContactById(int32(cId))

	if err != nil {
		log.Println(err.Error())
		return models.Contact{}
	}
	return contact
}

func (app *application) ConvertToUnsignedInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return i
}
