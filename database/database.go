package database

import (
	"btelli-customersupport-app/models"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_CONNECTION_STRING")

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
		&models.User{},
		&models.Role{},
		&models.Customer{},
		&models.ComplaintCategory{},
		&models.Complaint{},
		&models.Comment{},
	)

	if err != nil {
		log.Fatalf("Error while database migration. Error: %v", err)
	}

	log.Println("Database migrated successfully.")
}

func SeedData() {
	SeedRolesData(DB)
	SeedCategoriesData(DB)
	SeedCustomerData(DB)
	SeedComplaintData(DB)
}

func SeedComplaintData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.Complaint{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		complaints := []models.Complaint{
			{
				Description: "Product was delivered late.",
				CustomerID:  1,
				CategoryID:  1,
			},
			{
				Description: "Package arrived damaged.",
				CustomerID:  2,
				CategoryID:  1,
			},
			{
				Description: "Wrong product was sent.",
				CustomerID:  3,
				CategoryID:  2,
			},
			{
				Description: "Return process was not accepted.",
				CustomerID:  4,
				CategoryID:  3,
			},
			{
				Description: "Support hotline is too slow.",
				CustomerID:  5,
				CategoryID:  1,
			},
			{
				Description: "Incomplete product shipment.",
				CustomerID:  6,
				CategoryID:  2,
			},
			{
				Description: "Product advertisement was misleading.",
				CustomerID:  2,
				CategoryID:  1,
			},
			{
				Description: "Invoice was incorrect.",
				CustomerID:  3,
				CategoryID:  2,
			},
			{
				Description: "Not covered under warranty.",
				CustomerID:  4,
				CategoryID:  2,
			},
			{
				Description: "Technical service is inadequate.",
				CustomerID:  5,
				CategoryID:  1,
			},
		}

		for _, complaint := range complaints {
			err := db.Create(&complaint).Error
			if err != nil {
				log.Println("Error while seeding complaint data. Error: ", err)
				return err
			}
		}
	}

	return nil

}

func SeedRolesData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.Customer{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		roles := []models.Role{
			{Name: "Admin"},
			{Name: "User"},
			{Name: "Customer"},
		}

		for _, role := range roles {
			err := db.Create(&role).Error
			if err != nil {
				log.Println("Error while seeding role data. Error: ", err)
				return err
			}
		}
	}

	return nil

}

func SeedCustomerData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.Customer{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		customers := []models.Customer{
			{Name: "John Doe", Email: "john@example.com", Phone: "0 555 333 22 11"},
			{Name: "Jane Smith", Email: "jane@example.com", Phone: "0 555 333 22 12"},
			{Name: "Sam Brown", Email: "sam@example.com", Phone: "0 555 333 22 13"},
			{Name: "Jenny Ack", Email: "jenny@example.com", Phone: "0 555 333 22 14"},
			{Name: "Luis Wutz", Email: "luis@example.com", Phone: "0 555 333 22 15"},
			{Name: "Commander Logar", Email: "logar@example.com", Phone: "0 555 333 22 16"},
		}

		for _, customer := range customers {
			err := db.Create(&customer).Error
			if err != nil {
				log.Println("Error while seeding customer data. Error: ", err)
				return err
			}
		}
	}

	return nil

}

func SeedCategoriesData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.ComplaintCategory{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		categories := []models.ComplaintCategory{
			{
				Name:        "Delivery Issue",
				Description: "Complaints related to delivery problems, such as late or damaged shipments.",
			},
			{
				Name:        "Product Issue",
				Description: "Complaints related to issues with the product, such as defects or incorrect items.",
			},
			{
				Name:        "Customer Service",
				Description: "Complaints related to poor customer service experiences.",
			},
		}

		for _, category := range categories {
			err := db.Create(&category).Error
			if err != nil {
				log.Println("Error while seeding category data. Error: ", err)
				return err
			}
		}
	}

	return nil
}
