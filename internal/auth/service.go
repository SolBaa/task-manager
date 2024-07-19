package auth

import (
	"os"
	"time"

	"github.com/SolBaa/task-manager/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(models.UserRequest) (string, error)
	Register(string, string, string) error
}

type authService struct {
	authRepo AuthRepository
}

func NewService(authRepo AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Login(user models.UserRequest) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenjwt := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(tokenjwt))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func (s *authService) Register(email, username, password string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.authRepo.Register(email, username, hashedPassword)
}
