package models

import "database/sql"

type Contact struct {
	ID             int32  `form:"id"`
	Name           string `form:"name"`
	PixKey         string `form:"pix_key"`
	RelationshipID int32  `form:"relationship_id"`
}

type ContactModel struct {
	DB *sql.DB
}

func (m *ContactModel) Insert(name string, relationshipID int32) (int64, error) {

	stmt := "INSERT INTO contacts (name, relationship_id) VALUE (?, ?)"

	result, err := m.DB.Exec(stmt, name, relationshipID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *ContactModel) ListAllContacts() ([]Contact, error) {

	stmt := "SELECT id, name, pix_key, relationship_id FROM contacts"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contacts []Contact
	for rows.Next() {

		var contact Contact
		err = rows.Scan(&contact.ID, &contact.Name, &contact.PixKey, &contact.RelationshipID)

		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (m *ContactModel) GetContactById(id int32) (Contact, error) {

	stmt := "SELECT id, name, pix_key, relationship_id FROM contacts WHERE id := ?"

	result := m.DB.QueryRow(stmt, id)

	var contact Contact
	err := result.Scan(&contact.ID, &contact.Name, &contact.PixKey, &contact.RelationshipID)

	if err != nil {
		return Contact{}, err
	}

	return contact, nil

}
