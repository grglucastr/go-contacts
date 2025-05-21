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
