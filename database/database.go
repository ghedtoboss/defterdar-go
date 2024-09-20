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
	// Her bir AutoMigrate işlemi sonrasında hata kontrolü ekleyin
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate User model: ", err)
	}
	if err := DB.AutoMigrate(&models.Shop{}); err != nil {
		log.Fatal("Failed to migrate Shop model: ", err)
	}
	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatal("Failed to migrate Product model: ", err)
	}
	if err := DB.AutoMigrate(&models.Customer{}); err != nil {
		log.Fatal("Failed to migrate Customer model: ", err)
	}
	if err := DB.AutoMigrate(&models.Request{}); err != nil {
		log.Fatal("Failed to migrate Request model: ", err)
	}
	if err := DB.AutoMigrate(&models.Sale{}); err != nil {
		log.Fatal("Failed to migrate Sale model: ", err)
	}
	if err := DB.AutoMigrate(&models.Ledger{}); err != nil {
		log.Fatal("Failed to migrate Ledger model: ", err)
	}
	if err := DB.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatal("Failed to migrate Expense model: ", err)
	}
	if err := DB.AutoMigrate(&models.Subscription{}); err != nil {
		log.Fatal("Failed to migrate Subscription model: ", err)
	}
	log.Println("Database migrated successfully.")
}
