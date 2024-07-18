package auth

type AuthService interface {
	Login()
	Register()
}

type authService struct {
	authRepo AuthRepository
}

func NewService(authRepo AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Login() {
}

func (s *authService) Register() {
}
