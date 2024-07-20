package recipe

import (
	"database/sql"

	"github.com/SolBaa/task-manager/internal/models"
)

// RecipeRepository defines the methods that the repository should implement
type RecipeRepository interface {
	GetAll() ([]models.Recipe, error)
	GetByID(id int) (models.Recipe, error)
	CreateRecipe(recipe models.Recipe) (int, error)
}

type recipeRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) RecipeRepository {
	return &recipeRepository{db}
}

func (r *recipeRepository) GetAll() ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := "SELECT * FROM task_manager.Recipes"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (r *recipeRepository) GetByID(id int) (models.Recipe, error) {
	var recipe models.Recipe
	query := "SELECT * FROM task_manager.Recipes WHERE id = ?"
	err := r.DB.QueryRow(query, id).Scan(&recipe.ID, &recipe.Name, &recipe.Description)
	if err != nil {
		return models.Recipe{}, err
	}
	return recipe, nil
}

func (r *recipeRepository) CreateRecipe(recipe models.Recipe) (int, error) {
	query := "INSERT INTO task_manager.Recipes (name, description) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(recipe.Name, recipe.Description)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
