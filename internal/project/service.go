package project

import "github.com/SolBaa/task-manager/internal/models"

type ProjectService interface {
	GetAll() ([]models.Project, error)
	GetByID(id int) (models.Project, error)
	CreateProject(project models.Project) (int, error)
}

type projectService struct {
	repository RepositoryProject
}

func NewService(repository RepositoryProject) ProjectService {
	return &projectService{repository}
}

func (s *projectService) GetAll() ([]models.Project, error) {
	projects, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return projects, nil

}

func (s *projectService) GetByID(id int) (models.Project, error) {
	project, err := s.repository.GetByID(id)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}

func (s *projectService) CreateProject(project models.Project) (int, error) {
	id, err := s.repository.CreateProject(project)
	if err != nil {
		return 0, err
	}
	return id, nil
}
