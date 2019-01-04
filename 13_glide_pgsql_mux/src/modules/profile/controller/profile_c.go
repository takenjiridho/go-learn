package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-learn/13_glide_pgsql_mux/src/modules/profile/repository"
	"log"
	"net/http"
)

// var profiles []model.Profile

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetMonitoring
func (c Controller) GetProfile(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		// var profile model.Profile
		// profiles = []model.Profile{}
		// profileRepo := repository.ProfileRepository{}
		profileRepo := repository.NewProfileImpl(db)

		p, e := profileRepo.FindAll("P2", "wjd")

		if e != nil {
			fmt.Println(e.Error())
		}

		// books = bookRepo.GetBooks(db, book, books)

		// for i := 0; i < len(books); i++ {
		// 	fmt.Printf("elemen %d : %s\n", i, books[i])
		// }

		json.NewEncoder(w).Encode(p)
		// json.NewEncoder(w).Encode(bookRepo.GetBooks(db))
	}
}
