package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "gopkg.in/goracle.v2"
	// _ "github.com/mattn/go-oci8"
)

// const dsn = `ICMS/icmsganteng@192.168.100.75:1521/devicms`

// const dsn = `ICMS/icmsganteng@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=192.168.100.75)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=devicms)))`

// var db *sql.DB

func ConnectDB() (*sql.DB, error) {

	// fmt.Println("connection string " + os.Getenv("ORACLE_URL"))
	fmt.Println("driver oracle load..... " + os.Getenv("ORACLE_URL"))

	db, err := sql.Open("goracle", os.Getenv("ORACLE_URL"))
	// db, err := sql.Open("oci8", dsn)

	// if err != nil {
	// 	fmt.Println("errrrror ..... ")
	// 	fmt.Println(err)
	// }
	// defer db.Close()
	//
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println("ping ..... ")
	// 	fmt.Println(err)
	// }

	return db, err
}
