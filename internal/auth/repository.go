package auth

import (
	"database/sql"
	"fmt"
)

type AuthRepository interface {
	Login()
	Register(string, string, []byte) error
	GetUserPasswd(string) (string, error)
}

type authRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Login() {
}

func (r *authRepository) Register(email, username string, password []byte) error {
	query := "INSERT INTO task_manager.Users (email,username, password, role) VALUES (?,?, ?, 'user')"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(email, username, password)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil

}

func (r *authRepository) GetUserPasswd(username string) (string, error) {
	var password string
	query := "SELECT password FROM task_manager.Users WHERE username = ?"
	err := r.DB.QueryRow(query, username).Scan(&password)
	if err != nil {
		return "", fmt.Errorf("failed to get user password: %v", err)
	}
	return password, nil
}
