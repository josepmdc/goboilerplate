package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/conf"
	"github.com/josepmdc/goboilerplate/domain"
	"github.com/josepmdc/goboilerplate/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

// UserService is an interface that defines the opearations that you can do
// with the User domain entity
type UserService interface {
	FindUser(ID uuid.UUID) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	CheckEmail(email string) bool
	CheckUsername(username string) bool
}

type userService struct {
	UserRepo domain.UserRepo
}

// NewUserService creates a new UserService implementation
func NewUserService(conf *conf.Config) (UserService, error) {
	db, err := infrastructure.NewDB(&conf.PostgresConfig)
	if err != nil {
		return nil, err
	}
	return &userService{
		UserRepo: infrastructure.NewUserRepo(db),
	}, nil
}

func (us *userService) FindUser(ID uuid.UUID) (*domain.User, error) {
	user, err := us.UserRepo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) CreateUser(user *domain.User) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	ok := user.Validate()
	if !ok {
		return nil, errors.New("Invalid user")
	}

	if _, err = us.UserRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) CheckEmail(email string) bool {
	return us.UserRepo.CheckEmail(email)
}

func (us *userService) CheckUsername(username string) bool {
	return us.UserRepo.CheckUsername(username)
}
