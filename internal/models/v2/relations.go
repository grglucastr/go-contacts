package models

import "database/sql"

type Relationship struct {
	ID   int
	Name string
}

type RelationshipModel struct {
	DB *sql.DB
}

func (m *RelationshipModel) ListAll() ([]Relationship, error) {

	stmt := "SELECT id, name FROM relationships"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var relationships []Relationship
	defer rows.Close()

	for rows.Next() {

		var rel Relationship
		err := rows.Scan(&rel.ID, &rel.Name)
		if err != nil {
			return nil, err
		}

		relationships = append(relationships, rel)
	}

	return relationships, nil
}
