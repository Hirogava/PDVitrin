package db

import "vitrina/models"

func (manager *Manager) GetThematic() ([]*models.Thematic, error) {
	var themes []*models.Thematic

	rows, err := manager.Conn.Query("SELECT id, name FROM thematic")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var theme models.Thematic

		if err := rows.Scan(&theme.Id, &theme.Name); err != nil {
            return nil, err
        }

        themes = append(themes, &theme)
	}

	return themes, nil
}