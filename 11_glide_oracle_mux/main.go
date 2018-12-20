package main

import (
	"database/sql"
	"fmt"
	"go-learn/11_glide_oracle_mux/driver"

	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	fmt.Println("glide oracle mux")

	db = driver.ConnectDB()

	rows, err := db.Query("select sysdate from dual")

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		// return
	}

	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)

}
