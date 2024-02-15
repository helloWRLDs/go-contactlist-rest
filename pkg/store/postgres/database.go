package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetPostgresConnection(host, port, user, password, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbName,
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
