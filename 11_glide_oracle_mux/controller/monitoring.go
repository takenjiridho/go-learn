package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-learn/11_glide_oracle_mux/models"
	"go-learn/11_glide_oracle_mux/repository/schedule"
	"io/ioutil"
	"log"
	"net/http"
)

var schedules []models.Schedule

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetMonitoring
func (c Controller) GetMonitoring(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var schedule models.Schedule
		schedules = []models.Schedule{}
		scheduleRepo := scheduleRepository.ScheduleRepository{}

		// books = bookRepo.GetBooks(db, book, books)

		// for i := 0; i < len(books); i++ {
		// 	fmt.Printf("elemen %d : %s\n", i, books[i])
		// }

		json.NewEncoder(w).Encode(scheduleRepo.GetScheuldes(db, schedule, schedules))
		// json.NewEncoder(w).Encode(bookRepo.GetBooks(db))
	}
}

// GetMonitoringByOrgIdThbl
func (c Controller) GetMonitoringByOrgIdThbl(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var q models.Schedule

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("load ... GetMonitoringByOrgIdThbl ")

		var data2 interface{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		err = json.Unmarshal(body, &data2)
		if err != nil {
			panic(err)
		}
		// log.Println()

		var decodedData = data2.(map[string]interface{})
		fmt.Println("thbl :", decodedData["thbl"])
		fmt.Println("org_id_pemasok  :", decodedData["org_id_pemasok"])

		// thbl, err := strconv.Atoi(decodedData["thbl"])
		// logFatal(err)

		// var schedule models.Schedule
		// json.NewDecoder(r.Body).Decode(&schedule)

		q.Org_id_pemasok = decodedData["org_id_pemasok"]
		q.Thbl = decodedData["thbl"]

		schedules = []models.Schedule{}
		scheedulRepo := scheduleRepository.ScheduleRepository{}

		// book = bookRepo.GetBook(db, book, id)
		json.NewEncoder(w).Encode(scheedulRepo.GetScheuldeByOrgId(db, q))

	}
}
