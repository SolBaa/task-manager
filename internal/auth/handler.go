package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SolBaa/task-manager/internal/models"
)

func NewHandler(authService AuthService) *LoginHandler {
	return &LoginHandler{
		authService: authService,
	}
}

type LoginHandler struct {
	authService AuthService
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

func (h *LoginHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	err = h.authService.Register(user.Email, user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
