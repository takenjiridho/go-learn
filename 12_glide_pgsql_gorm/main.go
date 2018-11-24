package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// pProfile
type Profile struct {
	// gorm.Model
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func main() {
	fmt.Println("gorm + postgres + mux + glide")

	db, err := getPostgresDB()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.DB().Ping()

	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	fmt.Println("connection success !!!")

	var profile Profile
	// fmt.Println(db.First(&profile, "email = ?", "farid@gmail.com")) // find product with code l1212
	db.First(&profile)

	fmt.Println("first name ", profile.FirstName)
}

func createConnection(setting string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", setting)

	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	// db.DB().SetConnMaxLifetime(time.Hour)

	return db, nil
}

// gGetPostgresDB()
func getPostgresDB() (*gorm.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")

	s := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable port=5433", user, pass, host, dbName)

	db, err := createConnection(s)

	if err != nil {
		return nil, err
	}

	return db, nil

}
