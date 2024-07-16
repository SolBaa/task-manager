package project

type ProjectService interface {
}

type projectService struct {
}

func NewProjectService() ProjectService {
	return &projectService{}
}
