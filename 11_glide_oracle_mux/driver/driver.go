package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "gopkg.in/goracle.v2"
)

var db *sql.DB

func connectDB() *sql.DB {

	db, err := sql.Open("goracle", os.Getenv("ORACLE_URL"))
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	return db
}
