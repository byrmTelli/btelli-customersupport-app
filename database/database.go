package database

import (
	"btelli-customersupport-app/models"
	"btelli-customersupport-app/utils"
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
	err := seedRolesData(DB)
	if err != nil {
		log.Println("Error seeding roles:", err)
		return
	}

	err = seedCategoriesData(DB)
	if err != nil {
		log.Println("Error seeding categories:", err)
		return
	}

	err = seedUserData(DB)
	if err != nil {
		log.Println("Error seeding users:", err)
		return
	}

	err = seedComplaintData(DB)
	if err != nil {
		log.Println("Error seeding complaints:", err)
		return
	}

	err = seedCommentData(DB)
	if err != nil {
		log.Println("Error seeding comments:", err)
		return
	}
}

func seedCommentData(db *gorm.DB) error {
	var count int64
	err := db.Model(&models.Comment{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		comments := []models.Comment{
			{
				ComplaintID: 5,
				UserID:      2,
				CommentText: "Will be  fixed soon.",
			},
			{
				ComplaintID: 5,
				UserID:      5,
				CommentText: "Alright! Thank you for your efforts.",
			},
			{
				ComplaintID: 10,
				UserID:      2,
				CommentText: "Our team will be handle this problem between 10 or 15 days.",
			},
			{
				ComplaintID: 10,
				UserID:      5,
				CommentText: "C'mon is too long to solve.",
			},
			{
				ComplaintID: 10,
				UserID:      5,
				CommentText: "Have to be solved quicker!",
			},
		}
		for _, comment := range comments {
			err := db.Create(&comment).Error
			if err != nil {
				log.Println("Error while seeding comment data. Error: ", err)
				return err
			}
		}
	}

	return nil
}

func seedComplaintData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.Complaint{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		complaints := []models.Complaint{
			{
				Title:       "Product was delivered late.",
				Description: "Product was delivered late.",
				UserID:      3,
				CategoryID:  1,
				Status:      models.Resolved,
			},
			{
				Title:       "Package arrived damaged.",
				Description: "Package arrived damaged.",
				UserID:      3,
				CategoryID:  1,
				Status:      models.Resolved,
			},
			{
				Title:       "Wrong product was sent.",
				Description: "Wrong product was sent.",
				UserID:      3,
				CategoryID:  2,
				Status:      models.Cancelled,
			},
			{
				Title:       "Return process was not accepted.",
				Description: "Return process was not accepted.",
				UserID:      4,
				CategoryID:  3,
				Status:      models.Cancelled,
			},
			{
				Title:       "Support hotline is too slow.",
				Description: "Support hotline is too slow.",
				UserID:      5,
				CategoryID:  1,
				Status:      models.InProgress,
			},
			{
				Title:       "Incomplete product shipment.",
				Description: "Incomplete product shipment.",
				UserID:      6,
				CategoryID:  2,
				Status:      models.InProgress,
			},
			{
				Title:       "Product advertisement was misleading.",
				Description: "Product advertisement was misleading.",
				UserID:      4,
				CategoryID:  1,
				Status:      models.InProgress,
			},
			{
				Title:       "Invoice was incorrect.",
				Description: "Invoice was incorrect.",
				UserID:      3,
				CategoryID:  2,
				Status:      models.InProgress,
			},
			{
				Title:       "Not covered under warranty.",
				Description: "Not covered under warranty.",
				UserID:      7,
				CategoryID:  2,
				Status:      models.InProgress,
			},
			{
				Title:       "Technical service is inadequate.",
				Description: "Technical service is inadequate.",
				UserID:      5,
				CategoryID:  1,
				Status:      models.InProgress,
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

func seedRolesData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.Role{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		roles := []models.Role{
			{Name: "Admin"},
			{Name: "Help Desk"},
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

func seedUserData(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return err
	}
	password, err := utils.HashPassword("Password123.")

	if err != nil {
		log.Fatalf("Error occured while hashing seeder passwords.")
	}

	if count == 0 {
		users := []models.User{
			{
				UserName:     "bayramtelli",
				Name:         "Bayram",
				Surname:      "Telli",
				Email:        "bayram@example.com",
				Phone:        "0 555 333 22 10",
				PasswordHash: password,
				RoleID:       1,
			},

			{
				UserName:     "johndoe",
				Name:         "John",
				Surname:      "Doe",
				Email:        "john@example.com",
				Phone:        "0 555 333 22 11",
				PasswordHash: password,
				RoleID:       2,
			},
			{
				UserName:     "janesmith",
				Name:         "Jane",
				Surname:      "Smith",
				Email:        "jane@example.com",
				Phone:        "0 555 333 22 12",
				PasswordHash: password,
				RoleID:       3,
			},
			{
				UserName:     "sambrown",
				Name:         "Sam",
				Surname:      "Brown",
				Email:        "sam@example.com",
				Phone:        "0 555 333 22 13",
				PasswordHash: password,
				RoleID:       3,
			},
			{
				UserName:     "jennyack",
				Name:         "Jenny",
				Surname:      "Ach",
				Email:        "jenny@example.com",
				Phone:        "0 555 333 22 14",
				PasswordHash: password,
				RoleID:       3,
			},
			{
				UserName:     "luiswutz",
				Name:         "Luis",
				Surname:      "Wutz",
				Email:        "luis@example.com",
				Phone:        "0 555 333 22 15",
				PasswordHash: password,
				RoleID:       3,
			},
			{
				UserName:     "commanderlogar",
				Name:         "Commander",
				Surname:      "Logar",
				Email:        "logar@example.com",
				Phone:        "0 555 333 22 16",
				PasswordHash: password,
				RoleID:       3,
			},
		}

		for _, user := range users {
			err := db.Create(&user).Error
			if err != nil {
				log.Println("Error while seeding customer data. Error: ", err)
				return err
			}
		}
	}

	return nil

}

func seedCategoriesData(db *gorm.DB) error {
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
