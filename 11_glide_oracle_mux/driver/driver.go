package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "gopkg.in/goracle.v2"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("driver oracle load..... " + os.Getenv("ORACLE_URL"))

	db, err := sql.Open("goracle", os.Getenv("ORACLE_URL"))
	return db, err
}
