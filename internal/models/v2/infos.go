package models

import "database/sql"

type Info struct {
	ID              int
	Email           string
	Phone           string
	TypeID          int
	TypeDescription string
	ContactID       int
}

type InfoModel struct {
	DB *sql.DB
}

func (m *InfoModel) Insert(email string, phone string, typeId int, contactId int) (int, error) {

	stmt := "INSERT INTO infos (email, phone, type_id, contact_id) VALUES (?, ?, ?, ?)"

	result, err := m.DB.Exec(stmt, email, phone, typeId, contactId)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *InfoModel) ListAllByContactsId(contactId int) ([]Info, error) {

	stmt := `SELECT i.id, i.email, i.phone, i.type_id, t.name, i.contact_id
			FROM infos i
			INNER JOIN types t ON t.id = i.type_id
			WHERE i.contact_id = ?
			ORDER BY id ASC`

	rows, err := m.DB.Query(stmt, contactId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var infos []Info

	for rows.Next() {

		var info Info

		err := rows.Scan(&info.ID, &info.Email, &info.Phone, &info.TypeID, &info.TypeDescription, &info.ContactID)
		if err != nil {
			return nil, err
		}

		infos = append(infos, info)
	}

	return infos, nil
}
