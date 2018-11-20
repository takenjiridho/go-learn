package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetPostgresDB(*sql.DB, err) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	db := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, db)

	db, err := createConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConnection(desc string) (*sql.DB, err) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
