package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabse connections - returns a pointer  to a database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up the database connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil

}
