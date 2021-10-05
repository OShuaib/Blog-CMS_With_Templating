package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = 5432
	DB_USER	= "postgres"
	DB_PASS	= "postgres"
	DB_NAME = "blog"
)

func ConnectDb() (*sql.DB, error) {
	conf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", conf)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
		db.Close()
	}
	return db, nil
}