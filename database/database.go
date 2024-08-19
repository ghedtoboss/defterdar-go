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
