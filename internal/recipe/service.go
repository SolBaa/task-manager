package recipe

import "github.com/SolBaa/task-manager/internal/models"

// RecipeService is a service that provides recipe-related operations.
type RecipeService interface {
	// GetAll returns all recipes.
	GetAll() ([]models.Recipe, error)
	// GetByID returns a recipe by its ID.
	GetByID(id int) (models.Recipe, error)
	// CreateRecipe creates a new recipe.
	CreateRecipe(recipe models.Recipe) (int, error)
}

// recipeService is a service that provides recipe-related operations.
type recipeService struct {
	repository RecipeRepository
}

// NewService creates a new recipe service.
func NewService(repository RecipeRepository) RecipeService {
	return &recipeService{repository}
}

// GetAll returns all recipes.
func (s *recipeService) GetAll() ([]models.Recipe, error) {
	recipes, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return recipes, nil

}

// GetByID returns a recipe by its ID.
func (s *recipeService) GetByID(id int) (models.Recipe, error) {
	recipe, err := s.repository.GetByID(id)
	if err != nil {
		return models.Recipe{}, err
	}
	return recipe, nil
}

// CreateRecipe creates a new recipe.
func (s *recipeService) CreateRecipe(recipe models.Recipe) (int, error) {

	id, err := s.repository.CreateRecipe(recipe)
	if err != nil {
		return 0, err
	}
	return id, nil
}
