package db

import "vitrina/models"

func (manager *Manager) GetSpecializations() ([]*models.Specialization, error) {
	var specializations []*models.Specialization

	rows, err := manager.Conn.Query("SELECT id, name FROM specializations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var specialization models.Specialization

		if err := rows.Scan(&specialization.Id, &specialization.Name); err != nil {
            return nil, err
        }

        specializations = append(specializations, &specialization)
	}

	return specializations, nil
}