package db

import "vitrina/models"

func (manager *Manager) GetSpecializations() ([]*models.Specialization, error) {
	var specializations []*models.Specialization

	err := manager.Conn.QueryRow(`SELECT * FROM specializations;`).Scan(&specializations)
	if err != nil {
		return nil, err
	}

	return specializations, nil
}