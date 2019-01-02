package bookRepository

import (
	"database/sql"
	"fmt"
	"go-learn/91_codix/books-list/models"
	"log"
)

type BookRepository struct{}

type Ldata struct {
	Status string        `json:"status"`
	Data   []models.Book `json:"data"`
}

func logFatal(err error) {
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)

		// log.Panic(err)
	}
}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) models.Rdata {
	// func (b BookRepository) GetBooks(db *sql.DB) models.Rdata {
	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	var v models.Rdata

	for rows.Next() {

		var c models.Book

		err := rows.Scan(&c.ID, &c.Title, &c.Author, &c.Year)
		if err != nil {
			v.Status = err.Error()
		} else {
			v.Status = "success"
		}

		v.Data = append(v.Data, c)

		// books = append(books, book)
	}

	return v
}

// func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) Ldata {
func (b BookRepository) GetBook(db *sql.DB, id int) Ldata {
	rows := db.QueryRow("select * from books where id=$1", id)

	var v Ldata
	var c models.Book

	// err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	err := rows.Scan(&c.ID, &c.Title, &c.Author, &c.Year)

	if err != nil {
		v.Status = err.Error()
		fmt.Println(err.Error())
		v.Data = append(v.Data, c)
		// v.Data = append(v.Data, book)
	} else {
		v.Status = "success"
		// v.Data = append(v.Data, book)
		v.Data = append(v.Data, c)
	}

	// if err != sql.ErrNoRows {
	// 	fmt.Println(sql.ErrNoRows.Error())
	// }

	// logFatal(err)

	return v
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	logFatal(err)

	return book.ID
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
