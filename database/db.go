package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=psql_db user=admin password=admin dbname=gofiber_db port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return
	}
	fmt.Println("Successfully connected to the database!")
}
