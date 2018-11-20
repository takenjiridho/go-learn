package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// car struct
type Car struct {
	ID          string `json:"id"`
	Manufacture string `json:"manufacture"`
	Product     string `json:"product"`
	Year        string `json:"year"`
	Engine      string `json:"engine"`
	Model       *Model `json:"model"`
}

// Model struc
type Model struct {
	Code  string `json:"code"`
	Type  string `json:"type"`
	Price string `json:"price"`
	Color string `json:"color"`
}

// initiate car var as a slice car struc
var cars []Car

// get all cars
func getCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sameone call me in rest api getCars")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// get specific car
func getCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sameone call me in rest api getCar")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get param from browser

	fmt.Println("params id is : ", params["id"])

	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Car{})
}

// add new car
func createCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sameone call me in rest api createCars")
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(1000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)

}
func updateCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sameone call me in rest api updateCars")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:i], cars[i+1:]...)
			var car Car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
		}
	}

}
func deleteCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sameone call me in rest api updateCars")
	w.Header().Set("Content-TYpe", "application/json")
	params := mux.Vars(r)

	for i, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:i], cars[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cars)

}

// init main function
func main() {
	// say hay
	fmt.Println("hey this rest api")

	// init router
	r := mux.NewRouter()

	// hardcode data cars
	// , Code: "2CRV", Type: "Disel", Price: "550 jt", Color: "Silver", Code: "3CRV", Type: "Practise", Price: "450 jt", Color: "Black"
	// , Code: "2HRV", Type: "Sport", Price: "310 jt", Color: "Brown"
	// , Code: "2DX", Type: "TwinLamp", Price: "10 jt", Color: "white"
	cars = append(cars, Car{ID: "1", Manufacture: "Honda", Product: "CRV", Year: "2018", Engine: "20L", Model: &Model{Code: "1CRV", Type: "Facelift", Price: "500 jt", Color: "white"}})
	cars = append(cars, Car{ID: "2", Manufacture: "Honda", Product: "HRV", Year: "2018", Engine: "15L", Model: &Model{Code: "1HRV", Type: "Facelift", Price: "300 jt", Color: "Red"}})
	cars = append(cars, Car{ID: "3", Manufacture: "Toyota", Product: "DX", Year: "1990", Engine: "15L", Model: &Model{Code: "1DX", Type: "QuardLamp", Price: "20 jt", Color: "Red"}})

	// route and handle endpoint
	r.HandleFunc("/api/cars", getCars).Methods("GET")
	r.HandleFunc("/api//cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/api/cars", createCars).Methods("POST")
	r.HandleFunc("/api/cars/{id}", updateCars).Methods("PUT")
	r.HandleFunc("/api/cars/{id}", deleteCars).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8888", r))

}
