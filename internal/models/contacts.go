package models

import (
	"database/sql"
	"errors"
)

type Contact struct {
	ID    int32
	Name  string
	Phone string
	Email string
}

type ContactModel struct {
	DB *sql.DB
}

func (m *ContactModel) Insert(name string, phone string, email string) (int, error) {

	stmt := "INSERT INTO contacts (name, phone, email) VALUES (?, ?, ?)"

	result, err := m.DB.Exec(stmt, name, phone, email)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ContactModel) ListAll() ([]Contact, error) {

	stmt := "SELECT id, name, phone, email FROM contacts"

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email)

		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

func (m *ContactModel) FindById(id int) (Contact, error) {

	stmt := "SELECT id, name, phone, email FROM contacts WHERE id = ?"

	row := m.DB.QueryRow(stmt, id)

	var contact Contact

	err := row.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Contact{}, sql.ErrNoRows
		} else {
			return Contact{}, err
		}
	}

	return contact, nil

}
