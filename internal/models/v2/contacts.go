package models

import "database/sql"

type Contact struct {
	ID               int32  `form:"id"`
	Name             string `form:"name"`
	PixKey           string `form:"pix_key"`
	RelationshipID   int32  `form:"relationship_id"`
	RelationshipName string
}

type ContactModel struct {
	DB *sql.DB
}

func (m *ContactModel) Insert(name string, pixKey string, relationshipID int32) (int64, error) {

	stmt := "INSERT INTO contacts (name, pix_key, relationship_id) VALUE (?, ?, ?)"

	result, err := m.DB.Exec(stmt, name, pixKey, relationshipID)
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

	stmt := `
		SELECT c.id, c.name, c.pix_key, c.relationship_id, r.name as relationship_name FROM contacts c
		INNER JOIN relationships r ON r.id = c.relationship_id
		ORDER BY c.id ASC
	`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contacts []Contact
	for rows.Next() {

		var contact Contact
		err = rows.Scan(&contact.ID, &contact.Name, &contact.PixKey, &contact.RelationshipID, &contact.RelationshipName)

		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (m *ContactModel) GetContactById(id int32) (Contact, error) {

	stmt := "SELECT id, name, pix_key, relationship_id FROM contacts WHERE id = ?"

	result := m.DB.QueryRow(stmt, id)

	var contact Contact
	err := result.Scan(&contact.ID, &contact.Name, &contact.PixKey, &contact.RelationshipID)

	if err != nil {
		return Contact{}, err
	}

	return contact, nil

}

func (m *ContactModel) DeleteContactById(id int32) error {

	stmt := "DELETE FROM contacts WHERE id = ?"

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *ContactModel) UpdateContact(id int, name string, pixKey string, relationshipId int) (int, error) {
	stmt := "UPDATE contacts SET name = ?, pix_key =?, relationship_id = ? WHERE id = ?"

	result, err := m.DB.Exec(stmt, name, pixKey, relationshipId, id)
	if err != nil {
		return 0, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(i), nil
}
