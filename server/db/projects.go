package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"vitrina/models"

	"github.com/lib/pq"
)

// CREATE TABLE projects (
//     id serial PRIMARY KEY,
//     name VARCHAR(255) NOT NULL,
//     purpose text NOT NULL,
//     relevance text NOT NULL,
//     result text NOT NULL,
//     created_at TIMESTAMP DEFAULT NOW()
// );

// CREATE TABLE thematic (
//     id serial PRIMARY KEY,
//     name VARCHAR(255) NOT NULL UNIQUE,
//     description text
// );

// CREATE TABLE project_thematics (
//     project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
//     thematic_id INTEGER REFERENCES thematic(id) ON DELETE CASCADE,
//     PRIMARY KEY (project_id, thematic_id)
// );

// CREATE TABLE specializations (
//     id serial PRIMARY KEY,
//     name VARCHAR(255) NOT NULL UNIQUE,
//     description text
// );

// CREATE TABLE project_specializations (
//     project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
//     specialization_id INTEGER REFERENCES specializations(id) ON DELETE CASCADE,
//     PRIMARY KEY (project_id, specialization_id)
// );

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

func (manager *Manager) GetProject(id int) (*models.Project, error) {
	var project models.Project

	query := `
		SELECT 
			p.id, p.name, p.purpose, p.relevance, p.result,
			(SELECT json_agg(json_build_object('id', t.id, 'name', t.name)) FROM thematic t JOIN project_thematics pt ON t.id = pt.thematic_id WHERE pt.project_id = p.id) AS thematic,
			(SELECT json_agg(json_build_object('id', s.id, 'name', s.name)) FROM specializations s JOIN project_specializations ps ON s.id = ps.specialization_id WHERE ps.project_id = p.id) AS specializations
			FROM projects p
			WHERE p.id = $1
			LIMIT 1
	`

	rows, err := manager.Conn.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var thematics []byte
		var specializations []byte
		if err := rows.Scan(&project.Id, &project.Name, &project.Purpose, &project.Relevance, &project.Result, &thematics, &specializations); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(thematics, &project.Thematics); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(specializations, &project.Specializations); err != nil {
			return nil, err
		}
	}

	return &project, nil
}

func (manager *Manager) GetProjectsByFilters(filters *models.ProjectFilter) ([]*models.ProjectMin, error) {
	query := `
		SELECT DISTINCT p.id, p.name, p.created_at
		FROM projects p
	`
	var whereClauses []string
	var args []interface{}
	argIdx := 1

	if filters.Specializations != nil && len(filters.Specializations.Id) > 0 {
		query += `
			JOIN project_specializations ps ON p.id = ps.project_id
		`
		whereClauses = append(whereClauses, fmt.Sprintf("ps.specialization_id = ANY($%d)", argIdx))
		args = append(args, pq.Array(filters.Specializations.Id))
		argIdx++
	}

	if filters.Thematics != nil && len(filters.Thematics.Id) > 0 {
		query += `
			JOIN project_thematics pt ON p.id = pt.project_id
		`
		whereClauses = append(whereClauses, fmt.Sprintf("pt.thematic_id = ANY($%d)", argIdx))
		args = append(args, pq.Array(filters.Thematics.Id))
		argIdx++
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " ORDER BY p.created_at DESC"

	rows, err := manager.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.ProjectMin
	for rows.Next() {
		var project models.ProjectMin
		var createdAt time.Time
		if err := rows.Scan(&project.Id, &project.Name, &createdAt); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}


	return projects, nil
}