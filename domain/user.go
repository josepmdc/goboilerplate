package domain

import "github.com/google/uuid"

// User is the domain entity that defines a user of the application
type User struct {
	ID       uuid.UUID
	FullName string
	UserName string
	Password string
	Email    string
	Score    int
	GitHubID string
}

// UserRepo is the interface that defines all of the operations the User entity
// can do with the persistance layer
type UserRepo interface {
	FindByUsername(userName string) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
	FindAll() (*[]User, error)
	Create(u *User) (*User, error)
	CheckEmail(email string) bool
	CheckUsername(username string) bool
}

// Validate checks that the required values of the user are filled and that they
// meet specific  requirements
func (user *User) Validate() bool {
	if user.UserName == "" {
		return false
	}
	if user.Password == "" || len(user.Password) < 6 {
		return false
	}
	if user.Email == "" {
		return false
	}
	return true
}
