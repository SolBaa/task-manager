package auth

import "database/sql"

type AuthRepository interface {
	Login()
	Register()
}

type authRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) AuthService {
	return &authRepository{db}
}

func (r *authRepository) Login() {
}

func (r *authRepository) Register() {
}
