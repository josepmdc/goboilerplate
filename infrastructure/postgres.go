package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/josepmdc/goboilerplate/conf"

	// Blank import to use the Postgres SQL driver
	_ "github.com/lib/pq"
)

// NewDB creates a new postgres *sql.DB instance
func NewDB(pg *conf.PostgresConfig) (*sql.DB, error) {
	pgInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.Database)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
