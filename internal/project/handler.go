package project

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SolBaa/task-manager/internal/middleware"
	"github.com/SolBaa/task-manager/internal/models"
	"github.com/SolBaa/task-manager/pkg/web"
	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	service ProjectService
}

func NewHandler(service ProjectService) *ProjectHandler {
	return &ProjectHandler{service}
}

// GetAll godoc
// @Summary Get all projects
// @Description Get all projects
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Project
// @Router /projects [get]
func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.RespondJSON(w, http.StatusOK, projects)
}

// GetByID godoc
func (h *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := h.service.GetByID(idInt)
	if err != nil {
		web.Error(w, "Project not found"+err.Error(), http.StatusNotFound)
		return
	}

	web.RespondJSON(w, http.StatusOK, project)
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserKey).(string)
	project.OwnerID = userID

	id, err := h.service.CreateProject(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.RespondJSON(w, http.StatusCreated, map[string]int{"id": id})
}
