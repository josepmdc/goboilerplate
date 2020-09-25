package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key") // TODO Add secret key in config

// Credentials defines the necessary parameters for a user to
// sign in and sign up
type Credentials struct {
	Username string
	Password string
	Email    string
}

// Claims defines the values we will add to the JWT claims
type Claims struct {
	Username string
	jwt.StandardClaims
}

// GetToken returns a JWT token with the username as a claim
func GetToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // TODO Set the time in config

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates if a JWT is correct
func ValidateToken(token string) (bool, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}
	if !tkn.Valid {
		return false, nil
	}
	return true, nil
}

// Validate checks that the required values of the credentials are filled and that they
// meet specific  requirements
func (credentials *Credentials) Validate() bool {
	if credentials.Username == "" {
		return false
	}
	if credentials.Password == "" || len(credentials.Password) < 8 {
		return false
	}
	if credentials.Email == "" {
		return false
	}
	return true
}
