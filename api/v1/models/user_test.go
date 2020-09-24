package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/domain"
)

var ID = uuid.Must(uuid.NewRandom())

const (
	USER_NAME = "asdf"
	FULL_NAME = "asd asdf asdf"
	SCORE     = 24
	EMAIL     = "asdf@gmail.com"
)

func TestMapUserToAPI(t *testing.T) {
	result := MapUserToAPI(GetDomainUser())
	expected := GetApiUser()

	if result == nil {
		t.Error("Mapped user is nil")
	}

	if *result != *expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToAPI(nil)
	expected = &User{}
	if *result != *expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToAPI(&domain.User{})
	expected = &User{}
	if *result != *expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}
}

func TestMapUserToDomain(t *testing.T) {
	result := MapUserToDomain(GetApiUser())
	expected := GetDomainUser()

	if result == nil {
		t.Error("Mapped user is nil")
	}

	if *result != *expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToDomain(nil)
	expected = &domain.User{}
	if *result != *expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}

	result = MapUserToDomain(&User{})
	expected = &domain.User{}
	if *result != *expected {
		t.Errorf("Mapped user is different than expected:\n result: %+v \n expected: %+v", result, expected)
	}
}

func GetDomainUser() *domain.User {
	return &domain.User{
		ID:       ID,
		Username: USER_NAME,
		FullName: FULL_NAME,
		Score:    SCORE,
		Email:    EMAIL,
	}
}

func GetApiUser() *User {
	return &User{
		ID:       ID,
		Username: USER_NAME,
		FullName: FULL_NAME,
		Score:    SCORE,
		Email:    EMAIL,
	}
}
