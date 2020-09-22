package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key") // TODO Add secret key in config

type Credentials struct {
	Username string
	Password string
	Email    string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func GetToken(credentials *Credentials) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // TODO Set the time in config

	claims := &Claims{
		Username: credentials.Username,
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
