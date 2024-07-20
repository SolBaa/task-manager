package recipe

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SolBaa/task-manager/internal/models"
	"github.com/SolBaa/task-manager/pkg/web"
	"github.com/go-chi/chi/v5"
)

type recipeHanlder struct {
	service RecipeService
}

// NewHandler creates a new recipe handler.
func NewHandler(service RecipeService) *recipeHanlder {
	return &recipeHanlder{service}
}

// GetAll returns all recipes.
func (h *recipeHanlder) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.RespondJSON(w, http.StatusOK, recipes)
}

// GetByID returns a recipe by its ID.
func (h *recipeHanlder) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipe, err := h.service.GetByID(idInt)
	if err != nil {
		web.Error(w, "Recipe not found"+err.Error(), http.StatusNotFound)
		return
	}

	web.RespondJSON(w, http.StatusOK, recipe)
}

// CreateRecipe creates a new recipe.
func (h *recipeHanlder) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.RespondJSON(w, http.StatusCreated, map[string]int{"id": id})
}
