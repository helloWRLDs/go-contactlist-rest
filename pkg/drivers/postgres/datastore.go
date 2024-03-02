package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresCFG struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func GetPostgresConnection(cfg *PostgresCFG) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
