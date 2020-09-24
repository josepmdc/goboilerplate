package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/domain"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Score    int       `json:"score"`
	Email    string    `json:"email"`
}

func DecodeUser(body io.ReadCloser) (*User, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON payload: %v", err)
	}
	return &user, nil
}

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
