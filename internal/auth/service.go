package auth

import (
	"errors"
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

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *authService) Login(user models.UserRequest) (string, error) {

	// Get the user from the database
	passwd, err := s.authRepo.GetUserPasswd(user.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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
