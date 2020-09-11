package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/domain"
	"io"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	FullName string    `json:"full_name"`
	Score    int       `json:"score"`
	Email    string    `json:"email"`
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func MapUserToAPI(user *domain.User) *User {
	if user == nil {
		return &User{}
	}
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
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
		UserName: user.UserName,
		FullName: user.FullName,
		Score:    user.Score,
		Email:    user.Email,
	}
}

func DecodeUser(body io.ReadCloser) (*User, error) {
	var user User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON payload: %v", err)
	}
	return &user, nil
}

func DecodeCredentials(body io.ReadCloser) (*Credentials, error) {
	var c Credentials
	err := json.NewDecoder(body).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON payload: %v", err)
	}
	return &c, nil
}
