package database

import (
	"defterdar-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DSL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.Customer{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.CashEntry{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.CustomerTransaction{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.Invoice{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		return
	}

	err = DB.AutoMigrate(&models.Receipt{})
	if err != nil {
		return
	}
}

func Setup() {
	err := os.Setenv("DSL", "root:Eses147852@tcp(127.0.0.1:3306)/defterdar?parseTime=true")
	if err != nil {
		return
	}
	Connect()
	Migrate()
	seedDatabase()
}

func Teardown() {
	err := DB.Migrator().DropTable(&models.User{}, &models.Customer{}, &models.CashEntry{}, &models.Invoice{}, &models.Product{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}
}

func seedDatabase() {
	user := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "user",
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
}
