package models

import "database/sql"

type Type struct {
	ID   int
	Name string
}

type TypeModel struct {
	DB *sql.DB
}

func (m *TypeModel) ListAll() ([]Type, error) {
	stmt := "SELECT id, name FROM types"

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var types []Type

	for rows.Next() {
		var t Type

		err := rows.Scan(&t.ID, &t.Name)
		if err != nil {
			return nil, err
		}

		types = append(types, t)
	}

	return types, nil
}
