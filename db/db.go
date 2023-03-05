package db

import (
	"fmt"
	"os"
	"product-api/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Destination{})
	database.AutoMigrate(&model.Order{})
	database.AutoMigrate(&model.OrderItem{})
	database.AutoMigrate(&model.ProductPackage{})
	database.AutoMigrate(&model.Transportation{})
	database.AutoMigrate(&model.Customer{})

	DB = database
	return DB
}
