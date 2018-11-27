package main

import (
	"fmt"
	"go-learn/13_glide_pgsql_mux/src/config"
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

}
