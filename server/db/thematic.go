package db

import "vitrina/models"

func (manager *Manager) GetThematic() ([]*models.Thematic, error) {
	var themes []*models.Thematic

	err := manager.Conn.QueryRow("SELECT * FROM thematic").Scan(&themes)
	if err != nil {
		return nil, err
	}

	return themes, nil
}