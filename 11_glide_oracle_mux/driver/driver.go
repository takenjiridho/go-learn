package driver

import (
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
	// _ "github.com/mattn/go-oci8"
)

func ConnectDB() (*sql.DB, error) {
	/*
		//-- enable if ide vscode and depoly
		fmt.Println("driver oracle load..... " + os.Getenv("ORACLE_URL"))
		db, err := sql.Open("goracle", os.Getenv("ORACLE_URL"))
		db, err := sql.Open("oci8", os.Getenv("ORACLE_URL"))
	*/

	//-- enable if ide goland
	ORACLE_URL_1 := "ICMS/icmsganteng@192.168.100.75:1521/devicms"
	fmt.Println("driver oracle load..... " + ORACLE_URL_1) //
	db, err := sql.Open("goracle", ORACLE_URL_1)
	return db, err
}
