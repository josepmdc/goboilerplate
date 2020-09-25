package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/domain"
)

// User defienes the user info that will be exposed through the API
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Score    int       `json:"score"`
	Email    string    `json:"email"`
}

// DecodeUser takes the request body in JSON format
// and decodes it into the User struct
func DecodeUser(body io.ReadCloser) (*User, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON payload: %v", err)
	}
	return &user, nil
}

// MapUserToAPI takes a domain object of User and maps it to the API model
func MapUserToAPI(user *domain.User) *User {
	if user == nil {
		return &User{}
	}
	return &User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Score:    user.Score,
		Email:    user.Email,
	}
}

// MapUserToDomain takes an API object of User and maps it to the domain model
func MapUserToDomain(user *User) *domain.User {
	if user == nil {
		return &domain.User{}
	}
	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Score:    user.Score,
		Email:    user.Email,
	}
}
