package scheduleRepository

import (
	"database/sql"
	"fmt"
	"go-learn/11_glide_oracle_mux/models"

	// "log"

	_ "gopkg.in/goracle.v2"
	// _ "github.com/mattn/go-oci8"
)

type ScheduleRepository struct{}

func logFatal(err error) {
	if err != nil && err != sql.ErrNoRows {
		// log.Fatal(err)
		logFatal(err)
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

// GetScheuldes
func (b ScheduleRepository) GetScheuldeByOrgId(db *sql.DB, schedules []models.Schedule, vorg_id string, vthbl string) models.ReturnData {
	fmt.Println(" thbl : ", vthbl)
	fmt.Println("org id :", vorg_id)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// const qry = "SELECT * FROM user_tab_cols"
	qry := "select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=:1 and org_id_pemasok= :2"
	// cols, err := goracle.QueryContext(ctx, db, qry)
	// if err != nil {
	// 	logFatal(errors.Wrap(err, qry))
	// }

	fmt.Println(qry)

	// thbl := vthbl
	// org_id_pemasok := vorg_id

	// db.QueryContext(goracle.ContextWithLog(ctx, logger.Log), qry)

	// rows, err := db.Query("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=$1", t)
	// rows, err := db.Query("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=? and org_id_pemasok='20PBB'", t)
	// rows, err := db.Query("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=:1 and org_id_pemasok=:2", t, o)
	// rows, err := db.Query(fmt.Sprintf("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=:1 and org_id_pemasok=:2"), t, o)

	// query := "select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=$1 and org_id_pemasok=$2"
	// 	query := "select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl='201811' and org_id_pemasok='20PBB'"
	rows, err := db.Query(qry, vthbl, vorg_id)
	// rows = stmt.Query(t, o)

	// db.QueryContext(goracle.ContextWithLog(ctx, logger.Log), qry)

	// goracle.Log(rows)

	// goracle.DescribeQuery(ctx, db, qry)

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
// func (b ScheduleRepository) GetScheuldeByOrgId(db *sql.DB, vorg_id string, vthbl string) models.ReturnData {

// 	// query := `select trx_id, org_id_pemasok, thbl, volume from "t_fl02a1" where "thbl"= and "org_id_pemasok"=$2`
// 	query := "select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl='201811' and org_id_pemasok='20PBB'"
// 	stmt, err := db.Prepare(query)
// 	rows, err := stmt.Query()

// 	// fmt.Println("stmt ", stmt)

// 	// rows, err := db.Query("select trx_id, org_id_pemasok, thbl, volume from t_fl02a1 where thbl=? and org_id_pemasok=?", vthbl, vorg_id)
// 	logFatal(err)
// 	defer rows.Close()

// 	fmt.Println("jml rows ", rows.)

// 	var v models.ReturnData

// 	for rows.Next() {
// 		var c models.Schedule

// 		err = rows.Scan(&c.TRX_ID, &c.Org_id_pemasok, &c.Volume, &c.Thbl)

// 		fmt.Println(" TRX_ID ", c.TRX_ID)

// 		if err != nil {
// 			v.Status = err.Error()
// 		} else {
// 			v.Status = "success"
// 		}

// 		v.Data = append(v.Data, c)
// 	}

// 	// if err = rows.Err(); err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	return v

// }
