package infrastructure

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/josepmdc/goboilerplate/domain"
	"github.com/josepmdc/goboilerplate/log"
)

type pgUserRepo struct {
	DB *sql.DB
}

// NewUserRepo creates a new pgUser repo, wich is a Postgres implementation of
// the UserRepo interface
func NewUserRepo(db *sql.DB) domain.UserRepo {
	return &pgUserRepo{
		DB: db,
	}
}

func (repo *pgUserRepo) FindByUsername(username string) (*domain.User, error) {
	var ID uuid.UUID
	var fullName, email, password string
	var isAdmin bool
	var score int
	var createdAt time.Time
	err := repo.DB.QueryRow(`SELECT id, email, full_name, admin, created_at, score, password FROM users WHERE username = $1`, username).
		Scan(&ID, &email, &fullName, &isAdmin, &createdAt, &score, &password)

	if err != nil {
		log.Logger.Warnf("Could not get user from database: %s", err.Error())
		return nil, err
	}

	user := domain.User{
		ID:       ID,
		FullName: fullName,
		Username: username,
		Email:    email,
		Score:    score,
		Password: password,
	}
	return &user, nil
}

// FindByID searches on the database for a user with an ID. If it doesn't find
// it returns a nil object
func (repo *pgUserRepo) FindByID(ID uuid.UUID) (*domain.User, error) {
	var fullName, username, email string
	var isAdmin bool
	var score int
	var createdAt time.Time
	err := repo.DB.QueryRow(`SELECT username, email, full_name, admin, created_at, score FROM users WHERE id = $1`, ID).
		Scan(&username, &email, &fullName, &isAdmin, &createdAt, &score)

	if err != nil {
		log.Logger.Warnf("Could not get user from database: %s", err.Error())
		return nil, err
	}

	user := domain.User{
		ID:       ID,
		FullName: fullName,
		Username: username,
		Email:    email,
		Score:    score,
	}
	return &user, nil
}

func (repo *pgUserRepo) FindAll() (*[]domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *pgUserRepo) Create(u *domain.User) (*domain.User, error) {
	var err error
	u.ID, err = uuid.NewRandom()
	if err != nil {
		log.Logger.Warnf("Could not insert user %s into the database: %s", u.Username, err.Error())
		return nil, err
	}

	_, err = repo.DB.Exec(`INSERT INTO users (id, username, password, email, full_name) VALUES ($1, $2, $3, $4, $5)`,
		u.ID, u.Username, u.Password, u.Email, u.FullName)
	if err != nil {
		log.Logger.Warnf("Could not insert user %s into the database: %s", u.Username, err.Error())
		return nil, err
	}
	return u, nil
}

func (repo *pgUserRepo) CheckEmail(email string) bool {
	err := repo.DB.QueryRow(`SELECT email FROM users WHERE email = $1`, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Logger.Warnf("Unexpected error: %s", err.Error())
		}
		return false
	}
	return true
}

func (repo *pgUserRepo) CheckUsername(username string) bool {
	err := repo.DB.QueryRow(`SELECT username FROM users WHERE username = $1`, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Logger.Warnf("Unexpected error: %s", err.Error())
		}
		return false
	}
	return true
}
