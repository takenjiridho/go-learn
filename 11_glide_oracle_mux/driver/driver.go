package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "gopkg.in/goracle.v2"
	// _ "github.com/mattn/go-oci8"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("driver oracle load..... " + os.Getenv("ORACLE_URL"))

	// db, err := sql.Open("oci8", os.Getenv("ORACLE_URL"))
	db, err := sql.Open("goracle", os.Getenv("ORACLE_URL"))
	return db, err
}
