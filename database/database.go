package database

import (
	"Streaming-Website-Master/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "root:123456789@tcp(localhost:3306)/streaming_anime?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate schema
	DB.AutoMigrate(&models.Anime{})

	return DB // ✅ Now returns the database instance
}
