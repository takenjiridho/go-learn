package main

import (
	"database/sql"
	"fmt"
	"go-learn/11_glide_oracle_mux/controller"
	"go-learn/11_glide_oracle_mux/driver"
	"go-learn/11_glide_oracle_mux/models"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	_ "gopkg.in/goracle.v2"
)

var schedules []models.Schedule
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("glide oracle mux port 8000")
	db, err := driver.ConnectDB()
	logFatal(err)

	router := mux.NewRouter()

	c := controller.Controller{}

	router.HandleFunc("/schedule", c.GetMonitoring(db)).Methods("GET")
	router.HandleFunc("/schedule", c.GetMonitoringByOrgIdThbl(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// func main() {
// 	fmt.Println("glide oracle mux")
//
// 	db, err := driver.ConnectDB()
//
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	rows, err := db.Query("select sysdate from dual")
//
// 	if err != nil {
// 		fmt.Println("Error running query")
// 		panic(err)
// 	}
//
// 	defer rows.Close()
//
// 	var thedate string
// 	for rows.Next() {
//
// 		rows.Scan(&thedate)
// 	}
// 	fmt.Printf("The date is: %s\n", thedate)
//
// }
