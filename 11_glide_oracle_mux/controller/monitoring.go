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

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("load ... GetMonitoringByOrgIdThbl ")

		body, err := ioutil.ReadAll(r.Body)

		var jsonData = []byte(body)

		var data models.Q

		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Println(err.Error())
		}

		// fmt.Println("org_id_pemasok :", data.Org_id_pemasok)
		// fmt.Println("thbl  :", data.Thbl)

		schedules = []models.Schedule{}
		scheedulRepo := scheduleRepository.ScheduleRepository{}
		json.NewEncoder(w).Encode(scheedulRepo.GetScheuldeByOrgId(db, schedules, data.Org_id_pemasok, data.Thbl))

	}
}
