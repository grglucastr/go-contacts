package models

import "database/sql"

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