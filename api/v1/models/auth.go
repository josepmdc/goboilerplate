package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/josepmdc/goboilerplate/domain"
)

// Credentials defines the necessary parameters for a user to
// sign in and sign up
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// DecodeCredentials takes the request body in JSON format
// and decodes it into the Credentials struct
func DecodeCredentials(body io.ReadCloser) (*Credentials, error) {
	var c Credentials
	err := json.NewDecoder(body).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("Invalid JSON payload: %v", err)
	}
	return &c, nil
}

// MapCredentialsToAPI takes a domain object of credentials and maps it to the API model
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

// MapCredentialsToDomain takes an API object of credentials and maps it to the domain model
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
