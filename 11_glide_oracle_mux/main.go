package main

import (
	"fmt"
	"go-learn/11_glide_oracle_mux/driver"

	// _ "github.com/mattn/go-oci8"
	"github.com/subosito/gotenv"
	_ "gopkg.in/goracle.v2"
)

const dsn = `ICMS/icmsganteng@192.168.100.75:1521/devicms`

func init() {
	gotenv.Load()
}

// var db *sql.DB

func main() {
	// https://jbowens.org/p/using-oracle-from-golang-on-mac-os-x
	// jdbc:oracle:thin:@192.168.100.75:1521:devicms
	// 10.8.0.219

	fmt.Println("glide oracle mux")

	db, err := driver.ConnectDB()
	// ------------------------------------------------------
	//  goracle
	// db, err := sql.Open("goracle", dsn)
	//
	// if err != nil {
	// 	panic(err)
	// }
	// // defer db.Close()
	//
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select sysdate from dual")

	if err != nil {
		fmt.Println("Error running query")
		panic(err)
	}

	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)

	// -----------------------------------------------------------
	// user go-oci
	// db, err := sql.Open("oci8", dsn)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	//
	// rows, err := db.Query("SELECT sysdate FROM dual")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	err = rows.Scan(&name)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

}
