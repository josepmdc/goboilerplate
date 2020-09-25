package app

import (
	"github.com/josepmdc/goboilerplate/domain"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is an interface that defines the opearations related to
// authentication
type AuthService interface {
	SignIn(credentials *domain.Credentials) (string, error)
}

type authService struct {
	userService UserService
}

// NewAuthService creates a new AuthService implementation
func NewAuthService(us UserService) (AuthService, error) {
	return &authService{
		userService: us,
	}, nil
}

func (as *authService) SignIn(credentials *domain.Credentials) (string, error) {
	u, err := as.userService.FindByUsername(credentials.Username)
	if err != nil {
		return "", err
	}

	plainPassword := []byte(credentials.Password)
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), plainPassword)
	if err != nil {
		return "", err
	}

	token, err := domain.GetToken(credentials.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
