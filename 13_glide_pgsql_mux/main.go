package main

import (
	"fmt"
	"go-learn/13_glide_pgsql_mux/src/config"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hi ... mux , pgsql good morning !!!!!")

	db, err := config.GetConnectionDb()

	if err != nil {
		fmt.Print(err.Error())
	}

	p := db.Ping()

	if p != nil {
		fmt.Println(p.Error())
	} else {
		fmt.Println("success !!!!")
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/cars", getCars).Methods("GET")
	r.HandleFunc("/api//cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/api/cars", createCars).Methods("POST")
	r.HandleFunc("/api/cars/{id}", updateCars).Methods("PUT")
	r.HandleFunc("/api/cars/{id}", deleteCars).Methods("DELETE")

}
