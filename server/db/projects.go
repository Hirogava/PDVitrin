package db

import "vitrina/models"

func (manager *Manager) GetProjectsMin(page int, limit int) ([]*models.ProjectMin, error) {
	offset := (page - 1) * limit
    
    query := `SELECT id, name
              FROM projects 
              ORDER BY created_at DESC 
              LIMIT $1 OFFSET $2`

	rows, err := manager.Conn.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.ProjectMin

	for rows.Next() {
		var project models.ProjectMin
		if err := rows.Scan(&project.Id, &project.Name); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}

func (manager *Manager) GetProjectsCount() (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM projects`

	row := manager.Conn.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}