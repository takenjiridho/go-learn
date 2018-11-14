package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sameone call me in rest api")
}

func getCar(w http.ResponseWriter, r *http.Request) {

}

func createCars(w http.ResponseWriter, r *http.Request) {

}

func updateCars(w http.ResponseWriter, r *http.Request) {

}
func deleteCars(w http.ResponseWriter, r *http.Request) {

}

// init main function
func main() {
	// say hay
	fmt.Println("hey this rest api")

	// init router
	r := mux.NewRouter()

	// route and handle endpoint
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/cars", createCars).Methods("POST")
	r.HandleFunc("/cars/{id}", updateCars).Methods("PUT")
	r.HandleFunc("/cars/{id}", deleteCars).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8888", r))

}
