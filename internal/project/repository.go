package project

import (
	"database/sql"

	"github.com/SolBaa/task-manager/internal/models"
)

type RepositoryProject interface {
	GetAll() ([]models.Project, error)
	GetByID(id int) (models.Project, error)
	CreateProject(project models.Project) (int, error)
}

type repositoryProject struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) RepositoryProject {
	return &repositoryProject{db}
}

func (r *repositoryProject) GetAll() ([]models.Project, error) {
	var projects []models.Project
	query := "SELECT * FROM task_manager.Projects"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (r *repositoryProject) GetByID(id int) (models.Project, error) {
	var project models.Project
	query := "SELECT * FROM task_manager.Projects WHERE id = ?"
	err := r.DB.QueryRow(query, id).Scan(&project.ID, &project.Name, &project.Description)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}

func (r *repositoryProject) CreateProject(project models.Project) (int, error) {
	query := "INSERT INTO task_manager.Projects (name, description) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(project.Name, project.Description)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
