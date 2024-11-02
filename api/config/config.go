package config

import (
	"content-app-backend/api/models"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    dsn := "host=localhost user=dds password=123 dbname=content-app port=5432 sslmode=disable" 
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    log.Println("Database connected successfully")
	err = DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}