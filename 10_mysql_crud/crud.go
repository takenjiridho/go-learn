package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3308)/blogapp")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlQueryRow() {
	var db, err = connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var r = student{}
	var id = "E001"

	err = db.QueryRow("select name,grade from tb_student where id = ?", id).Scan(&r.name, &r.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name : %s\ngrade : %d\n", r.name, r.grade)
}

func sqlPrepare() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("select name ,grade from tb_student where id =?")

	if err != nil {
		fmt.Println(err.Error())
	}

	r1 := student{}
	stmt.QueryRow("E001").Scan(&r1.name, &r1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", r1.name, r1.grade)

	r2 := student{}
	stmt.QueryRow("W001").Scan(&r2.name, &r2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", r2.name, r2.grade)

	r3 := student{}
	stmt.QueryRow("B001").Scan(&r3.name, &r3.grade)
	fmt.Printf("name: %s\ngrade: %d\n", r3.name, r3.grade)

}

func sqlExec() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?,?,?,?)", "G001", "Golangbro", 29, 2)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("insert success!!")

	_, err = db.Exec("update tb_student set age=? where id=?", 28, "G001")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("update success!!")

	_, err = db.Exec("delete from tb_student where id =?", "G001")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("delete success!!")

}
func main() {
	fmt.Println("test")

	sqlExec()

	// sqlPrepare()
	// sqlQuery()
	// sqlQueryRow()
}
