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

	// Obtener los ingredientes
	for i := range recipes {
		ingredientQuery := "SELECT * FROM task_manager.Ingredients WHERE recipe_id = ?"
		ingredientRows, err := r.DB.Query(ingredientQuery, recipes[i].ID)
		if err != nil {
			return nil, err
		}
		defer ingredientRows.Close()
		for ingredientRows.Next() {
			var ingredient models.Ingredient
			err := ingredientRows.Scan(&ingredient.ID, &ingredient.RecipeID, &ingredient.Name, &ingredient.Quantity)
			if err != nil {
				return nil, err
			}
			recipes[i].Ingredients = append(recipes[i].Ingredients, ingredient)
		}
	}
	return recipes, nil
}

func (r *recipeRepository) GetByID(id int) (models.Recipe, error) {
	var recipe models.Recipe
	query := `
	SELECT r.id, r.name, r.description, i.id, i.recipe_id, i.name, i.quantity
	FROM task_manager.Recipes r
	LEFT JOIN task_manager.Ingredients i ON r.id = i.recipe_id
	WHERE r.id = ?`
	err := r.DB.QueryRow(query, id).Scan(&recipe.ID, &recipe.Name, &recipe.Description)
	if err != nil {
		return models.Recipe{}, err
	}

	// Obtener los ingredientes
	ingredientQuery := "SELECT * FROM task_manager.Ingredients WHERE recipe_id = ?"
	ingredientRows, err := r.DB.Query(ingredientQuery, id)
	if err != nil {
		return models.Recipe{}, err
	}
	defer ingredientRows.Close()
	for ingredientRows.Next() {
		var ingredient models.Ingredient
		err := ingredientRows.Scan(&ingredient.ID, &ingredient.RecipeID, &ingredient.Name, &ingredient.Quantity)
		if err != nil {
			return models.Recipe{}, err
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	return recipe, nil
}

func (r *recipeRepository) CreateRecipe(recipe models.Recipe) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Insertar la receta
	recipeQuery := "INSERT INTO task_manager.Recipes (name, description) VALUES (?, ?)"
	recipeStmt, err := tx.Prepare(recipeQuery)
	if err != nil {
		return 0, err
	}
	defer recipeStmt.Close()

	res, err := recipeStmt.Exec(recipe.Name, recipe.Description)
	if err != nil {
		return 0, err
	}

	recipeID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Insertar los ingredientes
	ingredientQuery := "INSERT INTO task_manager.Ingredients (recipe_id, name, quantity) VALUES (?, ?, ?)"
	ingredientStmt, err := tx.Prepare(ingredientQuery)
	if err != nil {
		return 0, err
	}
	defer ingredientStmt.Close()

	for _, ingredient := range recipe.Ingredients {
		_, err = ingredientStmt.Exec(recipeID, ingredient.Name, ingredient.Quantity)
		if err != nil {
			return 0, err
		}
	}

	return int(recipeID), nil
}
