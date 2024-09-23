package database

import (
	"defterdar-go/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBWrite *gorm.DB
var DBRead *gorm.DB

func Connect() {
	dsnWrite := os.Getenv("DSL_WRITE")
	dbWrite, err := gorm.Open(mysql.Open(dsnWrite), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	DBWrite = dbWrite

	dsnRead := os.Getenv("DSL_READ")
	dbRead, err := gorm.Open(mysql.Open(dsnRead), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	DBRead = dbRead

}

func Migrate() {
	if err := DBWrite.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate User model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Shop{}); err != nil {
		log.Fatal("Failed to migrate Shop model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Product{}); err != nil {
		log.Fatal("Failed to migrate Product model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Customer{}); err != nil {
		log.Fatal("Failed to migrate Customer model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Request{}); err != nil {
		log.Fatal("Failed to migrate Request model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Sale{}); err != nil {
		log.Fatal("Failed to migrate Sale model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Ledger{}); err != nil {
		log.Fatal("Failed to migrate Ledger model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatal("Failed to migrate Expense model: ", err)
	}
	if err := DBWrite.AutoMigrate(&models.Subscription{}); err != nil {
		log.Fatal("Failed to migrate Subscription model: ", err)
	}
	log.Println("Database migrated successfully.")
}
