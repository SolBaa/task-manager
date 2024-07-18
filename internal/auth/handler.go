package auth

import (
	"net/http"
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
	w.Write([]byte("Login"))
}

func (h *LoginHandler) Register(w http.ResponseWriter, r *http.Request) {
}
