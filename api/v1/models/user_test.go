package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/domain"
)

var ID = uuid.Must(uuid.NewRandom())

const (
	userName = "asdf"
	fullName = "asd asdf asdf"
	score    = 24
	email    = "asdf@gmail.com"
)

func TestMapUserToAPI(t *testing.T) {
	u := getDomainUser()
	result := MapUserToAPI(&u)
	expected := getAPIUser()

	if *result != expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToAPI(nil)
	expected = User{}
	if *result != expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToAPI(&domain.User{})
	expected = User{}
	if *result != expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}
}

func TestMapUserToDomain(t *testing.T) {
	u := getAPIUser()
	result := MapUserToDomain(&u)
	expected := getDomainUser()

	if *result != expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToDomain(nil)
	expected = domain.User{}
	if *result != expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToDomain(&User{})
	expected = domain.User{}
	if *result != expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}
}

func getDomainUser() domain.User {
	return domain.User{
		ID:       ID,
		Username: userName,
		FullName: fullName,
		Score:    score,
		Email:    email,
	}
}

func getAPIUser() User {
	return User{
		ID:       ID,
		Username: userName,
		FullName: fullName,
		Score:    score,
		Email:    email,
	}
}
