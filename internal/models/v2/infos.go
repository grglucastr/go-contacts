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

func (m *InfoModel) DeleteById(infoId int) (bool, error) {
	stmt := "DELETE FROM infos WHERE id  = ?"

	_, err := m.DB.Exec(stmt, infoId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *InfoModel) GetById(infoId int) (Info, error) {
	stmt := `SELECT i.id, i.email, i.phone, i.type_id FROM infos i WHERE i.id = ?`

	row := m.DB.QueryRow(stmt, infoId)

	var info Info
	err := row.Scan(&info.ID, &info.Email, &info.Phone, &info.TypeID)

	if err != nil {
		return Info{}, err
	}

	return info, nil
}

func (m *InfoModel) UpdateInfo(infoId int, email string, phone string, typeId int) (int, error) {
	stmt := "UPDATE infos SET email = ?, phone = ?, type_id = ? WHERE id = ?"

	res, err := m.DB.Exec(stmt, email, phone, typeId, infoId)
	if err != nil {
		return 0, err
	}

	i, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(i), nil
}