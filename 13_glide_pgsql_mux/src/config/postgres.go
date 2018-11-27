package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetConnectionDb() (*sql.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=5433", host, user, pass, dbName)

	db, err := createConnection(conn)

	if err != nil {
		return nil, err
	}

	return db, nil

}

func createConnection(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	// db.SetConnMaxLifetime(d)

	return db, nil
}
