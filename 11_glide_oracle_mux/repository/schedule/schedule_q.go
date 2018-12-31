package scheduleRepository

import (
	"database/sql"
	"go-learn/11_glide_oracle_mux/models"
	"log"
)

type ScheduleRepository struct{}

func logFatal(err error) {
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
		// log.Panic(err)
	}
}

// GetScheuldes
func (b ScheduleRepository) GetScheuldes(db *sql.DB, schedule models.Schedule, schedules []models.Schedule) models.ReturnData {
	// func (b BookRepository) GetBooks(db *sql.DB) models.Rdata {
	rows, err := db.Query("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=201812 and org_id_pemasok='20PBB' ")
	logFatal(err)

	defer rows.Close()
	var v models.ReturnData

	for rows.Next() {
		var c models.Schedule
		err := rows.Scan(&c.TRX_ID, &c.Org_id_pemasok, &c.Thbl, &c.Volume)
		if err != nil {
			v.Status = err.Error()
		} else {
			v.Status = "success"
		}
		v.Data = append(v.Data, c)
	}

	return v
}

// GetScheuldeByOrgId
func (b ScheduleRepository) GetScheuldeByOrgId(db *sql.DB, vorg_id string, vthbl int) models.ReturnData {
	rows, err := db.Query("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=$1 and org_id_pemasok=$2", vthbl, vorg_id)
	logFatal(err)
	defer rows.Close()
	var v models.ReturnData

	for rows.Next() {
		var c models.Schedule

		err := rows.Scan(&c.TRX_ID, &c.Org_id_pemasok, &c.Volume, &c.Thbl)
		if err != nil {
			v.Status = err.Error()
		} else {
			v.Status = "success"
		}

		v.Data = append(v.Data, c)
	}

	return v

}
