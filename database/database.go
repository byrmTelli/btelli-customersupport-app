package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_CONNECTION_STRING")

	fmt.Println("dsn", dsn)
	if dsn == "" {
		log.Fatal("Db connection string not found.")
	}

	// Connect to db
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Something went wrong while connecting to db. Error: %v", err)
	}

	DB = db

	log.Println("Database connection initialized succesfully.")

	MigrateDatabase()

}

func MigrateDatabase() {
	err := DB.AutoMigrate(
		// Database Models to create related table.
		&User{},
		&Role{},
		&Customer{},
		&ComplaintCategory{},
		&Complaint{},
		&Comment{},
	)

	if err != nil {
		log.Fatalf("Error while database migration. Error: %v", err)
	}

	log.Println("Database migrated successfully.")
}
