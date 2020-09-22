package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/josepmdc/goboilerplate/domain"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func DecodeCredentials(body io.ReadCloser) (*Credentials, error) {
	var c Credentials
	err := json.NewDecoder(body).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON payload: %v", err)
	}
	return &c, nil
}

func MapCredentialsToAPI(credentials *domain.Credentials) *Credentials {
	if credentials == nil {
		return &Credentials{}
	}
	return &Credentials{
		Username: credentials.Username,
		Password: credentials.Password,
		Email:    credentials.Email,
	}
}

func MapCredentialsToDomain(credentials *Credentials) *domain.Credentials {
	if credentials == nil {
		return &domain.Credentials{}
	}
	return &domain.Credentials{
		Username: credentials.Username,
		Password: credentials.Password,
		Email:    credentials.Email,
	}
}
